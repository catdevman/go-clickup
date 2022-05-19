package main

import (
	"context"
	"fmt"
	"syscall"

	"github.com/catdevman/go-clickup/clickup"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/oauth2"
)

func main() {
	fmt.Print("ClickUp Token: ")
	byteToken, _ := terminal.ReadPassword(int(syscall.Stdin))
	println()
	token := string(byteToken)

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client, _ := clickup.NewClient(tc)

	workspaces, resp, err := client.Workspaces.Get(ctx)
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	fmt.Println(workspaces, resp)

}
