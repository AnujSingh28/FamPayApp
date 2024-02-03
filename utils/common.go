package utils

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

func Paginate(data []interface{}, pageNo int64, recordsPerPage int64) []interface{} {
	startIndex := recordsPerPage * (pageNo - 1)
	endIndex := recordsPerPage * pageNo
	length := int64(len(data))
	if startIndex > length {
		startIndex = length
	}
	if endIndex > length {
		endIndex = length
	}
	data = data[startIndex:endIndex]
	return data
}

func CreateRandomSearchSlug(data string) string {
	n, _ := rand.Int(rand.Reader, big.NewInt(1000))
	randomSuffix := strconv.Itoa(int(n.Int64()))
	data += randomSuffix
	return data
}
