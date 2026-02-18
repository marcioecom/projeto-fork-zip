package main

import "time"

func main() {
	custom(time.Now().String())
	custom("Original")
	Helper()
}

func Helper() {
	custom("Novo no Original")
}

func custom(arg string) {
	println(arg)
	println("oi")
}
