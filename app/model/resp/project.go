package resp

type ProjectInfo struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	OwnerUser string `json:"owner_user"`
	CreatedAt string `json:"created_at"`
}
