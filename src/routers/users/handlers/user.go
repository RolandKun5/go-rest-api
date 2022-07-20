package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/RolandKun5/go-rest-api/src/database"
	"github.com/RolandKun5/go-rest-api/src/services"
	"github.com/RolandKun5/go-rest-api/src/types"
	"github.com/RolandKun5/go-rest-api/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getUserById(id string, permissionLevels []types.PermissionLevel, context *gin.Context) {
	var rawUser types.RawUser

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

	var user = utils.SetRawUserToUser(rawUser, permissionLevels)

	context.IndentedJSON(http.StatusOK, user)
}

func GetUserById(context *gin.Context) {
	id := context.Param("userid")
	permissionLevels, _ := services.GetPermissionLevels(context)
	getUserById(id, permissionLevels, context)
}

func GetUsers(context *gin.Context) {
	var users []types.User
	permissionLevels, _ := services.GetPermissionLevels(context)

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
		var user = utils.SetRawUserToUser(rawUser, permissionLevels)
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, users)

}

func CreateNewUser(context *gin.Context) {
	uuid := uuid.New()
	var newUser types.User
	newUser.ID = uuid.String()
	permissionLevels, _ := services.GetPermissionLevels(context)

	if err := context.BindJSON(&newUser); err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	const query = "INSERT INTO users(id,user_name,permission_level_id,first_name,last_name,email,city) values($1,$2,$3,$4,$5,$6,$7)"

	var rawUser = utils.SetUserToRawUser(newUser, permissionLevels)

	_, err := database.Db.Exec(query, rawUser.ID, rawUser.UserName, rawUser.PermissionLevelID, rawUser.FirstName, rawUser.LastName, rawUser.Email, rawUser.City)

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	getUserById(uuid.String(), permissionLevels, context)

}

func UpdateUser(context *gin.Context) {
	var user types.User
	permissionLevels, _ := services.GetPermissionLevels(context)
	currentTime, _ := time.Parse(time.RFC3339, "2016-02-02T15:04:05.000Z")
	if err := context.BindJSON(&user); err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var rawUser = utils.SetUserToRawUser(user, permissionLevels)

	const query = "UPDATE users SET user_name=$1, permission_level_id=$2, first_name=$3, last_name=$4, email=$5, city=$6, updated_at=$7 WHERE id = $8"

	_, err := database.Db.Exec(query, rawUser.UserName, rawUser.PermissionLevelID, rawUser.FirstName, rawUser.LastName, rawUser.Email, rawUser.City, currentTime.Format(time.RFC3339), rawUser.ID)

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	getUserById(rawUser.ID, permissionLevels, context)

}
