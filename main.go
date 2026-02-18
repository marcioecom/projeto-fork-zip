package main

import "time"

func main() {
	logger(time.Now().String())
	logger("Original")
	custom(time.Now().String())
	custom("Original")
	Helper()
}

func Helper() {
	logger("Novo no Original")
}

func logger(arg string) {
	println(arg)
	custom("Novo no Original")
}

func custom(arg string) {
	println(arg)
	println("oi")
}
