package base63

import (
	"errors"
	"math"
	"strings"
)

//TODO: extract into config file and parse it;
const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"


type Base63 struct {
	Alphabet string
	Length uint64
}

func NewBase63Handler() Base63 {
	return Base63{
		Alphabet: alphabet,
		Length: uint64(len(alphabet)),
	}
}

func (b63 Base63)Encode(randomGeneratedInteger uint64) (string, error){
	if randomGeneratedInteger == 0 {
		return "", errors.New("generated number for encoding can't be equal with 0")
	}
	var stringBuilder strings.Builder

	for ; randomGeneratedInteger > 0 && stringBuilder.Len() != 10; randomGeneratedInteger = randomGeneratedInteger / b63.Length {
		stringBuilder.WriteByte(alphabet[(randomGeneratedInteger % b63.Length)])
	}

	return stringBuilder.String(), nil
}

func (b63 Base63) Decode(shortLink string) (uint64, error) {
	var number uint64

	for i, symbol := range shortLink {
		alphabeticPosition := strings.IndexRune(alphabet, symbol)
		if alphabeticPosition == -1 {
			return uint64(alphabeticPosition), errors.New("invalid character: " + string(symbol))
		}
		number += uint64(alphabeticPosition) * uint64(math.Pow(float64(b63.Length), float64(i)))
	}

	return number, nil
}