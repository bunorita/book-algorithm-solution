package util

import (
	"bufio"
	"io"
	"log"
	"testing"

	"github.com/google/go-cmp/cmp"
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

func TestDiff(t *testing.T, got, want interface{}) {
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("value mismatch (-want +got):\n%s", diff)
	}
}
