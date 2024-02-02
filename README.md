# Discord Poll Bot

This Discord bot enhances engagement within Discord servers by enabling users to create and participate in polls. It utilizes the `discordgo` library for Discord API interaction and `godotenv` for managing environment variables efficiently.

## Features

- **Poll Creation**: Allows users to create polls with customizable options.
- **Vote Using Reactions**: Participants can vote on polls by reacting with emojis.
- **Dynamic Poll Updates**: Updates poll messages in real-time to reflect current votes.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go 1.15 or newer
- A Discord bot token ([Creating a Discord bot](https://discord.com/developers/docs/intro))
- Git (for cloning the repository)

### Setup

1. **Clone the Repository**

   ```bash
   git clone [https://github.com/yourgithubusername/discord-poll-bot.git](https://github.com/MaratYskak/test-ON-esports.git)
   cd discord-poll-bot
   
2. **Configure Environment Variables**
    create a '.env' file and fill in your Discord bot token:
   ```bash
   DISCORD_BOT_TOKEN=your_discord_bot_token_here
   
3. **Install Dependencies**
   ```bash
   go mod tidy

4. **Run the Bot**
   ```bash
   go run .

### Usage

To initiate a poll, type !poll in the chat followed by your question and options separated by semicolons (;). For example:
```bash
!poll What is your favorite programming language?;Go;Python;JavaScript;Rust

