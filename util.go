package main

import "encoding/json"

func IsLowercaseHex(s string) bool {
	for _, ch := range s {
		if '0' <= ch && ch <= '9' || 'a' <= ch && ch <= 'f' {
			continue
		}
		return false
	}
	return true
}

func StructToString(v any) string {
	b, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(b)
}
