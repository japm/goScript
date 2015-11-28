package evalGo

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func evalBinaryExprADD(left interface{}, right interface{}) (interface{}, error) {
	switch left.(type) {
	case int64:
		switch right.(type) {
		case int64:
			return left.(int64) + right.(int64), nil
		case int:
			return left.(int64) + int64(right.(int)), nil
		case float64:
			return float64(left.(int64)) + right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return left.(int64) + vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int64)) + valf, nil
		case bool:
			if right.(bool) {
				return left.(int64) + int64(1), nil
			}
			return left, nil
		case nil:
			return left.(int64), nil
		}
	case int:
		switch right.(type) {
		case int64:
			return int64(left.(int)) + right.(int64), nil
		case int:
			return left.(int) + right.(int), nil
		case float64:
			return float64(left.(int)) + right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return int64(left.(int)) + vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int)) + valf, nil
		case bool:
			if right.(bool) {
				return left.(int) + 1, nil
			}
			return left, nil
		case nil:
			return left.(int), nil
		}
	case float64:
		switch right.(type) {
		case int64:
			return left.(float64) + float64(right.(int64)), nil
		case int:
			return left.(float64) + float64(right.(int)), nil
		case float64:
			return left.(float64) + right.(float64), nil
		case string:
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return left.(float64) + valf, nil
		case bool:
			if right.(bool) {
				return left.(float64) + float64(1), nil
			}
			return left, nil
		case nil:
			return left.(float64), nil
		}
	case string:
		switch right.(type) {
		case string:
			return left.(string) + right.(string), nil
		case int64:
			return left.(string) + strconv.FormatInt(right.(int64), 10), nil
		case int:
			return left.(string) + strconv.Itoa(right.(int)), nil
		case float64:
			return left.(string) + strconv.FormatFloat(right.(float64), 'f', -1, 64), nil
		case bool:
			return left.(string) + strconv.FormatBool(right.(bool)), nil
		case nil:
			return left.(string), nil
		}
	case nil:
		return nil, nil
	}

	return nil, fmt.Errorf("Unimplemented add for types  %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}

func evalBinaryExprSUB(left interface{}, right interface{}) (interface{}, error) {
	switch left.(type) {
	case int64:
		switch right.(type) {
		case int64:
			return left.(int64) - right.(int64), nil
		case int:
			return left.(int64) - int64(right.(int)), nil
		case float64:
			return float64(left.(int64)) - right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return left.(int64) - vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int64)) - valf, nil
		case bool:
			if right.(bool) {
				return left.(int64) - int64(1), nil
			}
			return left, nil
		case nil:
			return left, nil
		}
	case int:
		switch right.(type) {
		case int64:
			return int64(left.(int)) - right.(int64), nil
		case int:
			return left.(int) - right.(int), nil
		case float64:
			return float64(left.(int)) - right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return int64(left.(int)) - vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int)) - valf, nil
		case bool:
			if right.(bool) {
				return left.(int) - 1, nil
			}
			return left, nil
		case nil:
			return left, nil
		}
	case float64:
		switch right.(type) {
		case int64:
			return left.(float64) - float64(right.(int64)), nil
		case int:
			return left.(float64) - float64(right.(int)), nil
		case float64:
			return left.(float64) - right.(float64), nil
		case string:
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return left.(float64) - valf, nil
		case bool:
			if right.(bool) {
				return left.(float64) - float64(1), nil
			}
			return left, nil
		case nil:
			return left, nil
		}
	case string:
		switch right.(type) {
		case string:
			return strings.Replace(left.(string), right.(string), "", -1), nil
		case int64:
			return strings.Replace(left.(string), strconv.FormatInt(right.(int64), 10), "", -1), nil
		case int:
			return strings.Replace(left.(string), strconv.Itoa(right.(int)), "", -1), nil
		case float64:
			return strings.Replace(left.(string), strconv.FormatFloat(right.(float64), 'f', -1, 64), "", -1), nil
		case bool:
			return strings.Replace(left.(string), strconv.FormatBool(right.(bool)), "", -1), nil
		case nil:
			return left, nil
		}
	case nil:
		return nil, nil
	}

	return nil, fmt.Errorf("Unimplemented sub for types  %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}

func evalBinaryExprMUL(left interface{}, right interface{}) (interface{}, error) {
	switch left.(type) {
	case int64:
		switch right.(type) {
		case int64:
			return left.(int64) * right.(int64), nil
		case int:
			return left.(int64) * int64(right.(int)), nil
		case float64:
			return float64(left.(int64)) * right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return left.(int64) * vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int64)) * valf, nil
		case bool:
			if right.(bool) {
				return left, nil
			}
			return int64(0), nil
		case nil:
			return int64(0), nil
		}
	case int:
		switch right.(type) {
		case int64:
			return int64(left.(int)) * right.(int64), nil
		case int:
			return left.(int) * right.(int), nil
		case float64:
			return float64(left.(int)) * right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return int64(left.(int)) * vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int)) * valf, nil
		case bool:
			if right.(bool) {
				return left, nil
			}
			return 0, nil
		case nil:
			return 0, nil
		}
	case float64:
		switch right.(type) {
		case int64:
			return left.(float64) * float64(right.(int64)), nil
		case int:
			return left.(float64) * float64(right.(int)), nil
		case float64:
			return left.(float64) * right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return left.(float64) * float64(vali), nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return left.(float64) * valf, nil
		case bool:
			if right.(bool) {
				return left, nil
			}
			return float64(0), nil
		case nil:
			return float64(0), nil
		}
	case string:
		switch right.(type) {
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err != nil {
				valf, err := strconv.ParseFloat(right.(string), 10)
				if err != nil {
					return nil, err
				}
				vali = int64(valf)
			}
			return strings.Repeat(left.(string), int(vali)), nil
		case int64:
			return strings.Repeat(left.(string), int(right.(int64))), nil
		case int:
			return strings.Repeat(left.(string), right.(int)), nil
		case float64:
			return strings.Repeat(left.(string), int(right.(float64))), nil
		case bool:
			if right.(bool) {
				return left, nil
			}
			return "", nil
		case nil:
			return "", nil
		}
	case nil:
		return nil, nil
	}
	return nil, fmt.Errorf("Unimplemented mul for types  %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}

func evalBinaryExprQUO(left interface{}, right interface{}) (interface{}, error) {
	switch left.(type) {
	case int64:
		switch right.(type) {
		case int64:
			return left.(int64) / right.(int64), nil
		case int:
			return left.(int64) / int64(right.(int)), nil
		case float64:
			return float64(left.(int64)) / right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return left.(int64) / vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int64)) / valf, nil
		case bool:
			if right.(bool) {
				return left, nil
			}
			return nil, fmt.Errorf("Divide by false no allowed")
		case nil:
			return nil, fmt.Errorf("Divide by <nil> no allowed")
		}
	case int:
		switch right.(type) {
		case int64:
			return int64(left.(int)) / right.(int64), nil
		case int:
			return left.(int) / right.(int), nil
		case float64:
			return float64(left.(int)) / right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return int64(left.(int)) / vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int)) / valf, nil
		case bool:
			if right.(bool) {
				return left, nil
			}
			return nil, fmt.Errorf("Divide by false no allowed")
		case nil:
			return nil, fmt.Errorf("Divide by <nil> no allowed")
		}
	case float64:
		switch right.(type) {
		case int64:
			return left.(float64) / float64(right.(int64)), nil
		case int:
			return left.(float64) / float64(right.(int)), nil
		case float64:
			return left.(float64) / right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return left.(float64) / float64(vali), nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return left.(float64) / valf, nil
		case bool:
			if right.(bool) {
				return left, nil
			}
			return nil, fmt.Errorf("Divide by false no allowed")
		case nil:
			return nil, fmt.Errorf("Divide by <nil> no allowed")
		}
	case nil:
		return nil, nil
	}
	return nil, fmt.Errorf("Unimplemented quo for types  %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}

func evalBinaryExprREM(left interface{}, right interface{}) (interface{}, error) {
	switch left.(type) {
	case int64:
		switch right.(type) {
		case int64:
			return left.(int64) % right.(int64), nil
		case int:
			return left.(int64) % int64(right.(int)), nil
		case float64:
			return left.(int64) % int64(right.(float64)), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return left.(int64) % vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return left.(int64) % int64(valf), nil
		case bool:
			if right.(bool) {
				return int64(0), nil
			}
			return nil, fmt.Errorf("Mod by false no allowed")
		case nil:
			return nil, fmt.Errorf("Mod by <nil> no allowed")
		}
	case int:
		switch right.(type) {
		case int64:
			return int64(left.(int)) % right.(int64), nil
		case int:
			return left.(int) % right.(int), nil
		case float64:
			return int64(left.(int)) % int64(right.(float64)), nil
		case bool:
			if right.(bool) {
				return 0, nil
			}
			return nil, fmt.Errorf("Mod by false no allowed")
		case nil:
			return nil, fmt.Errorf("Mod by <nil> no allowed")
		}
	case float64:
		switch right.(type) {
		case int64:
			return int64(left.(float64)) % right.(int64), nil
		case int:
			return int64(left.(float64)) % int64(right.(int)), nil
		case float64:
			return int64(left.(float64)) % int64(right.(float64)), nil
		case bool:
			if right.(bool) {
				return int64(0), nil
			}
			return nil, fmt.Errorf("Mod by false no allowed")
		case nil:
			return nil, fmt.Errorf("Mod by <nil> no allowed")
		}
	case nil:
		return nil, nil
	}
	return nil, fmt.Errorf("Unimplemented rem for types  %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}
