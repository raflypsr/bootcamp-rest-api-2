package middlewares

import (
	"backend-tugas-reactjs/models"
	"backend-tugas-reactjs/utils/token"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func JwtAuthMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		var userId, _ = token.ExtractTokenID(c)

		var user models.User

		db := c.MustGet("db").(*gorm.DB)

		findUserErr := db.Model(models.User{}).Where("id = ?", userId).Take(&user).Error

		if findUserErr != nil {
			c.String(http.StatusUnauthorized, findUserErr.Error())
			c.Abort()
			return
		}

		if role == "all-user" || user.Role == role || user.Role == "admin" {
			c.Next()
		} else {
			roleError := errors.New("sorry your role cannot access this route")
			c.String(http.StatusUnauthorized, roleError.Error())
			c.Abort()
			return
		}

	}
}
