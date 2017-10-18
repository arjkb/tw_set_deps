// Program to set up the dependencies of taskwarror tasks

package main

import (
  "bytes"
  "fmt"
  "os"
  "unicode"
  "os/exec"
)

func isValid(s []string) (bool, string)  {
  for _, word := range s  {
    for _, ch := range word {
      if !unicode.IsDigit(ch)  {
        return false, word
      }
    }
  }
  return true, ""
}

func main()  {
  tasklist := os.Args[1:]
  valid, word := isValid(tasklist)
  if !valid {
    panic(fmt.Sprintf(" Invalid input parameter: %q\n\tUse space separated numbers", word))
  }

  for i := 1; i < len(tasklist); i++  {
    cmd := exec.Command("task", tasklist[i], "modify", fmt.Sprintf("depends:%v", tasklist[i-1]))
    var out bytes.Buffer
    cmd.Stdout = &out

    if err := cmd.Run(); err != nil {
      fmt.Println("Error: ", err.Error())
    }
  }
}
