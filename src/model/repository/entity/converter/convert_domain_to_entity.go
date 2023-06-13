package converter

import (
	"github.com/brnocorreia/go-movies-crud/src/model"
	"github.com/brnocorreia/go-movies-crud/src/model/repository/entity"
)

func ConvertDomainToEntity(
	domain model.UserDomainInterface,
) *entity.UserEntity {
	return &entity.UserEntity{
		Email:       domain.GetEmail(),
		Password:    domain.GetPassword(),
		Name:        domain.GetName(),
		Age:         domain.GetAge(),
		Nationality: domain.GetNationality(),
		CreatedAt:   domain.GetCreatedAt(),
	}
}
