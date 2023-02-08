package util

import (
	"bufio"
	"io"
	"log"
)

func Scan(r io.Reader) []string {
	texts := make([]string, 0)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		texts = append(texts, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return texts
}
