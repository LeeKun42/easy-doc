package request

type ProjectParams struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProjectUserParams struct {
	ProjectID int    `json:"project_id"`
	Account   string `json:"account"`
}

type ProjectDirectoryParams struct {
	ID        int    `json:"id"`
	ProjectID int    `json:"project_id"`
	ParentID  int    `json:"parent_id"`
	Name      string `json:"name"`
	Seq       int    `json:"seq"`
	Desc      string `json:"desc"`
}

type ProjectApiParams struct {
	ID             int    `json:"id"`
	ProjectID      int    `json:"project_id"`
	DirectoryID    int    `json:"directory_id"`
	Name           string `json:"name"`
	Path           string `json:"path"`
	Method         string `json:"method"`
	RequestHeaders string `json:"request_headers"`
	RequestQuery   string `json:"request_query"`
	RequestPath    string `json:"request_path"`
	RequestBody    string `json:"request_body"`
	ResponseBody   string `json:"response_body"`
	Desc           string `json:"desc"`
}
