package main

import "fmt"

func helper(a int, b int) int {
  return a * b
}

func sayHello(s1, s2 string) (string, error) {
  if s1 == "apple" {
    return "", fmt.Errorf("Error dwag")
  }

  return s1 + s2, nil
}

func skypia(a, b, c int) (int, error) {
  fmt.Errorf("Hello")
  return a, fmt.Errorf("Skypia error ")
}

func main() {
  if allTheval := helper(11, 10); allTheval == 100 {
    fmt.Println("You value is precisely ", allTheval)
  } else {
    fmt.Println(allTheval, " is not precise")
  }

  if val, err := skypia(2, 2, 3); err != nil {
    fmt.Println("An error occured ", val)
  } else {
    fmt.Println("Test passed all good")
  }

  var a string = "avacado"
  var b string = "kiwi"

  c, _ := sayHello(a, b)
  fmt.Println(c);

}
