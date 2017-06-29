package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World! I repeat what you type!")
	reader := bufio.NewReader(os.Stdin)
	for {
		in,_ := reader.ReadString('\n')
		fmt.Println("You said: " + in)
	}

}