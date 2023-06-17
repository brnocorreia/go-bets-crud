package response

import "time"

type UserResponse struct {
	ID          string    `json:"id"`
	Email       string    `json:"email"`
	Nickname    string    `json:"nickname"`
	Name        string    `json:"name"`
	Age         int8      `json:"age"`
	Nationality string    `json:"nationality"`
	CreatedAt   time.Time `json:"created_at"`
}
