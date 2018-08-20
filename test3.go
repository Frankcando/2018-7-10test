package main

import (
	"fmt"
	"strategy"
	"time"
)

func main() {

	t_start := time.Now()
	fmt.Println(t_start)
	fmt.Println("---data test begin---")

	strategy.Du_Init()
	fmt.Println("Init end")

	strategy.Du_gg_test()

	fmt.Println("---data test end---")
	t_end := time.Now()
	fmt.Println(t_end)
}
