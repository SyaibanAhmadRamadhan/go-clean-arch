package utils

import (
	"crypto/rand"
	"math/big"

	"github.com/rs/zerolog/log"
)

func charsetByte(length int, charset string) ([]byte, error) {
	b := make([]byte, length)
	max := big.NewInt(int64(len(charset)))

	for i := range b {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			log.Err(err).Msg("CANNOT GENERATE RAND INT")
			return nil, err
		}
		b[i] = charset[n.Int64()]
	}

	return b, nil
}

func RandomChar(length int) (string, error) {
	randomByte, err := charsetByte(length, charset)
	if err != nil {
		return "", nil
	}

	return string(randomByte), nil
}
