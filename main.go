package main

import (
	"fmt"
	"start-link/cmd"
)

func main() {
	fmt.Println("hello world")
	cmd.Execute()
}

//docker run -it -d --restart always --name purchase my-golang-app
