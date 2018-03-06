package main

import "CoursePlatform/Auth"

func main() {
	go Auth.StartServer()
}
