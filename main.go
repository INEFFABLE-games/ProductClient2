package main

import (
	"context"
	"main/src/client"
)

func main() {
	client.Start(context.Background())
}
