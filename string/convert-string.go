package string

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"
)

func String2Uint64(str string) uint64 {
	str = strings.TrimSpace(str)
	result, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Println("String2Uint64 ", err)
	}
	return uint64(result)
}

func String2Int64(str string) int64 {
	str = strings.TrimSpace(str)
	result, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Println("String2Int64 ", err)
	}
	return result
}

func String2Int(str string) int {
	str = strings.TrimSpace(str)
	result, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Println("String2Int ", err)
	}
	return int(result)
}

func String2Float64(str string) float64 {
	str = strings.TrimSpace(str)
	result, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Println("String2Float64 ", err)
	}
	return result
}

func DecodeString(hexStr string) ([]byte, error) {
	cleaned := strings.Replace(hexStr, "0x", "", -1)
	return hex.DecodeString(cleaned)
}

func StringCompare(str1, str2 string) bool {
	str1 = strings.TrimSpace(str1)
	str2 = strings.TrimSpace(str2)
	return strings.ToLower(str1) == strings.ToLower(str2)
}

func StringContains(str1, str2 string) bool {
	return strings.Contains(strings.ToLower(str1), strings.ToLower(str2))
}

func Hex2Int64(hexStr string) int64 {
	hexStr = Remove0x(hexStr)
	result, _ := strconv.ParseInt(hexStr, 16, 64)
	return result
}

func Hex2BigInt(hexStr string) *big.Int {
	hexStr = Remove0x(hexStr)
	result := new(big.Int)
	result.SetString(hexStr, 16)
	return result
}

func Remove0x(hexStr string) string {
	return strings.Replace(hexStr, "0x", "", -1)
}

func String2BigInt(str string) *big.Int {
	str = strings.TrimSpace(str)
	result := new(big.Int)
	result, ok := result.SetString(str, 10)
	if !ok {
		fmt.Println("SetString: error")
		return nil
	}
	return result
}
