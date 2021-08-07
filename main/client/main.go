package main

import (
	client "github.com/AfvanJaffer/racy/client"
)

func main() {
	config := client.LoadClientConfig()

	client.RunClient(config.Host, config.Port)
}
