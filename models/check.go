package models

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
)

type Check struct {
	ID            int     `json:"id" gorm:"primaryKey"`
	StoreName     string  `json:"store_name" gorm:"not null"`
	Total         float64 `json:"total" gorm:"not null"`
	PaymentMethod string  `json:"payment_method" gorm:"not null"`
	Tax           float64 `json:"tax" gorm:"default:0.00"`
}

// Структура для модели хеш-суммы чека
type CheckHash struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	CheckID int    `json:"check_id" gorm:"not null"`
	Hash    string `json:"hash" gorm:"not null"`
}

// Нужно для сортировки ключей в Json сериализации,
// так как они могут быть какждый раз по-разному
// сериализоваться, тем самым давать разный hash
type sortedMap map[string]interface{}

func (m sortedMap) MarshalJSON() ([]byte, error) {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	buf := bytes.NewBufferString("{")
	for i, k := range keys {
		v := m[k]
		jsonValue, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		if i != 0 {
			buf.WriteString(",")
		}
		buf.WriteString(fmt.Sprintf(`"%s":%s`, k, jsonValue))
	}
	buf.WriteString("}")
	return buf.Bytes(), nil
}

func CalculateCheckHash(check Check) (string, error) {
	// Serialize object with sorted keys
	checkMap := sortedMap{
		"store_name":     check.StoreName,
		"total":          check.Total,
		"payment_method": check.PaymentMethod,
		"tax":            check.Tax,
	}
	checkJSON, err := json.Marshal(checkMap)
	if err != nil {
		return "", err
	}

	// Calculate hash
	hash := md5.Sum(checkJSON)
	return hex.EncodeToString(hash[:]), nil
}
