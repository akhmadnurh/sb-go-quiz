package structs

type SUsers struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	CreatedAt  string `json:"created_at"`
	CreatedBy  int    `json:"created_by"`
	ModifiedAt string `json:"modified_at"`
	ModifiedBy int    `json:"modified_by"`
}
