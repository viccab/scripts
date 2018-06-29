// menu.go


package main

import (
    "fmt"
)


func add() {
  fmt.Println("Processing 1")
}

func change() {
  fmt.Println("Processing 2")
}

func view() {
  fmt.Println("Processing 3")
}

func menu() {
  fmt.Println("\nMain Menu")
  fmt.Println("1 - Add")
  fmt.Println("2 - Change")
  fmt.Println("3 - View")
  fmt.Println("4 - Quit\n")
  fmt.Print("Option: ")
}

func main() {
  var ch string

  for {
     menu()
     fmt.Scan(&ch)
     switch ch {
        case "1":
           add()
        case "2":
           change()
        case "3":
           view()
        case "4":
           return
      } 
  }
}

//eof
