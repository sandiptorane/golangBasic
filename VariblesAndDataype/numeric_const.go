package main

import (  
    "fmt"
)
//here value of a is 5 which is generic.it is compatible with other numeric format so below all allowed
func main() {  
    const a = 5
    var intVar int = a
    var int32Var int32 = a
    var float64Var float64 = a
    var complex64Var complex64 = a
    fmt.Println("int=",intVar, "\nint32=", int32Var, "\nfloat64=", float64Var, "\ncomplex64=",complex64Var)
}
