package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/joho/godotenv"

	"cPrices/client"
	telegramBotClient "cPrices/client"
)

func main() {
	err := godotenv.Load("../.env")
    if err != nil {
        fmt.Println("Error loading .env file")
        return
    }

 	go sendRates()

	runServer()
}

func sendRates() {
	waitUntilNextInterval()

	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for {
		symbols := strings.Split(os.Getenv("SYMBOLS"), ",")
		coins, error := client.GetRates(symbols)
		if error != nil {
			log.Fatal("Couldnt fetch rates");
		}

		for i := 0; i < len(coins); i++ {
			coin := coins[i];
			client.SendMessage(fmt.Sprintf("%s: Current price: %.2f (last 5 minutes)", coin.Name, coin.Price));
		}
		
		<-ticker.C
	}
}

func waitUntilNextInterval() {
	now := time.Now()
	next := now.Truncate(5 * time.Minute).Add(5 * time.Minute)

	sleepDuration := next.Sub(now)
	fmt.Printf("Waiting %v until next interval at %v\n", sleepDuration, next)

	time.Sleep(sleepDuration)
}

func runServer() {
	server := &http.Server{
		Addr:    os.Getenv("SERVER_ADDRESS"),
		Handler: nil,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		botUser := telegramBotClient.Auth()
		fmt.Println(botUser)
		if !botUser.Ok {
			log.Fatal("No bot available")
			<-quit
		}

		fmt.Println("Starting server on port 8080...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error starting server: %v\n", err)
		}
	}()

	<-quit
	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Error during shutdown: %v\n", err)
	}

	fmt.Println("Server stopped.")
}
