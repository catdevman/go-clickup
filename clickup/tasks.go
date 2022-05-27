package clickup

import (
	"context"
	"fmt"
)

type TasksService service

type TasksWrapper struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	ID          string `json:"id"`
	CustomID    string `json:"custom_id"`
	Name        string `json:"name"`
	TextContent string `json:"text_content"`
	Description string `json:"description"`
	Status      struct {
		Status     string `json:"status"`
		Type       string `json:"type"`
		OrderIndex int64  `json:"orderindex"`
		Color      string `json:"color"`
	} `json:"status"`
	OrderIndex  string `json:"orderindex"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
	DateClosed  string `json:"date_closed"`
	Creator     struct {
		ID             int64  `json:"id"`
		Username       string `json:"username"`
		Color          string `json:"color"`
		ProfilePicture string `json:"profilePicture"`
	} `json:"creator"`
	Assignees  []interface{} `json:"assignees"`
	Checklists []interface{} `json:"checklists"`
	Tags       []struct {
		Creator         int64  `json:"creator"`
		Name            string `json:"name"`
		BackgroundColor string `json:"tag_bg"`
		ForegroundColor string `json:"tag_fg"`
	} `json:"tags"`
	Parent       interface{} `json:"parent"`
	Priority     interface{} `json:"priority"`
	DueDate      string      `json:"due_date"`
	StartDate    string      `json:"start_date"`
	TimeEstimate interface{} `json:"time_estimate"`
	TimeSpent    interface{} `json:"time_spent"`
	CustomFields []struct {
		ID             string      `json:"id"`
		Name           string      `json:"name"`
		Type           string      `json:"type"`
		TypeConfig     interface{} `json:"type_config"`
		DateCreated    string      `json:"date_created"`
		HideFromGuests bool        `json:"hide_from_guests"`
		Value          interface{} `json:"value"`
		Required       bool        `json:"required"`
	} `json:"custom_fields"`
	List struct {
		ID string `json:"id"`
	} `json:"list"`
	Folder struct {
		ID string `json:"id"`
	} `json:"folder"`
	Space struct {
		ID string `json:"id"`
	} `json:"space"`
	Url string `json:"url"`
}

func (s *TasksService) Get(ctx context.Context, taskID string, query string) (*Task, *Response, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("task/%s%s", taskID, query), nil)
	if err != nil {
		return nil, nil, err
	}

	//fmt.Println(fmt.Sprintf("%+v", req))

	wResp := new(Task)
	resp, err := s.client.Do(ctx, req, wResp)
	if err != nil {
		return nil, resp, err
	}

	return wResp, resp, nil
}

func (s *TasksService) List(ctx context.Context, listID string, query string) (*TasksWrapper, *Response, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("list/%s/task%s", listID, query), nil)
	if err != nil {
		return nil, nil, err
	}

	//fmt.Println(fmt.Sprintf("%+v", req))

	wResp := new(TasksWrapper)
	resp, err := s.client.Do(ctx, req, wResp)
	if err != nil {
		return nil, resp, err
	}

	return wResp, resp, nil
}

func (s *TasksService) ForTeam(ctx context.Context, teamID string, query string) (*TasksWrapper, *Response, error) {

	req, err := s.client.NewRequest("GET", fmt.Sprintf("team/%s/task%s", teamID, query), nil)
	if err != nil {
		return nil, nil, err
	}

	//fmt.Println(fmt.Sprintf("%+v", req))

	wResp := new(TasksWrapper)
	resp, err := s.client.Do(ctx, req, wResp)
	if err != nil {
		return nil, resp, err
	}

	return wResp, resp, nil
}
