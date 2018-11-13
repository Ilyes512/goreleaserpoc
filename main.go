package main

import "fmt"

var (
	version = "NOT_SET"
	commit  = "NOT_SET"
	date    = "NOT_SET"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println("Version: ", version)
	fmt.Println("Commit: ", commit)
	fmt.Println("Date: ", date)
}
