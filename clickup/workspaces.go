package clickup

import "context"

type WorkspacesService service

type Workspace struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Color   string   `json:"color"`
	Avatar  string   `json:"avatar"`
	Members []Member `json:"members"`
}

type Member struct {
	User struct {
		ID             string `json:"id"`
		Username       string `json:"username"`
		Email          string `json:"email"`
		Color          string `json:"color"`
		ProfilePicture string `json:"profilePicture"`
		Initials       string `json:"initials"`
		Role           string `json:"role"`
		//CustomRole string???
		LastActive  string `json:"last_active"`
		DateJoined  string `json:"date_joined"`
		DateInvited string `json:"date_invited"`
	} `json:"user"`
	InvitedBy struct {
		ID             string `json:"id"`
		Username       string `json:"username"`
		Color          string `json:"color"`
		Email          string `json:"email"`
		Intials        string `json:"intials"`
		ProfilePicture string `json:"profilePicture"`
	} `json:"invited_by"`
}

func (s *WorkspacesService) Get(ctx context.Context) (*Workspace, *Response, error) {
	req, err := s.client.NewRequest("GET", "team", nil)
	if err != nil {
		return nil, nil, err
	}

	wResp := new(Workspace)
	resp, err := s.client.Do(ctx, req, wResp)
	if err != nil {
		return nil, resp, err
	}

	return wResp, resp, nil
}
