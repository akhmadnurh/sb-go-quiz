package middlewares

import (
	"database/sql"
	"net/http"
	"sb-go-quiz/structs"

	"github.com/gin-gonic/gin"
)

func MAuth(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		username, password, ok := c.Request.BasicAuth()

		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		var user structs.SUsers

		query := "SELECT * FROM users WHERE username = $1 AND password = $2"
		err := db.QueryRow(query, username, password).Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt, &user.CreatedBy, &user.ModifiedAt, &user.ModifiedBy)

		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}

			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}

}
