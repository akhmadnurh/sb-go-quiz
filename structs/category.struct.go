package structs

type SCategories struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	CreatedAt  string `json:"created_at"`
	CreatedBy  int    `json:"created_by"`
	ModifiedAt string `json:"modified_at"`
	ModifiedBy int    `json:"modified_by"`
}
