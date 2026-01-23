package channel

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/genai"
)

func StartGeminiChat() {
	// Создаём клиент Gemini один раз
	ctx := context.Background()
	cfg := &genai.ClientConfig{
		APIKey: "AIzaSyD5nylQUiZbzqcxeOIGedPM51taCPptqGE", // вставь сюда ключ
	}
	client, err := genai.NewClient(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		// 1️⃣ Чтение ввода пользователя
		fmt.Print("You: ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)
		if userInput == "" {
			continue // игнорируем пустой ввод
		}

		// 3️⃣ Отправляем запрос к Gemini и ждём ответа
		result, err := client.Models.GenerateContent(
			ctx,
			"gemini-3-flash-preview",
			genai.Text(userInput),
			nil,
		)
		if err != nil {
			fmt.Println("Ошибка Gemini:", err)
			continue
		}

		// 4️⃣ Печатаем ответ Gemini
		fmt.Println("Gemini:", result.Text())
	}
}

/*
большую часть ии сделал!!!
*/
