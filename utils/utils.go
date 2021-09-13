package utils

import (
	"crypto/rand"
	"math/big"
)

func RandomUINT64() (uint64,string, error) {
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(130), nil).Sub(max, big.NewInt(1))

	//Generate cryptographically strong pseudo-random between 0 - max
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, "", err
	}
	return n.Uint64(), n.String(), nil
}
