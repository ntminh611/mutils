package string

import (
	"fmt"
	"math/big"
	"strings"
)

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
