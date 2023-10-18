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
	Seq            int    `json:"seq"`
	Desc           string `json:"desc"`
}

type ProjectDirectory struct {
	Name string                `json:"name"`
	Apis []ProjectDirectoryApi `json:"list"`
}

type ProjectDirectoryApi struct {
	Name        string         `json:"title"`
	Path        string         `json:"path"`
	Method      string         `json:"method"`
	Headers     []HeaderParams `json:"req_headers"`
	Query       []QueryParams  `json:"req_query"`
	ReqBodyType string         `json:"req_body_type"`
	ReqBodyForm []QueryParams  `json:"req_body_form"`
	ReqBody     string         `json:"req_body_other"`
	ResBody     string         `json:"res_body"`
	Desc        string         `json:"desc"`
}
type HeaderParams struct {
	Required int    `json:"required"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Value    string `json:"value"`
}

type QueryParams struct {
	Required int    `json:"required"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Example  string `json:"example"`
}

type ApiBodyParams struct {
	Type       string              `json:"type"`
	Required   []string            `json:"required"`
	Properties map[string]Property `json:"properties"`
}

type Property struct {
	Type       string              `json:"type"`
	Desc       string              `json:"description"`
	Properties map[string]Property `json:"properties"`
	Required   []string            `json:"required"`
	Item       Item                `json:"items"`
}

type Item struct {
	Type       string              `json:"type"`
	Desc       string              `json:"description"`
	Required   []string            `json:"required"`
	Properties map[string]Property `json:"properties"`
}

type JsonParams struct {
	Name          string       `json:"name"`
	Type          string       `json:"type"`
	SubType       string       `json:"sub_type"`
	IsPlaceholder bool         `json:"is_placeholder"`
	TestVal       string       `json:"test_val"`
	Desc          string       `json:"desc"`
	Required      bool         `json:"required"`
	Children      []JsonParams `json:"children"`
}
