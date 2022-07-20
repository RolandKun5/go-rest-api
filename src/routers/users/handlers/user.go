package handlers

import (
	"database/sql"
	"net/http"

	"github.com/RolandKun5/go-rest-api/src/database"
	"github.com/RolandKun5/go-rest-api/src/services"
	"github.com/RolandKun5/go-rest-api/src/types"
	"github.com/RolandKun5/go-rest-api/src/utils"
	"github.com/gin-gonic/gin"
)

func GetUserById(context *gin.Context) {
	id := context.Param("userid")
	var rawUser types.RawUser

	permissionLevels, err := services.GetPermissionLevels(context)

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	const query = "SELECT * FROM users WHERE id = $1"

	if err := database.Db.QueryRow(query, id).Scan(&rawUser.ID, &rawUser.UserName, &rawUser.PermissionLevelID, &rawUser.FirstName, &rawUser.LastName,
		&rawUser.Email, &rawUser.City, &rawUser.CreatedAt, &rawUser.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var user = utils.SetPermissionLevel(rawUser, permissionLevels)

	context.IndentedJSON(http.StatusOK, user)
}

func GetUsers(context *gin.Context) {
	var users []types.User
	permissionLevels, err := services.GetPermissionLevels(context)

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	const query = "SELECT * FROM users"

	rows, err := database.Db.Query(query)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	defer rows.Close()

	for rows.Next() {
		var rawUser types.RawUser
		if err := rows.Scan(&rawUser.ID, &rawUser.UserName, &rawUser.PermissionLevelID, &rawUser.FirstName, &rawUser.LastName,
			&rawUser.Email, &rawUser.City, &rawUser.CreatedAt, &rawUser.UpdatedAt); err != nil {
			context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		var user = utils.SetPermissionLevel(rawUser, permissionLevels)
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, users)

}
