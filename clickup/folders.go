package clickup

import (
	"context"
	"fmt"
)

type FoldersService service

type FoldersWrapper struct {
	Folders []Folder `json:"folders"`
}

type Folder struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	OrderIndex       int64  `json:"orderindex"`
	OverrideStatuses bool   `json:"override_statuses"`
	Hidden           bool   `json:"hidden"`
	Space            struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"space"`
	TaskCount string      `json:"task_count"`
	Archived  bool        `json:"archived"`
	Statues   interface{} `json:"statues"`
	Lists     []struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		OrderIndex int64  `json:"orderindex"`
		Status     struct {
			Status    string `json:"status"`
			Color     string `json:"color"`
			HideLabel bool   `json:"hide_label"`
		} `json:"status"`
		Priority  interface{} `json:"priority"`
		Assignee  interface{} `json:"assignee"`
		TaskCount int64       `json:"task_count"`
		DueDate   string      `json:"due_date"`
		StartDate string      `json:"start_date"`
		Space     struct {
			ID     string `json:"id"`
			Name   string `json:"name"`
			Access bool   `json:"access"`
		} `json:"space"`
		Archived         bool `json:"archived"`
		OverrideStatuses bool `json:"override_statuses"`
		Statuses         []struct {
			ID         string `json:"id"`
			Status     string `json:"status"`
			OrderIndex int64  `json:"orderindex"`
			Color      string `json:"color"`
			Type       string `json:"type"`
		} `json:"statuses"`
		PermissionLevel string `json:"permission_level"`
	} `json:"lists"`
}

func (s *FoldersService) Get(ctx context.Context, folderID string, query string) (*Folder, *Response, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("folder/%s", folderID), nil)
	if err != nil {
		return nil, nil, err
	}

	//fmt.Println(fmt.Sprintf("%+v", req))

	wResp := new(Folder)
	resp, err := s.client.Do(ctx, req, wResp)
	if err != nil {
		return nil, resp, err
	}

	return wResp, resp, nil
}

func (s *FoldersService) List(ctx context.Context, spaceID string, query string) (*FoldersWrapper, *Response, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("space/%s/folder", spaceID), nil)
	if err != nil {
		return nil, nil, err
	}

	//fmt.Println(fmt.Sprintf("%+v", req))

	wResp := new(FoldersWrapper)
	resp, err := s.client.Do(ctx, req, wResp)
	if err != nil {
		return nil, resp, err
	}

	return wResp, resp, nil
}

func (s *FoldersService) Views(ctx context.Context, folderID string, query string) (*ViewsWrapper, *Response, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("folder/%s/view%s", folderID, query), nil)
	if err != nil {
		return nil, nil, err
	}

	//fmt.Println(fmt.Sprintf("%+v", req))

	wResp := new(ViewsWrapper)
	resp, err := s.client.Do(ctx, req, wResp)
	if err != nil {
		return nil, resp, err
	}

	return wResp, resp, nil
}
