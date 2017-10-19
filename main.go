// Program to set up the dependencies of taskwarror tasks

package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"unicode"
)

func isValid(s []string) error {
	for i, word := range s {
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
	err := isValid(tasklist)
	if err != nil {
		log.Fatalf("invalid input: %v\n", err)
		os.Exit(1)
	}

	for i := 1; i < len(tasklist); i++ {
		cmd := exec.Command("tlask", tasklist[i], "modify", fmt.Sprintf("depends:%v", tasklist[i-1]))
		var out bytes.Buffer
		cmd.Stdout = &out

		fmt.Println(cmd.Args)

		if err := cmd.Run(); err != nil {
			log.Fatalf("%s fails: %v", cmd.Args, err)
		}
	}
}
