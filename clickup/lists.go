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
