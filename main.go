package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/anthropics/anthropic-sdk-go"

	"agent/agent"
	"agent/logger"
	"agent/tools"
)

func main() {
	// Initialize logger
	logDir := filepath.Join(".", "logs")
	if err := logger.Initialize(logDir); err != nil {
		log.Fatal("Failed to initialize logger:", err)
	}
	defer logger.Close()

	client := anthropic.NewClient()

	scanner := bufio.NewScanner(os.Stdin)
	getUserMessage := func() (string, bool) {
		if !scanner.Scan() {
			return "", false
		}
		return scanner.Text(), true
	}

	toolDefs := []tools.ToolDefinition{
		tools.ReadFileDefinition,
		tools.ListFilesDefinition,
		tools.EditFileDefinition,
		tools.RunShellCommandDefinition,
	}
	myAgent := agent.NewAgent(&client, getUserMessage, toolDefs)
	err := myAgent.Run(context.TODO())
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}
