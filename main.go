// Program to set up the dependencies of taskwarror tasks

package main

import (
	"bytes"
	"fmt"
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
	tasklist := os.Args[1:]
	err := isValid(tasklist)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid input: %v\n", err)
		os.Exit(1)
	}

	for i := 1; i < len(tasklist); i++ {
		cmd := exec.Command("task", tasklist[i], "modify", fmt.Sprintf("depends:%v", tasklist[i-1]))
		var out bytes.Buffer
		cmd.Stdout = &out

		if err := cmd.Run(); err != nil {
			fmt.Println("Error: ", err.Error())
		}
	}
}
