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
		c.IndentedJSON(400, gin.H{
			"message": "Error",
			"error":   err.Error(),
		})
		return
	}
	defer rows.Close()

	users := make([]*models.RequestUser, 0)
	for rows.Next() {
		user := new(models.RequestUser)
		err := rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			c.IndentedJSON(400, gin.H{
				"message": "Error",
				"error":   err.Error(),
			})
			return
		}
		users = append(users, user)
	}

	c.IndentedJSON(200, users)
}

// FindUser this function return user by id
func FindUser(c *gin.Context) {
	var user models.RequestUser
	userId := c.Param("id")
	intUserId, _ := strconv.ParseInt(userId, 10, 64)

	err := initializers.DB.QueryRow("SELECT * FROM users WHERE id = ?", intUserId).Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"message": "Error",
			"error":   err.Error(),
		})
		return
	}

	c.IndentedJSON(200, user)
}

// NewUser this function create a new user
func NewUser(c *gin.Context) {
	var users models.ResponseUser
	if err := c.ShouldBindJSON(&users); err != nil {
		c.IndentedJSON(400, gin.H{
			"message": "Error",
			"error":   err.Error(),
		})
		return
	}
	query := `INSERT INTO users(name, age) VALUES (?,?)`
	insertDb, err := initializers.DB.Exec(query)
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"message": "Error",
			"error":   err.Error(),
		})
		return
	}

	lastId, err := insertDb.LastInsertId()
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"message": "Error",
			"error":   err.Error(),
		})
		return
	}

	c.IndentedJSON(200, gin.H{
		"message": fmt.Sprintf("User added, id: %d", lastId),
	})
}
