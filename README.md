# cPrices

cPrices is a Go-based cryptocurrency price tracking application that retrieves current prices from various exchanges and sends updates via Telegram. It provides an easy way to monitor cryptocurrency prices for users.

## Features

- Fetch real-time cryptocurrency prices for multiple currencies.
- Support for multiple cryptocurrencies like Bitcoin, Solana, etc.
- Sends price updates through a Telegram bot.
- Configurable via environment variables for flexibility.
- Built with Go, leveraging its concurrency and performance features.

## Getting Started

### Prerequisites

Before you begin, ensure you have the following installed:

- **Go 1.16 or later**
- **Git**
- A **Telegram bot token** (create one using [BotFather](https://core.telegram.org/bots#botfather))

### Installation Steps

1. **Clone the Repository**:
   Open your terminal and run the following command to clone the repository:
   ```bash
   git clone https://github.com/weakennN/cPrices.git
   cd cPrices

2. **Install Dependencies**: 
   Use Go modules for dependency management. Run:
   ```bash
   go mod tidy

3. **Create the .env file**:
   In the root of your project, create a file named ```.env``` and add the following environment variables:
   ```bash
   SERVER_ADDRESS=
   SYMBOLS=
   TELEGRAM_BOT_API_TOKEN=
   TELEGRAM_CHAR_ID=

4. **Run the Application**:
   Use the following command to start the application:
   ```bash
   go run main.go

### Usage

Once the application is running:

- It will begin fetching prices for the specified cryptocurrencies listed in the SYMBOLS environment variable.
- Updates will be sent through the configured Telegram bot.
- To track different cryptocurrencies, modify the SYMBOLS variable in the ```.env``` file.
