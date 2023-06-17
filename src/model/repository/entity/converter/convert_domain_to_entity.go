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
		Nickname:    domain.GetNickname(),
	}
}

func ConvertBetDomainToEntity(
	domain model.BetDomainInterface,
) *entity.BetEntity {
	return &entity.BetEntity{
		Nickname:    domain.GetNickname(),
		Date:        domain.GetDate(),
		Bookmaker:   domain.GetBookmaker(),
		Sport:       domain.GetSport(),
		Description: domain.GetDescription(),
		Odd:         domain.GetOdd(),
		Investment:  domain.GetInvestment(),
		Returned:    domain.GetReturned(),
		Profit:      domain.GetProfit(),
		Tipster:     domain.GetTipster(),
		Winner:      domain.GetWinner(),
	}
}
