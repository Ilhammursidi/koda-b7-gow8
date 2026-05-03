package number

import (
	"fmt"
)

// buat sebuah fungsi bernama ProcessNumber yang digunakan untuk menangani suatu daftar angka dengan ketentuan
// - jika inputnya bernilai nil, maka berikan error "No data Provided"
// - jika daftar inputnya kosong, maka timbulkan panic dengan pesan "empty list provided"
// - jika inputnya valid, maka tampilkan ke layar hasil perkalian dari masing-masing bilangan di dalam input tersebut dengan angka 2
// - jangan lupa tangani panicnya

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
