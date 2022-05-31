package main

import (
	"context"
	"fmt"
	"sync"
	"syscall"

	"github.com/catdevman/go-clickup/clickup"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	fmt.Print("ClickUp Token: ")
	byteToken, _ := terminal.ReadPassword(int(syscall.Stdin))
	println()
	token := string(byteToken)
	ctx := context.Background()
	pk := clickup.PersonalTokenTransport{PersonalToken: token}

	pages := make(chan int, 50)
	tasks := make(chan clickup.Task)

	go func() {
		for i := 0; i < 50; i++ {
			pages <- i
		}
		close(pages)
	}()
	wg := sync.WaitGroup{}
	client := clickup.NewClient(pk.Client())
	for p := range pages {
		wg.Add(1)
		go func(page int) {
			fmt.Println("Before API call for page", page)
			ts, _, err := client.Tasks.ForTeam(ctx, "", fmt.Sprintf("?page=%d", page))
			if err != nil {
				fmt.Printf("\nerror: %v\n", err)
				wg.Done()
				return
			}
			fmt.Println("Done with API call for page", page)
			if len(ts.Tasks) == 0 {
				fmt.Println("Page", page, "has no tasks")
				wg.Done()
				return
			}
			for _, t := range ts.Tasks {
				tasks <- t
			}
			wg.Done()
		}(p)

	}

	go func() {
		wg.Wait()
		close(tasks)
	}()

	for t := range tasks {
		fmt.Println(fmt.Sprintf("Task ID: %+v", t.ID))

	}

}
