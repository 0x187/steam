# Steam Session

This project provides a Go package for managing a Steam client session. It handles connecting to the Steam network, logging in with optional Steam Guard support, and maintaining the session.

## Features

-   Connect to the Steam network.
-   Login with username and password.
-   Supports Steam Guard with a prompt for the code.
-   Session caching to avoid repeated Steam Guard prompts.
-   Automatic reconnect with exponential backoff on disconnect.

## Getting Started

To use this package, you will need to have Go installed on your system.

\`\`\`bash
go get github.com/your-username/steam-session
\`\`\`

## Usage

\`\`\`go
package main

import (
	"log"

	"github.com/your-username/steam-session/internal/steamclient"
)

func main() {
	client := steamclient.NewClient()

	err := client.Connect()
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer client.Disconnect()

	err = client.Login("your-username", "your-password", "")
	if err != nil {
		log.Fatalf("Failed to login: %v", err)
	}

	// Your application logic here.
}
\`\`\`

## Security

This package does not persist your Steam password. It uses a sentry file to cache your session, which is stored in your user-specific configuration directory with secure permissions.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request with your changes.
