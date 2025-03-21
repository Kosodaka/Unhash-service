package entity

type PhoneNumber struct {
	Salt        int64  `json:"salt"`
	PhoneNumber string `json:"phone_number"`
}

type Hash struct {
	PhoneNumber string `json:"phone_number"`
	Salt        int64  `json:"salt"`
}
