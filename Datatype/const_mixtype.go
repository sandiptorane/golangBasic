package main

func main() {  
        var defaultName = "Sam" //allowed
        type customString string
        var customName customString = "Sam" //allowed
        customName = defaultName //not allowed because customName is of type customString and defaultName is string type

}
