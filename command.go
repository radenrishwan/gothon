package gothon

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Run(fileName string) error {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	token := NewScanner(string(bytes)).ScanTokens()
	for _, t := range token {
		fmt.Println(t)
	}

	return nil
}

func Help() {
	fmt.Println("usage: gothon <filename>")
}

func Prompt() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)

			continue
		}

		buffer := strings.TrimSpace(text)
		if buffer == "" {
			break
		}

		fmt.Println("You said: ", buffer)
	}
}
