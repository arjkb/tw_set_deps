package main

import (
  // "bytes"
  "fmt"
  "os"
  // "os/exec"
)

func main()  {
  tasklist := os.Args[1:]
  fmt.Println(tasklist, tasklist[0], tasklist[len(tasklist) -1])
  fmt.Println(len(tasklist))

  for i := 1; i < len(tasklist); i++  {
    command := "task " + tasklist[i] + " modify depends:" + tasklist[i-1]
    fmt.Println(command)
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
