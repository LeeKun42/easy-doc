package resp

type Api struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Directory struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Apis     []Api       `json:"apis"`
	Children []Directory `json:"children"`
}
