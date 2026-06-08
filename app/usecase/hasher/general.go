package hasher

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"unhashService/entity"
	"unhashService/pkg/logger"
)

type UseCase struct {
	log        *logger.Logger
	HMACSecret string
}

func New(log *logger.Logger, HMACSecret string) *UseCase {
	return &UseCase{
		log:        log,
		HMACSecret: HMACSecret,
	}
}

var _ HasherUC = (*UseCase)(nil)

// Вспомогательная функция для генерации HMAC
func generateHMAC(phoneNumber, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(phoneNumber))
	return hex.EncodeToString(h.Sum(nil))
}

// Вспомогательная функция для создания строки перед хешированием
func prepareDataForHashing(phoneNumber, secret string) string {
	hmacHash := generateHMAC(phoneNumber, secret)
	return fmt.Sprintf("%s:%s", phoneNumber, hmacHash)
}

// HashPhoneNumber принимает структуру []entity.Hash и возвращает слайс "хэшей" номера телефона.
func (uc *UseCase) HashPhoneNumber(hashes []entity.PhoneNumber, domain int64) ([]string, error) {

	hashedData := make([]string, 0, len(hashes))

	for _, h := range hashes {
		h.PhoneNumber = prepareDataForHashing(h.PhoneNumber, uc.HMACSecret)
		hashBytes := []byte(h.PhoneNumber)

		hashedBytes := make([]byte, len(hashBytes))
		for i, b := range hashBytes {
			hashedBytes[i] = b ^ byte(h.Salt) ^ byte(domain)
		}

		hashedHex := hex.EncodeToString(hashedBytes)

		hashedData = append(hashedData, hashedHex)
	}

	return hashedData, nil
}

// UnhashPhoneNumber принимает структуру []entity.Hash,
// выполняет обратную операцию XOR и возвращает слайс исходных номеров телефонов.
func (uc *UseCase) UnhashPhoneNumber(hashes []entity.Hash, domain int64) ([]entity.UserHash, error) {

	var unhashedData []entity.UserHash
	for _, h := range hashes {
		hashedBytes, err := hex.DecodeString(h.PhoneNumber)
		if err != nil {
			uc.log.Info("error decoding hex while unhashing: " + err.Error())
			return nil, err
		}

		hashBytes := make([]byte, len(hashedBytes))
		for i, b := range hashedBytes {
			hashBytes[i] = b ^ byte(h.Salt) ^ byte(domain)
		}

		unhashedData = append(unhashedData, entity.UserHash{
			UserID: h.UserID,
			Hashes: string(hashBytes),
		})

	}

	return unhashedData, nil
}
