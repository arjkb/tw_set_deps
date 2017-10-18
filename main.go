package main

import (
  // "bytes"
  "fmt"
  "os"
  // "os/exec"
)

func main()  {
  tasklist := os.Args[1:]
  fmt.Println(tasklist)

  // cmd := exec.Command("ls")

  // var out bytes.Buffer
  // cmd.Stdout = &out
  // err := cmd.Run()
  // if err != nil {
  //   fmt.Println("Something went wrong!")
  // }
  // fmt.Println(" >> ", out.String())
}
