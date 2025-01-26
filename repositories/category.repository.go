package repositories

import (
	"database/sql"
	"fmt"
	"net/http"
	"sb-go-quiz/structs"
)

func RFindCategory(db *sql.DB, id int) (result structs.SCategories, error structs.SError) {
	var category structs.SCategories
	query := "SELECT id, name, created_at, created_by, modified_at, modified_by FROM categories WHERE id = $1"

	err := db.QueryRow(query, id).Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy, &category.ModifiedAt, &category.ModifiedBy)

	if err != nil {
		if err == sql.ErrNoRows {
			return category, structs.SError{
				Message: fmt.Sprintf("no category found with ID %d", id),
				Status:  http.StatusNotFound,
			}
		}

		return category, structs.SError{
			Message: fmt.Sprintf("failed to fetch category: %s", err.Error()),
			Status:  http.StatusInternalServerError,
		}
	}

	return category, structs.SError{}
}

func RFindCategories(db *sql.DB) (result []structs.SCategories, error structs.SError) {
	query := "SELECT * FROM categories"

	rows, err := db.Query(query)

	if err != nil {
		return []structs.SCategories{}, structs.SError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	defer rows.Close()
	for rows.Next() {
		var category = structs.SCategories{}

		rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy, &category.ModifiedAt, &category.ModifiedBy)

		result = append(result, category)
	}

	return
}

func RCreateCategories(db *sql.DB, category structs.SCategories) (error structs.SError) {

	query := "INSERT INTO categories (name, created_by, modified_by) VALUES ($1, $2, $3)"

	_, err := db.Exec(query, category.Name, category.CreatedBy, category.ModifiedBy)

	if err != nil {
		return structs.SError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return
}

func RUpdateCategories(db *sql.DB, category structs.SCategories) (error structs.SError) {
	_, errFind := RFindCategory(db, category.ID)

	if errFind.Message != "" {
		return errFind
	}

	query := "UPDATE categories SET name = $1, modified_by = $2, modified_at = NOW() WHERE id = $3"

	_, err := db.Exec(query, category.Name, category.ModifiedBy, category.ID)

	if err != nil {
		return structs.SError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return
}

func RDeleteCategories(db *sql.DB, id int) (error structs.SError) {
	_, errFind := RFindCategory(db, id)

	if errFind.Message != "" {
		return errFind
	}

	query := "DELETE FROM categories WHERE id = $1"

	_, err := db.Exec(query, id)

	if err != nil {
		return structs.SError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return
}
