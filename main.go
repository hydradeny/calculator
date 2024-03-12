package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	result, _ := Calculate(text)
	fmt.Printf("%v", result)
}

func Calculate(text string) (string, error) {
	return text, nil
}
