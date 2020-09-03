package main

import (
	"fmt"

	"github.com/unsmoker/myminio/config"
)

func main() {
	config := config.New()
	fmt.Println(config)
}
