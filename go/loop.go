package main

import "fmt"

func main() {
  i := 1
  for i <= 3 {
    fmt.Println(i)
    i = i +1
  }

  for j:= 0; j < 3; j++ {
    fmt.Println(j)
  }

  // do this N times 
  for i:= range 3 {
    fmt.Println("range", i)
  }

  // without condition will loop until a break, need to use a "break"
  for {
    fmt.Println("break loop")
    break 
  }

  for n := range 6 {
    if n%2 == 0 {
       continue
    }
    fmt.Println(n)
  } 

}  
