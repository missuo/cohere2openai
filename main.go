/*
 * @Author: Vincent Yang
 * @Date: 2024-04-16 22:58:22
 * @LastEditors: Vincent Yang
 * @LastEditTime: 2024-04-18 04:33:58
 * @FilePath: /cohere2openai/main.go
 * @Telegram: https://t.me/missuo
 * @GitHub: https://github.com/missuo
 *
 * Copyright © 2024 by Vincent, All Rights Reserved.
 */

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func parseAuthorizationHeader(c *gin.Context) (string, error) {
	authorizationHeader := c.GetHeader("Authorization")
	if !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return "", fmt.Errorf("invalid Authorization header format")
	}
	return strings.TrimPrefix(authorizationHeader, "Bearer "), nil
}

func cohereRequest(c *gin.Context, openAIReq OpenAIRequest) {
	apiKey, err := parseAuthorizationHeader(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cohereReq := CohereRequest{
		Model:       openAIReq.Model,
		ChatHistory: []ChatMessage{},
		Message:     "",
		Stream:      openAIReq.Stream,
	}

	for _, msg := range openAIReq.Messages {
		if msg.Role == "user" {
			cohereReq.Message = msg.Content
		} else {
			var role string
			if msg.Role == "assistant" {
				role = "CHATBOT"
			} else if msg.Role == "system" {
				role = "SYSTEM"
			} else {
				role = "USER"
			}
			cohereReq.ChatHistory = append(cohereReq.ChatHistory, ChatMessage{
				Role:    role,
				Message: msg.Content,
			})
		}
	}

	reqBody, _ := json.Marshal(cohereReq)
	req, err := http.NewRequest("POST", "https://api.cohere.ai/v1/chat", bytes.NewBuffer(reqBody))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	reader := resp.Body
	buffer := make([]byte, 1024)

	isFirstChunk := true

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var cohereResp CohereResponse
		err = json.Unmarshal(buffer[:n], &cohereResp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if cohereResp.IsFinished {
			var resp OpenAIResponse
			resp.ID = "chatcmpl-123"
			resp.Object = "chat.completion.chunk"
			resp.Created = time.Now().Unix()
			resp.Model = openAIReq.Model
			resp.Choices = []OpenAIChoice{
				{
					Index:        0,
					Delta:        OpenAIDelta{},
					Logprobs:     nil,
					FinishReason: stringPtr("stop"),
				},
			}

			respBytes, _ := json.Marshal(resp)
			c.Data(http.StatusOK, "application/json", []byte("data: "))
			c.Data(http.StatusOK, "application/json", respBytes)
			c.Data(http.StatusOK, "application/json", []byte("\n\n"))

			c.Data(http.StatusOK, "application/json", []byte("data: [DONE]\n\n"))
			break
		} else {
			var resp OpenAIResponse
			resp.ID = "chatcmpl-123"
			resp.Object = "chat.completion.chunk"
			resp.Created = time.Now().Unix()
			resp.Model = openAIReq.Model

			if !isFirstChunk {
				resp.Choices = []OpenAIChoice{
					{
						Index:        0,
						Delta:        OpenAIDelta{Content: cohereResp.Text},
						Logprobs:     nil,
						FinishReason: nil,
					},
				}
			} else {
				resp.Choices = []OpenAIChoice{
					{
						Index:        0,
						Delta:        OpenAIDelta{},
						Logprobs:     nil,
						FinishReason: nil,
					},
				}
				isFirstChunk = false
			}

			respBytes, _ := json.Marshal(resp)
			c.Data(http.StatusOK, "application/json", []byte("data: "))
			c.Data(http.StatusOK, "application/json", respBytes)
			c.Data(http.StatusOK, "application/json", []byte("\n\n"))
		}
	}
}

func handler(c *gin.Context) {
	var openAIReq OpenAIRequest

	if err := c.BindJSON(&openAIReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	allowModels := []string{"command-r-plus", "command-r", "command", "command-light", "command-light-nightly", "command-nightly"}

	if !isInSlice(openAIReq.Model, allowModels) {
		openAIReq.Model = "command-r-plus"
	}
	if openAIReq.Stream {
		cohereRequest(c, openAIReq)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Stream is not enabled",
		})
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Cohere2OpenAI, Made by Vincent Yang. https://github.com/missuo/cohere2openai",
		})
	})
	r.POST("/v1/chat/completions", handler)
	r.GET("/v1/models", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"object": "list",
			"data": []gin.H{
				{
					"id":       "command-r",
					"object":   "model",
					"created":  1692901427,
					"owned_by": "system",
				},
				{
					"id":       "command-r-plus",
					"object":   "model",
					"created":  1692901427,
					"owned_by": "system",
				},
				{
					"id":       "command-light",
					"object":   "model",
					"created":  1692901427,
					"owned_by": "system",
				},
				{
					"id":       "command-light-nightly",
					"object":   "model",
					"created":  1692901427,
					"owned_by": "system",
				},
				{
					"id":       "command",
					"object":   "model",
					"created":  1692901427,
					"owned_by": "system",
				},
				{
					"id":       "command-nightly",
					"object":   "model",
					"created":  1692901427,
					"owned_by": "system",
				},
			},
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Path not found",
		})
	})
	r.Run(":6600")
}
