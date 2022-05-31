package clickup

import (
	"context"
	"fmt"
)

type ListsService service

type ListsWrapper struct {
	Lists []List `json:"lists"`
}

type List struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Deleted    bool   `json:"deleted"`
	OrderIndex int64  `json:"orderindex"`
	Content    string `json:"content"`
	Status     struct {
		Status    string `json:"status"`
		Color     string `json:"color"`
		HideLabel bool   `json:"hide_label"`
	} `json:"status"`
	Priority struct {
		Priority string `json:"priority"`
		Color    string `json:"color"`
	} `json:"priority"`
	TaskCount int64  `json:"task_count"`
	DueDate   string `json:"due_date"`
	StartDate string `json:"start_date"`
	Folder    struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Access bool   `json:"access"`
	} `json:"folder"`
	Space struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Access bool   `json:"access"`
	} `json:"space"`
	InboundAddress   string `json:"inbound_address"`
	Archived         bool   `json:"archived"`
	OverrideStatuses bool   `json:"override_statuses"`
	Statuses         []struct {
		ID         string `json:"id"`
		Status     string `json:"status"`
		OrderIndex int64  `json:"order_index"`
		Color      string `json:"color"`
		Type       string `json:"type"`
	} `json:"statuses"`
	PermissionLevel string `json:"permission_level"`
}

type ListMembersWrapper struct {
	Members []ListMember `json:"members"`
}

type ListMember struct {
	ID             int64  `json:"id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	Color          string `json:"color"`
	Initials       string `json:"initials"`
	ProfilePicture string `json:"profilePicture"`
	ProfileInfo    struct {
		DisplayProfile           bool `json:"display_profile"`
		VerifiedAmbassador       bool `json:"verified_ambassador"`
		VerifiedConsultant       bool `json:"verified_consultant"`
		TopTierUser              bool `json:"top_tier_user"`
		ViewedVerifiedEmbassador bool `json:"viewed_verified_embassador"`
		ViewedVerifiedConsultant bool `json:"viewed_verified_consultant"`
		ViewedTopTierUser        bool `json:"viewed_top_tier_user"`
	} `json:"profileInfo"`
}

type ListCommentsWrapper struct {
	Comments []ListComment `json:"comments"`
}

type ListComment struct {
	ID          string              `json:"id"`
	Comment     []map[string]string `json:"comment"`
	CommentText string              `json:"comment_text"`
	User        struct {
		ID             int64  `json:"id"`
		Username       string `json:"username"`
		Initials       string `json:"initials"`
		Email          string `json:"email"`
		Color          string `json:"color"`
		ProfilePicture string `json:"profilePicture"`
	} `json:"user"`
	Resolved bool `json:"resolved"`
	Assignee struct {
		ID             int64  `json:"id"`
		Username       string `json:"username"`
		Initials       string `json:"initials"`
		Email          string `json:"email"`
		Color          string `json:"color"`
		ProfilePicture string `json:"profilePicture"`
	} `json:"assignee"`
	AssignedBy struct {
		ID             int64  `json:"id"`
		Username       string `json:"username"`
		Initials       string `json:"initials"`
		Email          string `json:"email"`
		Color          string `json:"color"`
		ProfilePicture string `json:"profilePicture"`
	} `json:"assigned_by"`
	Reactions []interface{} `json:"reactions"`
	Date      string        `json:"date"`
}

func (s *ListsService) Get(ctx context.Context, listID string, query string) (*List, *Response, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("list/%s", listID), nil)
	if err != nil {
		return nil, nil, err
	}

	//fmt.Println(fmt.Sprintf("%+v", req))

	wResp := new(List)
	resp, err := s.client.Do(ctx, req, wResp)
	if err != nil {
		return nil, resp, err
	}

	return wResp, resp, nil
}

func (s *ListsService) GetFolderLists(ctx context.Context, folderID string, query string) (*ListsWrapper, *Response, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("folder/%s/list", folderID), nil)
	if err != nil {
		return nil, nil, err
	}

	//fmt.Println(fmt.Sprintf("%+v", req))

	wResp := new(ListsWrapper)
	resp, err := s.client.Do(ctx, req, wResp)
	if err != nil {
		return nil, resp, err
	}

	return wResp, resp, nil
}

func (s *ListsService) GetFolderlessLists(ctx context.Context, spaceID string, query string) (*ListsWrapper, *Response, error) {

	req, err := s.client.NewRequest("GET", fmt.Sprintf("space/%s/list", spaceID), nil)
	if err != nil {
		return nil, nil, err
	}

	//fmt.Println(fmt.Sprintf("%+v", req))

	wResp := new(ListsWrapper)
	resp, err := s.client.Do(ctx, req, wResp)
	if err != nil {
		return nil, resp, err
	}

	return wResp, resp, nil
}

func (s *ListsService) Members(ctx context.Context, listID string, query string) (*ListMembersWrapper, *Response, error) {

	req, err := s.client.NewRequest("GET", fmt.Sprintf("list/%s/member%s", listID, query), nil)
	if err != nil {
		return nil, nil, err
	}

	//fmt.Println(fmt.Sprintf("%+v", req))

	wResp := new(ListMembersWrapper)
	resp, err := s.client.Do(ctx, req, wResp)
	if err != nil {
		return nil, resp, err
	}

	return wResp, resp, nil
}

func (s *ListsService) Comments(ctx context.Context, listID string, query string) (*ListCommentsWrapper, *Response, error) {

	req, err := s.client.NewRequest("GET", fmt.Sprintf("list/%s/comment%s", listID, query), nil)
	if err != nil {
		return nil, nil, err
	}

	//fmt.Println(fmt.Sprintf("%+v", req))

	wResp := new(ListCommentsWrapper)
	resp, err := s.client.Do(ctx, req, wResp)
	if err != nil {
		return nil, resp, err
	}

	return wResp, resp, nil
}
