package structs

type SBooks struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	CategoryID  int    `json:"category_id"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	ReleaseYear int    `json:"release_year"`
	Price       int    `json:"price"`
	TotalPage   int    `json:"total_page"`
	Thickness   string `json:"thickness"`
	CreatedAt   string `json:"created_at"`
	CreatedBy   int    `json:"created_by"`
	ModifiedAt  string `json:"modified_at"`
	ModifiedBy  int    `json:"modified_by"`
}
