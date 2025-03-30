package main

import (
	"fmt"
	tu "github.com/mymmrac/telego/telegoutil"
	"log"
	"os"

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
	var responseText string
	var keyboard *telego.ReplyKeyboardMarkup

	switch message.Text {
	case "Путеводитель":
		responseText = "Вот ваш путеводитель:"
		_, err := b.telegoBot.SendMessage(&telego.SendMessageParams{
			ChatID: telego.ChatID{ID: chatID},
			Text:   responseText,
		})
		if err != nil {
			log.Printf("Failed to send message: %v", err)
		}

		// Открываем фото с диска
		file, err := os.Open("begite.jpg")
		if err != nil {
			log.Printf("Failed to open file: %v", err)
			return
		}
		defer file.Close() // Закрываем файл после отправки

		// Отправляем фото
		_, err = b.telegoBot.SendPhoto(&telego.SendPhotoParams{
			ChatID: telego.ChatID{ID: chatID},
			Photo:  telego.InputFile{File: file},
		})
		if err != nil {
			log.Printf("Failed to send photo: %v", err)
		}
		return // Завершаем выполнение, чтобы не отправлять лишний текст

	case "Сменить язык":
		responseText = "Сменить язык 2"

	case "Время ожидания ответа":
		responseText = "Время 3"

	case "Запросить звонок":
		responseText = "Запрос принят, ожидайте звонка."
		keyboard = tu.Keyboard(
			tu.KeyboardRow(tu.KeyboardButton("Отменить звонок")),
		)
	case "Отменить звонок":
		responseText = "Звонок отменён. Возвращаем вас в главное меню."
		keyboard = generateMainKeyboard()
	default:
		responseText = "Неизвестная команда. Выберите кнопку."
		keyboard = generateMainKeyboard()
	}

	// Отправляем текстовый ответ с клавиатурой (если нужно)
	_, err := b.telegoBot.SendMessage(&telego.SendMessageParams{
		ChatID:      telego.ChatID{ID: chatID},
		Text:        responseText,
		ReplyMarkup: keyboard,
	})
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
