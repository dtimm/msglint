package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	emoji "github.com/tmdvs/Go-Emoji-Utils"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	ex := 0
	em := 0
	empty := 0

	for s.Scan() {
		t := s.Text()
		if len(t) == 0 {
			if empty >= 1 {
				break
			}

			empty += 1
		}
		em = len(emoji.FindAll(t))

		for _, c := range t {
			if c == 33 {
				ex += 1
			}
		}
	}

	fmt.Printf("exclaimations: %v\nemoji: %v\n", ex, em)

	if err := s.Err(); err != nil {
		log.Println(err)
	}
}
