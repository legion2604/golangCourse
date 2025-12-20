package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	maps := map[string]any{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	res, _ := EncodeSimpleJSON(maps)
	fmt.Println(res)
}

func EncodeSimpleJSON(x any) (string, error) {
	switch v := x.(type) {
	case string:
		return `"` + v + `"`, nil
	case int:
		return strconv.Itoa(v), nil
	case bool:
		return strconv.FormatBool(v), nil
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64), nil
	case nil:
		return "null", nil
	case []any:
		var parts []string
		for _, elem := range v {
			s, err := EncodeSimpleJSON(elem) // рекурсия
			if err != nil {
				return "", err
			}
			parts = append(parts, s)
		}
		return "[" + strings.Join(parts, ",") + "]", nil
	case map[string]any:
		var parts []string
		for key, value := range v {
			s, err := EncodeSimpleJSON(value)
			if err != nil {
				return "", err
			}
			parts = append(parts, `"`+key+`":`+s)
		}
		return "{" + strings.Join(parts, ",") + "}", nil
	default:
		return "", nil
	}
}

/*
Мистер ГПТ помог с map

*/
