package clickup

import (
	"context"
	"fmt"
)

type GroupsService service

type GroupsWrapper struct {
	Groups []Group `json:"groups"`
}

type Group struct {
	ID          string `json:"id"`
	TeamID      string `json:"team_id"`
	UserID      int64  `json:"userid"`
	Name        string `json:"name"`
	Handle      string `json:"handle"`
	DateCreated string `json:"date_created"`
	Initials    string `json:"initials"`
	Members     []struct {
		ID             int64  `json:"id"`
		Username       string `json:"username"`
		Email          string `json:"email"`
		Color          string `json:"color"`
		Initials       string `json:"initials"`
		ProfilePicture string `json:"profilePicture"`
	} `json:"members"`
	Avatar struct {
		AttachmentID string `json:"attachment_id"`
		Color        string `json:"color"`
		Source       string `json:"source"`
		Icon         string `json:"icon"`
	} `json:"avatar"`
}

func (s *GroupsService) Get(ctx context.Context, query string) (*GroupsWrapper, *Response, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("group%s", query), nil)
	if err != nil {
		return nil, nil, err
	}

	//fmt.Println(fmt.Sprintf("%+v", req))

	wResp := new(GroupsWrapper)
	resp, err := s.client.Do(ctx, req, wResp)
	if err != nil {
		return nil, resp, err
	}

	return wResp, resp, nil

}
