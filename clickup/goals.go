package clickup

import (
	"context"
	"fmt"
)

type GoalsService service

type GoalsWrapper struct {
	Goals []Goal `json:"goals"`
}

type GoalWrapper struct {
	//TODO: Tell ClickUp they need better engineers and consistent return objects
	Goal Goal
}

type Goal struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	TeamID         string `json:"team_id"`
	DateCreated    string `json:"date_created"`
	StartDate      string `json:"start_date"`
	DueDate        string `json:"due_date"`
	Description    string `json:"description"`
	Private        bool   `json:"private"`
	Archived       bool   `json:"archived"`
	Creator        int64  `json:"creator"`
	Color          string `json:"color"`
	PrettyID       string `json:"pretty_id"`
	MultipleOwners bool   `json:"multiple_owners"`
	FolderID       string `json:"folder_id"`
	Members        []struct {
		ID              int64  `json:"id"`
		Username        string `json:"username"`
		Email           string `json:"email"`
		Color           string `json:"color"`
		PermissionLevel string `json:"permission_level"`
		ProfilePicture  string `json:"profilePicture"` //TODO: Tell ClickUp to pick a lane
		Initials        string `json:"initials"`
		IsCreator       bool   `json:"isCreator"` //TODO: Tell ClickUp to pick a lane
	} `json:"members"`
	Owners []struct {
		ID             int64  `json:"id"`
		Username       string `json:"username"`
		Initials       string `json:"initials"`
		Email          string `json:"email"`
		Color          string `json:"color"`
		ProfilePicture string `json:"profilePicture"` //TODO: Tell ClickUp to pick a lane
	} `json:"owners"`
	GroupMembers     []interface{} `json:"group_members"`
	KeyResults       []interface{} `json:"key_results"`
	PercentCompleted int64         `json:"percent_completed"`
	History          []interface{} `json:"history"`
	PrettyUrl        string        `json:"pretty_url"`
}

func (s *GoalsService) List(ctx context.Context, workspaceID string, query string) (*GoalsWrapper, *Response, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("goal/%s%s", workspaceID, query), nil)
	if err != nil {
		return nil, nil, err
	}

	//fmt.Println(fmt.Sprintf("%+v", req))

	wResp := new(GoalsWrapper)
	resp, err := s.client.Do(ctx, req, wResp)
	if err != nil {
		return nil, resp, err
	}

	return wResp, resp, nil
}

func (s *GoalsService) Get(ctx context.Context, goalID string, query string) (*GoalWrapper, *Response, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("team/%s/goal%s", goalID, query), nil)
	if err != nil {
		return nil, nil, err
	}

	//fmt.Println(fmt.Sprintf("%+v", req))

	wResp := new(GoalWrapper)
	resp, err := s.client.Do(ctx, req, wResp)
	if err != nil {
		return nil, resp, err
	}

	return wResp, resp, nil
}
