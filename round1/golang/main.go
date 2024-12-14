package main

import (
	"fmt"
	"math/rand"
	"time"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	const countNumbers int = 100000000
	fmt.Println(time.Now())
	var nums [countNumbers]int

	for i := 0; i < countNumbers; i++ {
		nums[i] = rand.Intn(100)
	}
	fmt.Println(time.Now()) //+1.276489701
	fmt.Println(nums[0])
	fmt.Println(nums[countNumbers-1])
}
