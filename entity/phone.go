package entity

type PhoneNumber struct {
	Domain int64
	Hash   []Hash
}

type Hash struct {
	PhoneNumber string `json:"phone_number"`
	Salt        int64  `json:"salt"`
}
