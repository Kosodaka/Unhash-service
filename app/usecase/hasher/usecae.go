package hasher

import (
	"unhashService/entity"
)

type HasherUC interface {
	// HashPhoneNumber принимает номер телефона (строка из 12 символов),
	// соль и домен, и возвращает "хэш" номера телефона.
	HashPhoneNumber(hash []entity.Hash, domain string) ([]string, error)

	// UnhashPhoneNumber принимает хэш, соль и домен,
	// выполняет обратную операцию XOR и возвращает исходный номер телефона.
	UnhashPhoneNumber(hash []entity.Hash, domain string) ([]string, error)
}
