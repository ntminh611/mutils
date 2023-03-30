package mutils

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"
)

func String2Uint64(str string) uint64 {
	str = Remove0x(str)
	result, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Println("String2Uint64 ", err)
	}
	return uint64(result)
}

func String2Int64(str string) int64 {
	str = Remove0x(str)
	result, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Println("String2Int64 ", err)
	}
	return result
}

func String2Int(str string) int {
	str = Remove0x(str)
	result, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Println("String2Int ", err)
	}
	return int(result)
}

func String2Float64(str string) float64 {
	str = Remove0x(str)
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
	// remove 0x suffix if found in the input string
	cleaned := strings.Replace(hexStr, "0x", "", -1)

	// base 16 for hexadecimal
	result, _ := strconv.ParseInt(cleaned, 16, 64)
	return result
}

func Hex2BigInt(hexStr string) *big.Int {
	cleaned := strings.Replace(hexStr, "0x", "", -1)
	result := new(big.Int)
	result.SetString(cleaned, 16)
	return result
}

func Remove0x(hexStr string) string {
	// remove 0x suffix if found in the input string
	return strings.Replace(hexStr, "0x", "", -1)
}

func HexString2Address(hexStr string) string {
	return "0x" + hexStr[24:]
}

const zeroHex = "0000000000000000000000000000000000000000000000000000000000000000"

func ToHexString64Chars(hexStr string) string {
	cleaned := strings.Replace(hexStr, "0x", "", -1)
	return zeroHex[:64-len(cleaned)] + cleaned
}

func ToHexString64CharsSuffix(hexStr string) string {
	cleaned := strings.Replace(hexStr, "0x", "", -1)
	return cleaned + zeroHex[:64-len(cleaned)]
}

func String2BigInt(str string) *big.Int {
	result := new(big.Int)
	result, ok := result.SetString(str, 10)
	if !ok {
		fmt.Println("SetString: error")
		return nil
	}
	return result
}

func Int2HexString(n int) string {
	return fmt.Sprintf("%x", n)
}

func Int642HexString(n int64) string {
	return fmt.Sprintf("%x", n)
}

func HexString24Chars(hex string) string {
	if len(hex) >= 4 {
		return hex
	}
	return zeroHex[:4-len(hex)] + hex
}

func Int24CharsHexString(n int) string {
	return HexString24Chars(Int2HexString(n))
}

func HexString28Chars(hex string) string {
	if len(hex) >= 8 {
		return hex
	}
	return zeroHex[:8-len(hex)] + hex
}

func Int28CharsHexString(n int) string {
	return HexString28Chars(Int2HexString(n))
}

func BigInt2HexString(n *big.Int) string {
	return fmt.Sprintf("%x", n)
}

const sixZero = "000000"

func BigInt2Hex6Byte(n *big.Int) string {
	num := BigInt2HexString(n)
	return sixZero[:6-len(num)] + num
}
