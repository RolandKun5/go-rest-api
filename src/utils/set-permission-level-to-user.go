package utils

import "github.com/RolandKun5/go-rest-api/src/types"

func SetPermissionLevel(rawUser types.RawUser, permissionLevels []types.PermissionLevel) types.User {
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
	return user
}
