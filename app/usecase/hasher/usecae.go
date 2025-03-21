package hasher

import (
	"unhashService/entity"
)

type HasherUC interface {
	// HashPhoneNumber принимает структуру []entity.Hash и возвращает слайс "хэшей" номера телефона.
	HashPhoneNumber(hash []entity.PhoneNumber, domain string) ([]string, error)

	// UnhashPhoneNumber принимает структуру []entity.Hash,
	// выполняет обратную операцию XOR и возвращает слайс исходных номеров телефонов.
	UnhashPhoneNumber(hash []entity.Hash, domain string) ([]string, error)
}
