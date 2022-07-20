package utils

import "github.com/RolandKun5/go-rest-api/src/types"

func SetRawUserToUser(rawUser types.RawUser, permissionLevels []types.PermissionLevel) types.User {
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

func SetUserToRawUser(user types.User, permissionLevels []types.PermissionLevel) types.RawUser {
	var rawUser types.RawUser
	for _, s := range permissionLevels {
		if s.PermissionLevel == user.PermissionLevel {
			rawUser.ID = user.ID
			rawUser.UserName = user.UserName
			rawUser.PermissionLevelID = s.ID
			rawUser.FirstName = user.FirstName
			rawUser.LastName = user.LastName
			rawUser.Email = user.Email
			rawUser.City = user.City
			rawUser.CreatedAt = user.CreatedAt
			rawUser.UpdatedAt = user.UpdatedAt
		}
	}
	return rawUser
}
