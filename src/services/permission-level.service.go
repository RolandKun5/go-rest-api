package services

import (
	"net/http"

	"github.com/RolandKun5/go-rest-api/src/database"
	"github.com/RolandKun5/go-rest-api/src/types"
	"github.com/gin-gonic/gin"
)

func GetPermissionLevels(context *gin.Context) ([]types.PermissionLevel, error) {
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
