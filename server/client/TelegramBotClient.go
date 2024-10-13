package client

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type BotUserResponse struct {
	Ok     bool                `json:"ok"`
	Result BotUserResponseBody `json:"result"`
}

type BotUserResponseBody struct {
	Id        int64  `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
}

type BotMessageResponse struct {
}

func Auth() BotUserResponse {
	url := "https://api.telegram.org/bot" + os.Getenv("TELEGRAM_BOT_API_TOKEN") +"/getMe"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
		return getError()
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Non-OK HTTP status: %v", resp.Status)
		return getError()
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return getError()
	}

	var botUserResponse BotUserResponse
	err = json.Unmarshal(body, &botUserResponse)
	if err != nil {
		log.Printf("Error parsing JSON: %v", err)
		return getError()
	}

	log.Printf("Response: %+v\n", botUserResponse)

	return botUserResponse
}

func SendMessage(message string) BotMessageResponse {
	url := "https://api.telegram.org/bot"+ os.Getenv("TELEGRAM_BOT_API_TOKEN") +"/sendMessage?chat_id=" + os.Getenv("TELEGRAM_CHAR_ID") + "&text=" + message

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error making GET request: %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	log.Print(string(body))

	var botMessageResponse BotMessageResponse
	err = json.Unmarshal(body, &botMessageResponse)
	if err != nil {
		log.Printf("Error parsing JSON: %v", err)
	}

	log.Printf("Response: %+v\n", botMessageResponse)

	return botMessageResponse
}

func getError() BotUserResponse {
	return BotUserResponse{
		Ok: false,
	}
}
