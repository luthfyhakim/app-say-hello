package main

import (
	"fmt"

	// update major menggunakan versi tags
	go_say_hello "github.com/luthfyhakim/go-say-hello/v2"
)

func main() {
	fmt.Println(go_say_hello.SayHello("Hakim"))
}
