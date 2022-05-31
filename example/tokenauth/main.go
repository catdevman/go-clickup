package main

import (
	"context"
	"fmt"
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

	client := clickup.NewClient(pk.Client())

	l, _, err := client.Groups.Get(ctx, "")
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	fmt.Println(fmt.Sprintf("%+v", l))
}
