package main

import (
	"fmt"
	service1api "github.com/kamolhasan/monorepo/single-module/service-1/api"
	"github.com/kamolhasan/monorepo/single-module/service-2/api"
)

func main() {
	fmt.Println(service1api.Perform())
	fmt.Println(api.Perform())
}
