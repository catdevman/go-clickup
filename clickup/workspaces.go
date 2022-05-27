package clickup

import (
	"context"
	"fmt"
)

type WorkspacesService service

type WorkspacesWrapper struct {
	Workspaces []Workspace `json:"teams"`
}

type Workspace struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Color   string   `json:"color"`
	Avatar  string   `json:"avatar"`
	Members []Member `json:"members"`
}

type Member struct {
	User struct {
		ID             int64  `json:"id"`
		Username       string `json:"username"`
		Email          string `json:"email"`
		Color          string `json:"color"`
		ProfilePicture string `json:"profilePicture"`
		Initials       string `json:"initials"`
		Role           int    `json:"role"`
		//CustomRole string???
		LastActive  string `json:"last_active"`
		DateJoined  string `json:"date_joined"`
		DateInvited string `json:"date_invited"`
	} `json:"user"`
	InvitedBy struct {
		ID             int64  `json:"id"`
		Username       string `json:"username"`
		Color          string `json:"color"`
		Email          string `json:"email"`
		Intials        string `json:"intials"`
		ProfilePicture string `json:"profilePicture"`
	} `json:"invited_by"`
}

type WorkspaceSeats struct {
	Members struct {
		FilledMemberSeats int64 `json:"filled_member_seats"`
		TotalMemberSeats  int64 `json:"total_member_seats"`
		EmptyMemberSeats  int64 `json:"empty_member_seats"`
	} `json:"members"`
	Guests struct {
		FilledGuestSeats int64 `json:"filled_guest_seats"`
		TotalGuestSeats  int64 `json:"total_guest_seats"`
		EmptyGuestSeats  int64 `json:"empty_guest_seats"`
	} `json:"guests"`
}

func (s *WorkspacesService) Get(ctx context.Context) (*WorkspacesWrapper, *Response, error) {
	req, err := s.client.NewRequest("GET", "team", nil)
	if err != nil {
		return nil, nil, err
	}

	//fmt.Println(fmt.Sprintf("%+v", req))

	wResp := new(WorkspacesWrapper)
	resp, err := s.client.Do(ctx, req, wResp)
	if err != nil {
		return nil, resp, err
	}

	return wResp, resp, nil
}

func (s *WorkspacesService) GetSeats(ctx context.Context, workspaceId string) (*WorkspaceSeats, *Response, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("team/%s/group", workspaceId), nil)
	if err != nil {
		return nil, nil, err
	}

	//	fmt.Println(fmt.Sprintf("%+v", req))

	wResp := new(WorkspaceSeats)
	resp, err := s.client.Do(ctx, req, wResp)
	if err != nil {
		return nil, resp, err
	}

	return wResp, resp, nil
}
