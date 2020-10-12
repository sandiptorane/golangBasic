package main

import (  
    "fmt"
)

func main() {  
  fmt.Println("Break if i==j")
    for i := 0; i < 3; i++ {
        for j := 1; j < 4; j++ {
            fmt.Printf("i = %d , j = %d\n", i, j)
            if i == j {
                break
            }
        }

    }
}
