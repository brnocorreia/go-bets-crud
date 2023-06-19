package response

type BetResponse struct {
	ID          string  `json:"id"`
	Nickname    string  `json:"nickname"`
	Date        string  `json:"date"`
	Bookmaker   string  `json:"bookmaker"`
	Sport       string  `json:"sport"`
	Description string  `json:"description"`
	Odd         float32 `json:"odd"`
	Investment  float32 `json:"investment"`
	Returned    float32 `json:"returned"`
	Profit      float32 `json:"profit"`
	Tipster     string  `json:"tipster"`
	Winner      string  `json:"winner"`
}
