package hasher

import (
	"encoding/hex"
	"strconv"
	"unhashService/entity"
	"unhashService/pkg/logger"
)

type UseCase struct {
	log *logger.Logger
}

func New(log *logger.Logger) *UseCase {
	return &UseCase{
		log: log,
	}
}

var _ HasherUC = (*UseCase)(nil)

// HashPhoneNumber принимает номер телефона (строка из 12 символов),
// соль и домен, и возвращает "хэш" номера телефона.
func (uc *UseCase) HashPhoneNumber(hashes []entity.Hash, domain string) ([]string, error) {
	domainInt, err := strconv.ParseInt(domain, 10, 64)
	if err != nil {
		uc.log.Info("Error parsing domain while hashing: " + err.Error())
		return nil, err
	}
	hashedData := make([]string, 0, len(hashes))

	for _, h := range hashes {
		hashBytes := []byte(h.PhoneNumber)

		hashedBytes := make([]byte, len(hashBytes))
		for i, b := range hashBytes {
			hashedBytes[i] = b ^ byte(h.Salt) ^ byte(domainInt)
		}

		hashedHex := hex.EncodeToString(hashedBytes)

		hashedData = append(hashedData, hashedHex)
	}

	return hashedData, nil
}

// UnhashPhoneNumber принимает хэш, соль и домен,
// выполняет обратную операцию XOR и возвращает исходный номер телефона.
func (uc *UseCase) UnhashPhoneNumber(hashes []entity.Hash, domain string) ([]string, error) {

	domainInt, err := strconv.ParseInt(domain, 10, 64)
	if err != nil {
		uc.log.Info("error parsing domain while unhashing: " + err.Error())
		return nil, err
	}

	unhashedData := make([]string, 0, len(hashes))

	for _, h := range hashes {
		hashedBytes, err := hex.DecodeString(h.PhoneNumber)
		if err != nil {
			uc.log.Info("error decoding hex while unhashing: " + err.Error())
			return nil, err
		}

		hashBytes := make([]byte, len(hashedBytes))
		for i, b := range hashedBytes {
			hashBytes[i] = b ^ byte(h.Salt) ^ byte(domainInt)
		}

		unhashedData = append(unhashedData, string(hashBytes))
	}

	return unhashedData, nil
}
