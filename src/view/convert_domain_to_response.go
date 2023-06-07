package view

import (
	"github.com/brnocorreia/go-movies-crud/src/controller/model/response"
	"github.com/brnocorreia/go-movies-crud/src/model"
)

func convertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
