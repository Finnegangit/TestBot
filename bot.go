package main

import (
	"fmt"
	tu "github.com/mymmrac/telego/telegoutil"
	"log"

	"github.com/mymmrac/telego"
)

type BotService struct {
	telegoBot *telego.Bot
}

func NewBotService() *BotService {
	botToken := "8192811696:AAH-W6EDVmZWJ1NnYLgEPtyjWgplrvVxcAI"
	telegoBot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Printf("failed to initialize telegram bot: %v\n", err)
		log.Fatal(err)
	}
	return &BotService{
		telegoBot: telegoBot,
	}
}

func (b *BotService) Run() {
	updatesChannel, err := b.telegoBot.UpdatesViaLongPolling(nil)
	if err != nil {
		return
	}
	defer b.telegoBot.StopLongPolling()
	for update := range updatesChannel {
		if update.Message != nil {
			b.processMessage(update.Message)
		}
	}
}

func (b *BotService) processMessage(message *telego.Message) {
	chatID := message.Chat.ID
	responseText := ""

	// Проверяем, какая кнопка была нажата
	switch message.Text {
	case "Путеводитель":
		responseText = "Путеводитель 1"
	case "Сменить язык":
		responseText = "Сменить язык 2"
	case "Время ожидания ответа":
		responseText = "Время 3"
	case "Запросить звонок":
		responseText = "Запросить звонок 4"
	default:
		responseText = "Неизвестная команда. Выберите кнопку."
	}

	// Отправляем ответ
	msg := &telego.SendMessageParams{
		ChatID: telego.ChatID{ID: chatID},
		Text:   responseText,
	}
	_, err := b.telegoBot.SendMessage(msg)
	if err != nil {
		log.Printf("Failed to send message: %v", err)
	}
}

// Функция для генерации главного меню
func generateMainKeyboard() *telego.ReplyKeyboardMarkup {
	return tu.Keyboard(
		tu.KeyboardRow(
			tu.KeyboardButton("Запросить звонок"),
			tu.KeyboardButton("Путеводитель"),
		),
		tu.KeyboardRow(
			tu.KeyboardButton("Сменить язык"),
			tu.KeyboardButton("Время ожидания ответа"),
		),
	)
}
