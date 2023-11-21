package middleware

import (
	"finalProject3/database"
	"finalProject3/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		userID, err := strconv.Atoi(c.Param("userID"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}

		user := entity.User{}

		err = db.Select("id").First(&user, uint(userID)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "data doesnt exist",
			})
			return
		}

	}
}
