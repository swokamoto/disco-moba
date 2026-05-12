package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("disco-moba starting...")
	_ = os.Getenv("DISCORD_TOKEN") // will be used when bot is wired up
}
