package parser

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func StringToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func StringToInt16(s string) (int16, error) {
	i, err := strconv.ParseInt(s, 10, 16)
	return int16(i), err
}

func StringToBool(s string) (bool, error) {
	return strconv.ParseBool(s)
}

func StringToFloat32(s string) (float32, error) {
	f, err := strconv.ParseFloat(s, 32)
	return float32(f), err
}

func StringToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func StringToRat(s string) (*big.Rat, error) {
	rat, isErr := new(big.Rat).SetString(s)
	if !isErr {
		return nil, fmt.Errorf("failed to parse big.Rat")
	}
	return rat, nil
}

func StringToTime(s string) (time.Time, error) {
	return time.Parse(time.RFC3339, s)
}

func StringToJSON(s string) (json.RawMessage, error) {
	return json.RawMessage(s), nil
}

func StringToUUID(s string) (uuid.UUID, error) {
	return uuid.Parse(s)
}

func StringToBytes(s string) ([]byte, error) {
	return []byte(s), nil
}

func StringToIPNet(s string) (net.IPNet, error) {
	_, ipnet, err := net.ParseCIDR(s)
	return *ipnet, err
}

func StringToIP(s string) (net.IP, error) {
	return net.ParseIP(s), nil
}

func StringToHardwareAddr(s string) (net.HardwareAddr, error) {
	return net.ParseMAC(s)
}

func StringToInterfaceSlice(s string) ([]interface{}, error) {
	var i []interface{}
	err := json.Unmarshal([]byte(s), &i)
	return i, err
}

func StringToStringMap(s string) (map[string]string, error) {
	var m map[string]string
	err := json.Unmarshal([]byte(s), &m)
	return m, err
}
