/*
 * @Author: Vincent Yang
 * @Date: 2024-04-16 22:58:27
 * @LastEditors: Vincent Yang
 * @LastEditTime: 2024-04-18 04:34:55
 * @FilePath: /cohere2openai/types.go
 * @Telegram: https://t.me/missuo
 * @GitHub: https://github.com/missuo
 *
 * Copyright © 2024 by Vincent, All Rights Reserved.
 */

package main

type OpenAIRequest struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
	Stream bool `json:"stream"`
}

type CohereRequest struct {
	Model       string        `json:"model"`
	ChatHistory []ChatMessage `json:"chat_history"`
	Message     string        `json:"message"`
	Stream      bool          `json:"stream"`
}

type ChatMessage struct {
	Role    string `json:"role"`
	Message string `json:"message"`
}

type CohereResponse struct {
	IsFinished bool   `json:"is_finished"`
	EventType  string `json:"event_type"`
	Text       string `json:"text"`
}

type OpenAIResponse struct {
	ID      string         `json:"id"`
	Object  string         `json:"object"`
	Created int64          `json:"created"`
	Model   string         `json:"model"`
	Choices []OpenAIChoice `json:"choices"`
}

type OpenAIChoice struct {
	Index        int         `json:"index"`
	Delta        OpenAIDelta `json:"delta"`
	Logprobs     interface{} `json:"logprobs"`
	FinishReason *string     `json:"finish_reason"`
}

type OpenAINonStreamResponse struct {
	ID      string                  `json:"id"`
	Object  string                  `json:"object"`
	Created int64                   `json:"created"`
	Model   string                  `json:"model"`
	Choices []OpenAINonStreamChoice `json:"choices"`
}

type OpenAINonStreamChoice struct {
	Index        int         `json:"index"`
	Message      OpenAIDelta `json:"message"`
	FinishReason *string     `json:"finish_reason"`
}

type OpenAIDelta struct {
	Role    string `json:"role,omitempty"`
	Content string `json:"content,omitempty"`
}
