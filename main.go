package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"clids/aimlapi"
	"clids/utils"
)

func main() {
	apiKey := os.Getenv("AIMLAPI_KEY")
	if apiKey == "" {
		utils.PrintError(fmt.Errorf("переменная окружения AIMLAPI_KEY не установлена"))
		return
	}

	client := aimlapi.NewAIMLAPIClient(apiKey)
	model := "gpt-4o"

	messages := []aimlapi.Message{
		{
			Role:    "system",
			Content: "Ты полезный ИИ-ассистент.",
		},
	}

	reader := bufio.NewReader(os.Stdin)

	utils.PrintBanner()
	utils.PrintInfo(fmt.Sprintf("Модель: %s", model))
	utils.PrintInfo("Введите 'exit' для выхода, 'clear' для очистки истории, 'history' для просмотра истории, 'help' для справки")
	utils.PrintSeparator()

	for {
		utils.PrintPrompt()
		input, err := reader.ReadString('\n')
		if err != nil {
			utils.PrintError(err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		if strings.ToLower(input) == "exit" {
			utils.PrintSuccess("До свидания!")
			break
		}

		if strings.ToLower(input) == "clear" {
			messages = []aimlapi.Message{{
				Role:    "system",
				Content: "Ты полезный ИИ-ассистент.",
			}}
			utils.PrintSuccess("История чата очищена")
			continue
		}

		if strings.ToLower(input) == "history" {
		    utils.PrintInfo("История сообщений:")
		    for _, msg := range messages {
		        if msg.Role == "system" {
		            continue
		        }
		        
		        role := "👤"
		        if msg.Role == "assistant" {
		            role = "🤖"
		        }
		        
		        fmt.Printf("%s %s:\n%s\n\n", role, strings.Title(msg.Role), msg.Content)
		    }
		    continue
		}

		if strings.ToLower(input) == "help" {
		    utils.PrintInfo("Доступные команды:")
		    fmt.Println("  exit    - выход из программы")
		    fmt.Println("  clear   - очистка истории чата")
		    fmt.Println("  history - показать историю сообщений")
		    fmt.Println("  help    - показать эту справку")
		    fmt.Println()
		    continue
		}

		messages = append(messages, aimlapi.Message{
			Role:    "user",
			Content: input,
		})

		fmt.Printf("%s%s🔄 Обрабатываю запрос...%s\r", utils.ColorYellow, utils.ColorBold, utils.ColorReset)

		response, err := client.Chat(model, messages)

		fmt.Print(strings.Repeat(" ", 30) + "\r")

		if err != nil {
			utils.PrintError(err)
			continue
		}

		messages = append(messages, aimlapi.Message{
			Role:    "assistant",
			Content: response,
		})

		utils.PrintResponse(response)
	}
}
