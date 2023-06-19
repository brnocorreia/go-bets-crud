package request

type BetRequest struct {
	Nickname    string  `json:"nickname" binding:"required,min=4,max=30"`
	Date        string  `json:"date" binding:"required"`
	Bookmaker   string  `json:"bookmaker" binding:"required,min=3,max=30"`
	Sport       string  `json:"sport" binding:"required,min=2,max=40"`
	Description string  `json:"description" binding:"required,min=4,max=100"`
	Odd         float32 `json:"odd" binding:"required,gt=1"`
	Investment  float32 `json:"investment" binding:"required"`
	Tipster     string  `json:"tipster" binding:"omitempty,min=2,max=30"`
	Winner      string  `json:"winner" binding:"required"`
}
