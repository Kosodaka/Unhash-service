package entity

type PhoneNumber struct {
	UserID      int64  `json:"user_id"`
	Salt        int64  `json:"salt"`
	PhoneNumber string `json:"phone_number"`
}

type Hash struct {
	UserID      int64  `json:"user_id"`
	PhoneNumber string `json:"phone_number"`
	Salt        int64  `json:"salt"`
}

type UserHash struct {
	UserID int64  `json:"user_id"`
	Hashes string `json:"phone_number"`
}
