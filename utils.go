/*
 * @Author: Vincent Yang
 * @Date: 2024-04-18 03:50:36
 * @LastEditors: Vincent Yang
 * @LastEditTime: 2024-04-18 03:50:56
 * @FilePath: /cohere2openai/utils.go
 * @Telegram: https://t.me/missuo
 * @GitHub: https://github.com/missuo
 *
 * Copyright © 2024 by Vincent, All Rights Reserved.
 */

package main

func isInSlice(str string, list []string) bool {
	for _, item := range list {
		if item == str {
			return true
		}
	}
	return false
}

func stringPtr(s string) *string {
	return &s
}
