package main

import (
	"fmt"
	"github.com/hachi-n/awesomeProject3/lib/config"
)

func main() {
	c := config.LoadConfig()
	fmt.Println(c)
}
