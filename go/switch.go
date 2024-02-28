// SWITCH
// time weeekday hour retrieve


package main

import (
  "fmt"
  "time"
)

func main(){

  i := 2
  fmt.Println("Write", i, " as")
  switch i {
  case 1:
    fmt.Println("one")
  case 2:
    fmt.Println("two")
  case 3:
    fmt.Println("three")
  }

  // You can use commas to separate multiple expressions in the same case statement. 
  switch time.Now().Weekday(){
  case time.Saturday, time.Sunday:
    fmt.Println("It's Weekend")
  default:
    fmt.Println("It's a weekday")  
  }

  // switch without an expression is an alternate way to express if/else logic.
  // Here we also show how the case expressions can be non-constants.
  t := time.Now()
  switch {
  case t.Hour() < 12:
    fmt.Println("It's before noon")
  default:
    fmt.Println("It's after noon")
  }

  // A type switch compares types instead of values. 
  // You can use this to discover the type of an interface value.
  // In this example, the variable t will have the type corresponding to its clause.
  whatAmI := func(i interface{}) {
    switch t := i.(type) {
    case bool:
      fmt.Println("bool")
    case int:
      fmt.Println("int")
    default:
      fmt.Println("Keine ano", t)
    }
  }
  whatAmI(true)
  whatAmI(1)
  whatAmI("hey")
}
