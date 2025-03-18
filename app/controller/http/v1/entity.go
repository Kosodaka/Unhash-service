package v1

import "unhashService/entity"

// DTO для полчения хеша из другого сервиса для его обработки и расхеширования.
type UnhashRequest struct {
	Hash   []entity.Hash `json:"hash"`
	Domain string        `json:"domain"`
}

// DTO для ответа в ручке /unhash.
type UnhashResponse struct {
	PhoneNumbers []string `json:"phone_numbers"`
}

// DTO для полчения исхлдного номера из другого сервиса.
type HashRequest struct {
	Hash   []entity.PhoneNumber `json:"hash"`
	Domain string               `json:"domain"`
}

// DTO для ответа в ручке /hash.
type HashResponse struct {
	Hashes []string `json:"hashes"`
}
