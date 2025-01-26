package repositories

import (
	"database/sql"
	"fmt"
	"net/http"
	"sb-go-quiz/structs"
)

func RFindBooks(db *sql.DB) (result []structs.SBooks, error structs.SError) {
	query := "SELECT * FROM books"

	rows, err := db.Query(query)

	if err != nil {
		return []structs.SBooks{}, structs.SError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	for rows.Next() {
		var book = structs.SBooks{}

		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.CategoryID,
			&book.Description,
			&book.ImageURL,
			&book.ReleaseYear,
			&book.Price,
			&book.TotalPage,
			&book.Thickness,
			&book.CreatedAt,
			&book.CreatedBy,
			&book.ModifiedAt,
			&book.ModifiedBy,
		)

		if err != nil {
			return []structs.SBooks{}, structs.SError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}

		result = append(result, book)
	}

	return
}

func RCreateBook(db *sql.DB, book structs.SBooks) (error structs.SError) {
	query := "INSERT INTO books (title, category_id, description, image_url, release_year, price, total_page, thickness, created_by, modified_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"

	_, err := db.Exec(query, book.Title, book.CategoryID, book.Description, book.ImageURL, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CreatedBy, book.ModifiedBy)

	if err != nil {
		return structs.SError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return
}

func RFindBook(db *sql.DB, id int) (result structs.SBooks, error structs.SError) {
	query := "SELECT * FROM books WHERE id = $1"

	row := db.QueryRow(query, id)

	err := row.Scan(
		&result.ID,
		&result.Title,
		&result.CategoryID,
		&result.Description,
		&result.ImageURL,
		&result.ReleaseYear,
		&result.Price,
		&result.TotalPage,
		&result.Thickness,
		&result.CreatedAt,
		&result.CreatedBy,
		&result.ModifiedAt,
		&result.ModifiedBy,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return result, structs.SError{
				Message: fmt.Sprintf("no book found with ID %d", id),
				Status:  http.StatusNotFound,
			}
		}

		return result, structs.SError{
			Message: fmt.Sprintf("failed to fetch book: %s", err.Error()),
			Status:  http.StatusInternalServerError,
		}
	}

	return
}

func RUpdateBook(db *sql.DB, book structs.SBooks) (error structs.SError) {
	_, errFind := RFindBook(db, book.ID)

	if errFind.Message != "" {
		return errFind
	}

	query := "UPDATE books SET title = $1, category_id = $2, description = $3, image_url = $4, release_year = $5, price = $6, total_page = $7, thickness = $8, modified_by = $9, modified_at = NOW() WHERE id = $10"

	_, err := db.Exec(query, &book.Title, &book.CategoryID, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.ModifiedBy, &book.ID)

	if err != nil {
		return structs.SError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return
}

func RDeleteBook(db *sql.DB, id int) (error structs.SError) {
	_, errFind := RFindBook(db, id)

	if errFind.Message != "" {
		return errFind
	}

	query := "DELETE FROM books WHERE id = $1"

	_, err := db.Exec(query, id)

	if err != nil {
		return structs.SError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return
}

func RFindBooksByCategory(db *sql.DB, categoryID int) (result []structs.SBooks, error structs.SError) {
	query := "SELECT * FROM books WHERE category_id = $1"

	rows, err := db.Query(query, categoryID)

	if err != nil {
		if err == sql.ErrNoRows {
			return result, structs.SError{
				Message: fmt.Sprintf("no book found with Category ID %d", categoryID),
				Status:  http.StatusNotFound,
			}
		}

		return result, structs.SError{
			Message: fmt.Sprintf("failed to fetch book: %s", err.Error()),
			Status:  http.StatusInternalServerError,
		}
	}

	count := 0
	for rows.Next() {
		count++

		var book = structs.SBooks{}

		err = rows.Scan(&book.ID, &book.CategoryID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy)

		if err != nil {
			panic(err)
		}

		result = append(result, book)
	}

	if count == 0 {
		return result, structs.SError{
			Message: fmt.Sprintf("no book found with Category ID %d", categoryID),
			Status:  http.StatusNotFound,
		}
	}

	return
}
