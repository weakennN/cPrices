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
