package main

import (
	"fmt"

	"wb-tech/l2/18/internal/config"
)

func main() {
	conf := config.NewConfig()

	fmt.Println(conf.Server.Port)
}
