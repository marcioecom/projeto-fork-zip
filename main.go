package main

import "time"

func main() {
	logger(time.Now().String())
	logger("Original")
	Helper()
	Helper()
}

func Helper() {
	logger("Novo no Original")
}

func logger(arg string) {
	println(arg)
}
