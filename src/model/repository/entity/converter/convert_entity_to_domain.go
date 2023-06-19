package converter

import (
	"github.com/brnocorreia/go-movies-crud/src/model"
	"github.com/brnocorreia/go-movies-crud/src/model/repository/entity"
)

func ConvertEntityToDomain(
	entity entity.UserEntity,
) model.UserDomainInterface {

	domain := model.NewUserDomain(entity.Email, entity.Password, entity.Name, entity.Age, entity.Nationality,
		entity.Nickname)

	domain.SetID(entity.ID.Hex())
	return domain
}

func ConvertBetEntityToDomain(
	entity entity.BetEntity,
) model.BetDomainInterface {

	domain := model.NewBetDomain(entity.Nickname, entity.Date, entity.Bookmaker, entity.Sport, entity.Description,
		entity.Odd, entity.Investment, entity.Tipster, entity.Winner)

	domain.SetID(entity.ID.Hex())
	return domain
}
