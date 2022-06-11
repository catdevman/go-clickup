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

type CustomRolesWrapper struct {
	CustomRoles []CustomRole `json:"custom_roles"`
}

type CustomRole struct {
	ID            int64   `json:"id"`
	TeamID        string  `json:"team_id"`
	InheritedRole int64   `json:"inherited_role"`
	DateCreated   string  `json:"date_created"`
	Members       []int64 `json:"members"`
}

type TaskTemplatesWrapper struct {
	TaskTemplates []TaskTemplate `json:"templates"`
}

type TaskTemplate struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type WebhooksWrapper struct {
	Webhooks []Webhook `json:"webhooks"`
}

type Webhook struct {
	ID       string   `json:"id"`
	UserId   int64    `json:"userid"`
	TeamID   int64    `json:"team_id"`
	Endpoint string   `json:"endpoint"`
	ClientID string   `json:"client_id"`
	Events   []string `json:"events"`
	TaskID   string   `json:"task_id"`
	ListID   string   `json:"list_id"`
	SpaceID  string   `json:"space_id"`
	Health   struct {
		Status    string `json:"status"`
		FailCount int64  `json:"fail_count"`
	} `json:"health"`
	Secret string `json:"secret"`
}

type SharedHierarchy struct {
	Shared struct {
		Tasks []interface{} `json:"tasks"`
		Lists []struct {
			ID         string      `json:"id"`
			Name       string      `json:"name"`
			OrderIndex int32       `json:"orderindex"`
			Content    string      `json:"content"`
			Status     string      `json:"status"`
			Priority   string      `json:"priority"`
			Assignee   interface{} `json:"assignee"`
			TaskCount  string      `json:"task_count"`
			DueDate    string      `json:"due_date"`
			StartDate  string      `json:"start_date"`
			Archived   bool        `json:"archived"`
		} `json:"lists"`
		Folders []struct {
			ID         string `json:"id"`
			Name       string `json:"name"`
			OrderIndex int32  `json:"orderindex"`
			Content    string `json:"content"`
			TaskCount  string `json:"task_count"`
			DueDate    string `json:"due_date"`
			Archived   bool   `json:"archived"`
		} `json:"folders"`
	} `json:"shared"`
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

func (s *WorkspacesService) CustomRoles(ctx context.Context, workspaceId string, query string) (*CustomRolesWrapper, *Response, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("team/%s/customroles%s", workspaceId, query), nil)
	if err != nil {
		return nil, nil, err
	}

	//	fmt.Println(fmt.Sprintf("%+v", req))

	wResp := new(CustomRolesWrapper)
	resp, err := s.client.Do(ctx, req, wResp)
	if err != nil {
		return nil, resp, err
	}

	return wResp, resp, nil
}

func (s *WorkspacesService) TaskTemplates(ctx context.Context, workspaceId string, query string) (*TaskTemplatesWrapper, *Response, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("team/%s/taskTemplate%s", workspaceId, query), nil)
	if err != nil {
		return nil, nil, err
	}

	//	fmt.Println(fmt.Sprintf("%+v", req))

	wResp := new(TaskTemplatesWrapper)
	resp, err := s.client.Do(ctx, req, wResp)
	if err != nil {
		return nil, resp, err
	}

	return wResp, resp, nil
}

func (s *WorkspacesService) Webhooks(ctx context.Context, workspaceId string, query string) (*WebhooksWrapper, *Response, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("team/%s/webhook%s", workspaceId, query), nil)
	if err != nil {
		return nil, nil, err
	}

	//	fmt.Println(fmt.Sprintf("%+v", req))

	wResp := new(WebhooksWrapper)
	resp, err := s.client.Do(ctx, req, wResp)
	if err != nil {
		return nil, resp, err
	}

	return wResp, resp, nil
}

func (s *WorkspacesService) SharedHierarchy(ctx context.Context, workspaceId string, query string) (*SharedHierarchy, *Response, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("team/%s/shared%s", workspaceId, query), nil)
	if err != nil {
		return nil, nil, err
	}

	//	fmt.Println(fmt.Sprintf("%+v", req))

	wResp := new(SharedHierarchy)
	resp, err := s.client.Do(ctx, req, wResp)
	if err != nil {
		return nil, resp, err
	}

	return wResp, resp, nil
}

func (s *WorkspacesService) Views(ctx context.Context, workspaceID string, query string) (*ViewsWrapper, *Response, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("team/%s/view%s", workspaceID, query), nil)
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
