package model

type BetDomainInterface interface {
	GetDate() string
	GetBookmaker() string
	GetSport() string
	GetDescription() string
	GetOdd() float32
	GetInvestment() float32
	GetReturned() float32
	GetProfit() float32
	GetTipster() string
	GetNickname() string
	GetID() string
	GetWinner() bool

	SetID(string)
	SetReturned(returned float32)
}

func NewBetDomain(nickname, date, bookmaker, sport, description string,
	odd float32, investment float32, tipster string, winner bool) BetDomainInterface {
	return &betDomain{
		nickname:    nickname,
		date:        date,
		bookmaker:   bookmaker,
		sport:       sport,
		description: description,
		odd:         odd,
		investment:  investment,
		tipster:     tipster,
		winner:      winner,
	}
}

func NewBetUpdateDomain(date, bookmaker, sport, description, tipster string,
	odd float32, investment, returned, profit float32) BetDomainInterface {
	return &betDomain{
		date:        date,
		bookmaker:   bookmaker,
		sport:       sport,
		description: description,
		odd:         odd,
		investment:  investment,
		returned:    returned,
		profit:      profit,
		tipster:     tipster,
	}
}

type betDomain struct {
	id          string
	nickname    string
	date        string
	bookmaker   string
	sport       string
	description string
	odd         float32
	investment  float32
	returned    float32
	profit      float32
	tipster     string
	winner      bool
}

func (bd *betDomain) GetWinner() bool {
	return bd.winner
}

func (bd *betDomain) SetReturned(returned float32) {
	winner := bd.winner
	if winner {
		odd := bd.odd
		investment := bd.investment
		returned = odd * investment
	} else {
		returned = 0
	}
	bd.returned = returned
}

func (bd *betDomain) GetNickname() string {
	return bd.nickname
}

func (bd *betDomain) GetDate() string {
	return bd.date
}

func (bd *betDomain) SetID(id string) {
	bd.id = id
}

func (bd *betDomain) GetID() string {
	return bd.id
}

func (bd *betDomain) GetBookmaker() string {
	return bd.bookmaker
}

func (bd *betDomain) GetSport() string {
	return bd.sport
}

func (bd *betDomain) GetDescription() string {
	return bd.description
}

func (bd *betDomain) GetOdd() float32 {
	return bd.odd
}

func (bd *betDomain) GetInvestment() float32 {
	return bd.investment
}

func (bd *betDomain) GetReturned() float32 {

	return bd.returned
}

func (bd *betDomain) GetProfit() float32 {
	returned := bd.returned
	investment := bd.investment
	profit := returned - investment
	bd.profit = profit
	return bd.profit
}

func (bd *betDomain) GetTipster() string {
	return bd.tipster
}
