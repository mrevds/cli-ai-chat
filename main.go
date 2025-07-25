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
	// Получаем API ключ из переменной окружения
	apiKey := os.Getenv("AIMLAPI_KEY")
	if apiKey == "" {
		utils.PrintError(fmt.Errorf("переменная окружения AIMLAPI_KEY не установлена"))
		return
	}

	// Создаем клиент
	client := aimlapi.NewAIMLAPIClient(apiKey)
	model := "gpt-4o"

	// Инициализируем историю сообщений
	messages := []aimlapi.Message{
		{
			Role:    "system",
			Content: "Ты полезный ИИ-ассистент.",
		},
	}

	// Создаем reader для ввода
	reader := bufio.NewReader(os.Stdin)

	// Показываем баннер
	utils.PrintBanner()
	utils.PrintInfo(fmt.Sprintf("Модель: %s", model))
	utils.PrintInfo("Введите 'exit' для выхода, 'clear' для очистки истории, 'help' для справки")
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

		// Команда выхода
		if strings.ToLower(input) == "exit" {
			utils.PrintSuccess("До свидания!")
			break
		}

		// Команда очистки истории
		if strings.ToLower(input) == "clear" {
			messages = []aimlapi.Message{{
				Role:    "system",
				Content: "Ты полезный ИИ-ассистент.",
			}}
			utils.PrintSuccess("История чата очищена")
			continue
		}

		// Команда справки
		if strings.ToLower(input) == "help" {
			utils.PrintInfo("Доступные команды:")
			fmt.Println("  exit  - выход из программы")
			fmt.Println("  clear - очистка истории чата")
			fmt.Println("  help  - показать эту справку")
			fmt.Println()
			continue
		}

		// Добавляем сообщение пользователя в историю
		messages = append(messages, aimlapi.Message{
			Role:    "user",
			Content: input,
		})

		// Показываем индикатор загрузки
		fmt.Printf("%s%s🔄 Обрабатываю запрос...%s\r", utils.ColorYellow, utils.ColorBold, utils.ColorReset)

		// Отправляем запрос
		response, err := client.Chat(model, messages)

		// Очищаем индикатор загрузки
		fmt.Print(strings.Repeat(" ", 30) + "\r")

		if err != nil {
			utils.PrintError(err)
			continue
		}

		// Добавляем ответ в историю
		messages = append(messages, aimlapi.Message{
			Role:    "assistant",
			Content: response,
		})

		// Выводим ответ
		utils.PrintResponse(response)
	}
}
