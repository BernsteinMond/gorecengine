package main

import "fmt"

func main() {
	err := run()
	if err != nil {
		fmt.Println("run() returned error: ", err.Error())
	}
}

func run() (err error) {
	return nil
}
