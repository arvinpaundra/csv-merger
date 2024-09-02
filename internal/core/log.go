package core

import (
	"fmt"
	"strings"
)

func LogInfo(messages ...any) {
	var sb strings.Builder

	for _, message := range messages {
		sb.WriteString(fmt.Sprintf("%v", message))
	}

	fmt.Printf("\033[34m[INFO] %s\033[0m\n", sb.String())
}

func LogError(messages ...any) {
	var sb strings.Builder

	for _, message := range messages {
		sb.WriteString(fmt.Sprintf("%v", message))
	}

	fmt.Printf("\033[31m[ERROR] %s\033[0m\n", sb.String())
}

func LogSuccess(messages ...any) {
	var sb strings.Builder

	for _, message := range messages {
		sb.WriteString(fmt.Sprintf("%v", message))
	}

	fmt.Printf("\033[32m[SUCCESS] %s\033[0m\n", sb.String())
}
