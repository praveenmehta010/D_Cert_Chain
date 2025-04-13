package blockchain

import (
	"encoding/hex"
	"errors"
	"strings"
)

// HexToBytes32 converts a hex string (64 characters) to [32]byte
func HexToBytes32(hexStr string) ([32]byte, error) {
	var b32 [32]byte
	decoded, err := hex.DecodeString(strings.TrimPrefix(hexStr, "0x"))
	if err != nil || len(decoded) != 32 {
		return b32, errors.New("invalid hex string")
	}
	copy(b32[:], decoded[:32])
	return b32, nil
}
