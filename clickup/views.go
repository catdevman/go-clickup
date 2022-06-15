package clickup

import (
	"context"
	"fmt"
)

type ViewsService service

type ViewWrapper struct {
	View View `json:"view"`
}

type View struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Parent struct {
		ID   string `json:"id"`
		Type int64  `json:"type"`
	} `json:"parent"`
	Grouping struct {
		Field     string        `json:"field"`
		Dir       int32         `json:"dir"`
		Collapsed []interface{} `json:"collapsed"`
		Ingore    bool          `json:"ingore"`
	} `json:"grouping"`
	Divide struct {
		Field     string        `json:"field"`
		Dir       int32         `json:"dir"`
		Collapsed []interface{} `json:"collapsed"`
	} `json:"divide"`
	Sorting struct {
		Fields []interface{} `json:"fields"`
	} `json:"sorting"`
	Filters struct {
		Op         string        `json:"op"`
		Fields     []interface{} `json:"fields"`
		Search     string        `json:"search"`
		ShowClosed bool          `json:"show_closed"`
	} `json:"filters"`
	TeamSidebar struct {
		Assignees        []interface{} `json:"assignees"`
		AssignedComments bool          `json:"assigned_comments"`
		UnassignedTasks  bool          `json:"unassigned_tasks"`
	} `json:"team_sidebar"`
	Settings struct {
		ShowTaskLocations      bool  `json:"show_task_locations"`
		ShowSubtasks           int32 `json:"show_subtasks"`
		ShowSubtaskParentNames bool  `json:"show_subtask_parent_names"`
		ShowClosedSubtasks     bool  `json:"show_closed_subtasks"`
		ShowAssignees          bool  `json:"show_assignees"`
		ShowImages             bool  `json:"show_images"`
		CollapseEmptyColumns   bool  `json:"collapse_empty_columns"` //Example shows null WTF
		MeComments             bool  `json:"me_comments"`
		MeSubtasks             bool  `json:"me_subtasks"`
		MeChecklists           bool  `json:"me_checklists"`
	} `json:"settings"`
}

type ViewsWrapper struct {
	Views []View `json:"views"`
}

type ChatViewCommentsWrapper struct {
	Comments []ChatViewComment `json:"comments"`
}

type ChatViewComment struct {
	ID          string        `json:"id"`
	Comment     []interface{} `json:"comment"`
	CommentText string        `json:"comment_text"`
	User        struct {
		ID             int64  `json:"id"`
		Username       string `json:"username"`
		Initials       string `json:"initials"`
		Email          string `json:"email"`
		Color          string `json:"color"`
		ProfilePicture string `json:"profilePicture"`
	} `json:"user"`
	Resolved   bool          `json:"resolved"`
	Assignee   interface{}   `json:"assignee"`
	AssignedBy interface{}   `json:"assigned_by"`
	Reactions  []interface{} `json:"reactions"`
	Date       string        `json:"date"`
}

func (s *ViewsService) Get(ctx context.Context, viewID string, query string) (*ViewWrapper, *Response, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("view/%s%s", query), nil)
	if err != nil {
		return nil, nil, err
	}

	//fmt.Println(fmt.Sprintf("%+v", req))

	wResp := new(ViewWrapper)
	resp, err := s.client.Do(ctx, req, wResp)
	if err != nil {
		return nil, resp, err
	}

	return wResp, resp, nil
}

func (s *ViewsService) Tasks(ctx context.Context, viewID string, query string) (*TasksWrapper, *Response, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("view/%s/task%s", query), nil)
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

func (s *ViewsService) Comments(ctx context.Context, viewID string, query string) (*ChatViewCommentsWrapper, *Response, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("view/%s/comment%s", query), nil)
	if err != nil {
		return nil, nil, err
	}

	//fmt.Println(fmt.Sprintf("%+v", req))

	wResp := new(ChatViewCommentsWrapper)
	resp, err := s.client.Do(ctx, req, wResp)
	if err != nil {
		return nil, resp, err
	}

	return wResp, resp, nil
}
