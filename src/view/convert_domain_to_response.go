package view

import (
	"github.com/brnocorreia/go-movies-crud/src/controller/model/response"
	"github.com/brnocorreia/go-movies-crud/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:          userDomain.GetID(),
		Email:       userDomain.GetEmail(),
		Name:        userDomain.GetName(),
		Age:         userDomain.GetAge(),
		Nationality: userDomain.GetNationality(),
		Nickname:    userDomain.GetNickname(),
	}
}

func ConvertBetDomainToResponse(
	betDomain model.BetDomainInterface,
) response.BetResponse {
	return response.BetResponse{
		ID:          betDomain.GetID(),
		Nickname:    betDomain.GetNickname(),
		Date:        betDomain.GetDate(),
		Bookmaker:   betDomain.GetBookmaker(),
		Sport:       betDomain.GetSport(),
		Description: betDomain.GetDescription(),
		Odd:         betDomain.GetOdd(),
		Investment:  betDomain.GetInvestment(),
		Returned:    betDomain.GetReturned(),
		Profit:      betDomain.GetProfit(),
		Tipster:     betDomain.GetTipster(),
		Winner:      betDomain.GetWinner(),
	}
}
