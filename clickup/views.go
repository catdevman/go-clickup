package clickup

type View struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
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
