package main

import (
  // "bytes"
  "fmt"
  "os"
  "reflect"
  "unicode"
  // "os/exec"
)

func isValid(s []string)(bool, string)  {
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
  fmt.Println(tasklist, tasklist[0], tasklist[len(tasklist) -1])
  fmt.Println(len(tasklist))
  fmt.Println(reflect.TypeOf(tasklist[0]))

  valid, word := isValid(tasklist)

  if !valid {
    panic(fmt.Sprintf(" Invalid input parameter: %q\n\tUse space separated numbers", word))
  }

  for i := 1; i < len(tasklist); i++  {
    command := "task " + tasklist[i] + " modify depends:" + tasklist[i-1]
    fmt.Println(command)
    for _, ch := range tasklist[i]  {
      fmt.Printf(" %c %v\n", ch, unicode.IsDigit(ch))
    }
  }

  // cmd := exec.Command("ls")

  // var out bytes.Buffer
  // cmd.Stdout = &out
  // err := cmd.Run()
  // if err != nil {
  //   fmt.Println("Something went wrong!")
  // }
  // fmt.Println(" >> ", out.String())
}
