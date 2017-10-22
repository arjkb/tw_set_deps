// Program to set up the dependencies of taskwarror tasks

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"unicode"
)

func isValid(words ...string) error {
	for i, word := range words {
		for _, ch := range word {
			if !unicode.IsDigit(ch) {
				return fmt.Errorf("bad word %q (param %d)", word, i+1)
			}
		}
	}
	return nil
}

func main() {
	log.SetPrefix("tw_set_deps: ")
	log.SetFlags(0)
	tasklist := os.Args[1:]
	if err := isValid(tasklist...); err != nil {
		log.Fatalf("invalid input: %v\n", err)
	}

	for i := 1; i < len(tasklist); i++ {
		cmd := exec.Command("task", tasklist[i], "modify", fmt.Sprintf("depends:%v", tasklist[i-1]))
		if err := cmd.Run(); err != nil {
			log.Fatalf("%s fails: %v", cmd.Args, err)
		}
	}
}
