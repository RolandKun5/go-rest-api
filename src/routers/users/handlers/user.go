package handlers

import (
	"net/http"

	"github.com/RolandKun5/go-rest-api/src/database"
	"github.com/RolandKun5/go-rest-api/src/types"
	"github.com/gin-gonic/gin"
)

func getPermissionLevels(context *gin.Context) ([]types.PermissionLevel, error) {
	var permissionLevels []types.PermissionLevel

	const query = "SELECT * FROM permission_levels"

	rows, err := database.Db.Query(query)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var permissionLevel types.PermissionLevel
		if err := rows.Scan(&permissionLevel.ID, &permissionLevel.PermissionLevel); err != nil {
			context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return nil, err
		}
		permissionLevels = append(permissionLevels, permissionLevel)
	}

	if err = rows.Err(); err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return nil, err
	}

	return permissionLevels, nil

}

func GetUsers(context *gin.Context) {
	var users []types.User
	permissionLevels, err := getPermissionLevels(context)

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
		var user types.User
		for _, s := range permissionLevels {
			if s.ID == rawUser.PermissionLevelID {
				user.ID = rawUser.ID
				user.UserName = rawUser.UserName
				user.PermissionLevel = s.PermissionLevel
				user.FirstName = rawUser.FirstName
				user.LastName = rawUser.LastName
				user.Email = rawUser.Email
				user.City = rawUser.City
				user.CreatedAt = rawUser.CreatedAt
				user.UpdatedAt = rawUser.UpdatedAt
			}
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, users)

}
