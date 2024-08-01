package controller

import (
	"fmt"
	"mvc/initializers"
	"mvc/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ListUsers this function list all users and return a json
func ListUsers(c *gin.Context) {
	rows, err := initializers.DB.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	users := make([]*models.User, 0)
	for rows.Next() {
		user := new(models.User)
		err := rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		users = append(users, user)
	}

	c.IndentedJSON(200, users)
}

// FindUser this function return user by id
func FindUser(c *gin.Context) {
	var user models.User
	userId := c.Param("id")
	intUserId, _ := strconv.ParseInt(userId, 10, 64)

	err := initializers.DB.QueryRow("SELECT * FROM users WHERE id = ?", intUserId).Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		c.IndentedJSON(200, gin.H{
			"message": "Error",
			"error":   err.Error(),
		})
		return
	}

	c.IndentedJSON(200, user)
}
