package main

import (
	"fmt"
	"net"
)

func main(){
	addr,err := net.LookupAddr("golangbot123.com")
	if err != nil{
		if dnsError,ok := err.(*net.DNSError);ok{
			if dnsError.Timeout(){
				fmt.Println("timeout error")
				return
			}
			if dnsError.Temporary(){
				fmt.Println("temporary error")
				return
			}
			fmt.Println("Generic error",err)
			return
		}
		fmt.Println(addr)

	}
}
