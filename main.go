package main

import (
	"fmt"
	"github.com/user-name/skeleton-name/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalln(fmt.Sprintf("failed to start the app: %v", err))
	}
}
