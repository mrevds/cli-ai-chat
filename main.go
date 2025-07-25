package main

import (
	"clids/api"
	"clids/config"
	"clids/utils"
	"fmt"
	"os"
	"strings"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		utils.PrintError(err)
		os.Exit(1)
	}

	client := api.NewDeepSeekClient(cfg.APIKey)

	utils.PrintSuccess("DeepSeek CLI Chat запущен")
	fmt.Println("Введите ваш запрос (или 'exit' для выхода):")

	for {
		userInput, err := utils.GetUserInput("> ")
		if err != nil {
			utils.PrintError(err)
			continue
		}

		if strings.ToLower(userInput) == "exit" {
			break
		}

		if utils.IsEmptyInput(userInput) {
			continue
		}

		response, err := client.SendMessage(userInput)
		if err != nil {
			utils.PrintError(err)
			continue
		}

		fmt.Println("\nОтвет:")
		fmt.Println(utils.FormatResponse(response))
		fmt.Println()
	}
}
