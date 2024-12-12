# Social-media-feeds-

## Installation

## Prerequisites

Go (version 1.20 or later)

A Discord bot token from the Discord Developer Portal.

API keys for YouTube, Twitch, and (optional) TikTok scraping.

## Steps

Clone the repository:

``` bash

git clone https://github.com/yourusername/social-media-feeds.git

```

`cd discord-social-media-bot ` 

Initialize the Go module:

`go mod tidy`  

Configure the bot by creating a config.json file:

``` json
{
  "botToken": "YOUR_DISCORD_BOT_TOKEN",
  "discordChannelID": "DISCORD_CHANNEL_ID_FOR_NOTIFICATIONS",
  "youtubeApiKey": "YOUR_YOUTUBE_API_KEY",
  "twitchClientId": "YOUR_TWITCH_CLIENT_ID",
  "twitchAccessToken": "YOUR_TWITCH_ACCESS_TOKEN",
  "tiktokUsers": ["example_user1", "example_user2"],
  "youtubeChannels": ["UC_x5XG1OV2P6uZZ5FSM9Ttw", "UCVHFbqXqoYvEWM1Ddxl0QDg"],
  "twitchUsers": ["example_user1", "example_user2"],
  "customMessage": "🎉 Check this out! New content is live!"
}

```


Run the bot:

`go run main.go ` 

## Usage

Commands
```
/help
Displays a list of available commands.
/channelinfo <platform> <channel>
Fetches information about a YouTube channel, Twitch streamer, or TikTok user.
Example:
/channelinfo youtube UC_x5XG1OV2P6uZZ5FSM9Ttw  
Notifications
The bot automatically monitors configured YouTube, Twitch, and TikTok accounts and posts updates in the specified Discord channel.
Notifications include a customizable message from config.json.
APIs and Authentication

```


YouTube API

Obtain a YouTube API key from the Google Cloud Console.
Twitch API

Create an application in the Twitch Developer Console.
Generate a client ID and access token.

TikTok API
TikTok does not have an official API for post monitoring. This implementation may rely on third-party APIs or scraping.

Contributing

Fork the repository.
Create a new branch:

``git checkout -b feature-name``  

Make your changes and commit them:

``git commit -m "Added feature" `` 

Push to the branch:

`git push origin feature-name ` 

Open a pull request.

## License

This project is licensed under the MIT License. See LICENSE for more details.


