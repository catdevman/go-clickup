package clickup

import (
	"context"
	"fmt"
)

type SpacesService service

type SpacesWrapper struct {
	Spaces []Space `json:"spaces"`
}

type Space struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Private        bool   `json:"private"`
	Color          string `json:"color"`
	Avatar         string `json:"avatar"`
	AdminCanManage bool   `json:"admin_can_manage"`
	Statues        []struct {
		ID         string `json:"id"`
		Status     string `json:"status"`
		Type       string `json:"type"`
		OrderIndex int64  `json:"order_index"`
		Color      string `json:"color"`
	} `json:"statues"`
	MultipleAssignees bool                   `json:"multiple_assignees"`
	Features          map[string]interface{} `json:"features"` // Most are bools but priorities also had more detail so I'll need a better way of marshaling this
	Archived          bool                   `json:"archived"`
	Members           []struct {
		User struct {
			ID             int64  `json:"id"`
			Username       string `json:"username"`
			Color          string `json:"color"`
			ProfilePicture string `json:"profilePicture"`
		} `json:"user"`
	} `json:"members"`
}

func (s *SpacesService) Get(ctx context.Context, spaceID string, query string) (*Space, *Response, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("space/%s", spaceID), nil)
	if err != nil {
		return nil, nil, err
	}

	//fmt.Println(fmt.Sprintf("%+v", req))

	wResp := new(Space)
	resp, err := s.client.Do(ctx, req, wResp)
	if err != nil {
		return nil, resp, err
	}

	return wResp, resp, nil
}

func (s *SpacesService) List(ctx context.Context, workspaceID string, query string) (*SpacesWrapper, *Response, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("team/%s/space", workspaceID), nil)
	if err != nil {
		return nil, nil, err
	}

	//fmt.Println(fmt.Sprintf("%+v", req))

	wResp := new(SpacesWrapper)
	resp, err := s.client.Do(ctx, req, wResp)
	if err != nil {
		return nil, resp, err
	}

	return wResp, resp, nil
}
