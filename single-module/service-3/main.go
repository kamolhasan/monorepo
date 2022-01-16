package main

import (
	"fmt"
	sapi1 "github.com/kamolhasan/monorepo/single-module/service-1/api"
	sapi2 "github.com/kamolhasan/monorepo/single-module/service-2/api"
)

func main() {
	fmt.Println(sapi1.Perform())
	fmt.Println(sapi2.Perform())
	fmt.Println("Hello from service-3")
}
