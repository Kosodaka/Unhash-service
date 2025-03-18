package v1

import "unhashService/entity"

// DTO для полчения хеша из другого сервиса
type UnhashRequest struct {
	Hash   []entity.Hash `json:"hash"`
	Domain string        `json:"domain"`
}

// DTO для ответа в ручке /unhash
type UnhashResponse struct {
	PhoneNumbers []string `json:"phone_numbers"`
}

type HashRequest struct {
	Hash   []entity.Hash `json:"hash"`
	Domain string        `json:"domain"`
}

type HashResponse struct {
	Hashes []string `json:"hashes"`
}
