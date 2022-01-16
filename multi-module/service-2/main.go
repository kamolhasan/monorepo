package main

import (
	"fmt"
	service1api "github.com/kamolhasan/monorepo/multi-module/service-1/api"
	"github.com/kamolhasan/monorepo/multi-module/service-2/api"
)

func main() {
	fmt.Println(service1api.Perform())
	fmt.Println(api.Perform())
}
