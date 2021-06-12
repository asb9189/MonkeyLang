package main

import (
  "fmt"
  "os"
  //"os/user"
  "MonkeyLang/repl"
)

func main() {
  //user, err := user.Current()
  //if err != nil {
    //panic(err)
  //}
  fmt.Printf("🙊 MonkeyLang 🙈 Interpreter! 😏 \n")
  repl.Start(os.Stdin, os.Stdout)
}
