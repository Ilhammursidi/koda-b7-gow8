package number

import (
	"fmt"
)

func ProcessNumber(n []int) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("PANIC : ", err)
		}
	}()

	if n == nil {
		fmt.Printf("No data Provided %d\n", n)
	}

	if 0 == len(n) {
		panic("empty list provided")
	}

	for i := 1; i < len(n)+1; i++ {
		result := i * 2
		fmt.Printf("%d X 2 = %d\n", i, result)
	}
	return nil
}
