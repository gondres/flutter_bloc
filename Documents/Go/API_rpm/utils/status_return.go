package utils

import "api_rpm/model"

func StatusUnauthorized() model.ErrorStatus {
	errorUnathorized := model.ErrorStatus{
		StatusCode: 401,
		Message:    "not authorized",
	}

	return errorUnathorized
}

func UserNotRegistered() model.ErrorStatus {
	userNotRegistered := model.ErrorStatus{
		StatusCode: 201,
		Message:    "user not yet registered",
	}

	return userNotRegistered
}
func TokenEmpty() model.ErrorStatus {
	errorTokenEmpty := model.ErrorStatus{
		StatusCode: 201,
		Message:    "token is needed",
	}

	return errorTokenEmpty
}
