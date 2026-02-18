package main

import "time"

func main() {
	println(time.Now().String())
	println("Original")
	Helper()
}

func Helper() {
	println("Novo no Original")
}
