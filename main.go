package main

import (
	"fmt"
	"sync"

	"github.com/ilhammursidi/koda-b7-gow8/internals/number"
	"github.com/ilhammursidi/koda-b7-gow8/internals/task3"
	"github.com/ilhammursidi/koda-b7-gow8/internals/task4"
	"github.com/ilhammursidi/koda-b7-gow8/internals/webFetcher"
)

func main() {
	// 1
	fmt.Println()
	fmt.Println("nomor 1")
	fmt.Println()
	// list angka
	n := []int{1, 2, 3, 4, 5}
	// list kosong
	kosong := []int{}
	// process list n
	number.ProcessNumber(n)
	// process input nil
	number.ProcessNumber(nil)
	// process input kosong atau slice kosong
	number.ProcessNumber(kosong)

	// 2
	fmt.Println()
	fmt.Println("nomor 2")
	fmt.Println()
	urlChan := make(chan string)
	var wg sync.WaitGroup
	urls := []string{"url1com", "url2com", "url3com", "url4com", "url5com"}

	var wgRecv sync.WaitGroup
	wgRecv.Go(func() {
		webFetcher.FetchResult(urlChan)
	})

	for _, u := range urls {
		wg.Add(1)
		go webFetcher.WebFetcher(urlChan, u, &wg)
	}
	wg.Wait()
	close(urlChan)
	wgRecv.Wait()

	// 3
	fmt.Println()
	fmt.Println("nomor 3")
	fmt.Println()
	users := task3.NewUserManager()

	users.AddUser(1, "ilham")
	users.AddUser(2, "adi")
	users.AddUser(3, "aco")

	// dengan id yang sudah ada
	users.AddUser(3, "baru")

	// tampilkan semua users yang sudah ditambahkan
	for _, u := range users.AllUsers {
		fmt.Printf("%d %s\n", u.Id, u.Name)
	}

	// menampilkan 1 users yang sudah ada
	printUser(users.GetUser(1))
	// menampilkan user yang tidak ada
	printUser(users.GetUser(99))

	//4
	fmt.Println()
	fmt.Println("nomor 4")
	fmt.Println()

	Rectangle := task4.Rectangle{Width: 20, Height: 16}
	Circle := task4.Circle{Radius: 10}
	fmt.Print("Rectangle ")
	fmt.Println(task4.Calculator(Rectangle))
	fmt.Println("=================")
	fmt.Print("Circle ")
	fmt.Println(task4.Calculator(Circle))
}

func printUser(user *task3.User, err error) {
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user.Id, user.Name)
	}
}
