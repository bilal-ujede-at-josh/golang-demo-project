package model

type UserDetails struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	Slug             string `json:"slug"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	IsMobileVerified int    `json:"isMobileVerified"`
	IsEmailVerified  int    `json:"isEmailVerified"`
	IsGuestUser      int    `json:"isGuestUser"`
	Fname            string `json:"fname"`
	Lname            string `json:"lname"`
	Is_approved      int    `json:"is_approved"`
	Is_completed     int    `json:"is_completed"`
}
