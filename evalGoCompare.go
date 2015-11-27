package evalGo

import (
	"fmt"
	"reflect"
	"strconv"
)

func evalBinaryExprLEQ(left interface{}, right interface{}) (interface{}, error) {
	if right == nil {
		return nil, nil
	}
	switch left.(type) {
	case int64:
		switch right.(type) {
		case int64:
			return left.(int64) <= right.(int64), nil
		case int:
			return left.(int64) <= int64(right.(int)), nil
		case float64:
			return float64(left.(int64)) <= right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return left.(int64) <= vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int64)) <= valf, nil
		case bool:
			if right.(bool) {
				return left.(int64) != int64(0), nil
			}
			return left.(int64) <= int64(0), nil
		}
	case int:
		switch right.(type) {
		case int64:
			return int64(left.(int)) <= right.(int64), nil
		case int:
			return left.(int) <= right.(int), nil
		case float64:
			return float64(left.(int)) <= right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return int64(left.(int)) <= vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int)) <= valf, nil
		case bool:
			if right.(bool) {
				return left.(int) != 0, nil
			}
			return left.(int) <= 0, nil
		}
	case float64:
		switch right.(type) {
		case int64:
			return left.(float64) <= float64(right.(int64)), nil
		case int:
			return left.(float64) <= float64(right.(int)), nil
		case float64:
			return left.(float64) <= right.(float64), nil
		case string:
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return left.(float64) <= valf, nil
		case bool:
			if right.(bool) {
				return left.(float64) != float64(0), nil
			}
			return left.(float64) <= float64(0), nil
		}
	case string:
		switch right.(type) {
		case string:
			return left.(string) <= right.(string), nil
		case int64:
			return left.(string) <= strconv.FormatInt(right.(int64), 10), nil
		case int:
			return left.(string) <= strconv.Itoa(right.(int)), nil
		case float64:
			return left.(string) <= strconv.FormatFloat(right.(float64), 'f', -1, 64), nil
		case bool:
			return left.(string) <= strconv.FormatBool(right.(bool)), nil
		}
	case nil:
		return nil, nil
	}
	return nil, fmt.Errorf("Unimplemented add for types  %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}

func evalBinaryExprNEQ(left interface{}, right interface{}) (interface{}, error) {
	val, err := evalBinaryExprEQL(left, right)
	if err != nil {
		return nil, err
	}
	return !val.(bool), nil
}

func evalBinaryExprEQL(left interface{}, right interface{}) (interface{}, error) {
	switch left.(type) {
	case int64:
		switch right.(type) {
		case int64:
			return left.(int64) == right.(int64), nil
		case int:
			return left.(int64) == int64(right.(int)), nil
		case float64:
			return float64(left.(int64)) == right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return left.(int64) == vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int64)) == valf, nil
		case bool:
			if right.(bool) {
				return left.(int64) != int64(0), nil
			}
			return left.(int64) == int64(0), nil
		case nil:
			return left.(int64) == int64(0), nil
		}
	case int:
		switch right.(type) {
		case int64:
			return int64(left.(int)) == right.(int64), nil
		case int:
			return left.(int) == right.(int), nil
		case float64:
			return float64(left.(int)) == right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return int64(left.(int)) == vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int)) == valf, nil
		case bool:
			if right.(bool) {
				return left.(int) != 0, nil
			}
			return left.(int) == 0, nil
		case nil:
			return left.(int) == 0, nil
		}
	case float64:
		switch right.(type) {
		case int64:
			return left.(float64) == float64(right.(int64)), nil
		case int:
			return left.(float64) == float64(right.(int)), nil
		case float64:
			return left.(float64) == right.(float64), nil
		case string:
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return left.(float64) == valf, nil
		case bool:
			if right.(bool) {
				return left.(float64) != float64(0), nil
			}
			return left.(float64) == float64(0), nil
		case nil:
			return left.(float64) == 0, nil
		}
	case string:
		switch right.(type) {
		case string:
			return left.(string) == right.(string), nil
		case int64:
			return left.(string) == strconv.FormatInt(right.(int64), 10), nil
		case int:
			return left.(string) == strconv.Itoa(right.(int)), nil
		case float64:
			return left.(string) == strconv.FormatFloat(right.(float64), 'f', -1, 64), nil
		case bool:
			return left.(string) == strconv.FormatBool(right.(bool)), nil
		case nil:
			return left.(string) == "", nil
		}
	case bool:
		switch right.(type) {
		case string:
			valb, err := strconv.ParseBool(right.(string))
			if err == nil {
				return left.(bool) == valb, nil
			}
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return left.(bool) == (vali != int64(0)), nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return left.(bool) == (valf != float64(0)), nil
		case int64:
			return left.(bool) == (right.(int64) != int64(0)), nil
		case int:
			return left.(bool) == (right.(int) != int(0)), nil
		case float64:
			return left.(bool) == (right.(float64) != float64(0)), nil
		case bool:
			return left.(bool) == right.(bool), nil
		case nil:
			return !left.(bool), nil
		}
	case nil:
		return nil == right, nil
	}
	return nil, fmt.Errorf("Unimplemented add for types  %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}

func evalBinaryExprLSS(left interface{}, right interface{}) (interface{}, error) {
	if right == nil {
		return nil, nil
	}
	switch left.(type) {
	case int64:
		switch right.(type) {
		case int64:
			return left.(int64) < right.(int64), nil
		case int:
			return left.(int64) < int64(right.(int)), nil
		case float64:
			return float64(left.(int64)) < right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return left.(int64) < vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int64)) < valf, nil
		case bool:
			if right.(bool) {
				return left.(int64) != int64(0), nil
			}
			return left.(int64) < int64(0), nil
		}
	case int:
		switch right.(type) {
		case int64:
			return int64(left.(int)) < right.(int64), nil
		case int:
			return left.(int) < right.(int), nil
		case float64:
			return float64(left.(int)) < right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return int64(left.(int)) < vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int)) < valf, nil
		case bool:
			if right.(bool) {
				return left.(int) != 0, nil
			}
			return left.(int) < 0, nil
		}
	case float64:
		switch right.(type) {
		case int64:
			return left.(float64) < float64(right.(int64)), nil
		case int:
			return left.(float64) < float64(right.(int)), nil
		case float64:
			return left.(float64) < right.(float64), nil
		case string:
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return left.(float64) < valf, nil
		case bool:
			if right.(bool) {
				return left.(float64) != float64(0), nil
			}
			return left.(float64) < float64(0), nil
		}
	case string:
		switch right.(type) {
		case string:
			return left.(string) < right.(string), nil
		case int64:
			return left.(string) < strconv.FormatInt(right.(int64), 10), nil
		case int:
			return left.(string) < strconv.Itoa(right.(int)), nil
		case float64:
			return left.(string) < strconv.FormatFloat(right.(float64), 'f', -1, 64), nil
		case bool:
			return left.(string) < strconv.FormatBool(right.(bool)), nil
		}
	case nil:
		return nil, nil
	}
	return nil, fmt.Errorf("Unimplemented add for types  %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}

func evalBinaryExprGTR(left interface{}, right interface{}) (interface{}, error) {
	if right == nil {
		return nil, nil
	}
	switch left.(type) {
	case int64:
		switch right.(type) {
		case int64:
			return left.(int64) > right.(int64), nil
		case int:
			return left.(int64) > int64(right.(int)), nil
		case float64:
			return float64(left.(int64)) > right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return left.(int64) > vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int64)) > valf, nil
		case bool:
			if right.(bool) {
				return left.(int64) != int64(0), nil
			}
			return left.(int64) > int64(0), nil
		}
	case int:
		switch right.(type) {
		case int64:
			return int64(left.(int)) > right.(int64), nil
		case int:
			return left.(int) > right.(int), nil
		case float64:
			return float64(left.(int)) > right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return int64(left.(int)) > vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int)) > valf, nil
		case bool:
			if right.(bool) {
				return left.(int) != 0, nil
			}
			return left.(int) > 0, nil
		}
	case float64:
		switch right.(type) {
		case int64:
			return left.(float64) > float64(right.(int64)), nil
		case int:
			return left.(float64) > float64(right.(int)), nil
		case float64:
			return left.(float64) > right.(float64), nil
		case string:
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return left.(float64) > valf, nil
		case bool:
			if right.(bool) {
				return left.(float64) != float64(0), nil
			}
			return left.(float64) > float64(0), nil
		}
	case string:
		switch right.(type) {
		case string:
			return left.(string) > right.(string), nil
		case int64:
			return left.(string) > strconv.FormatInt(right.(int64), 10), nil
		case int:
			return left.(string) > strconv.Itoa(right.(int)), nil
		case float64:
			return left.(string) > strconv.FormatFloat(right.(float64), 'f', -1, 64), nil
		case bool:
			return left.(string) > strconv.FormatBool(right.(bool)), nil
		}
	case nil:
		return nil, nil
	}
	return nil, fmt.Errorf("Unimplemented add for types  %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}

func evalBinaryExprGEQ(left interface{}, right interface{}) (interface{}, error) {
	if right == nil {
		return nil, nil
	}
	switch left.(type) {
	case int64:
		switch right.(type) {
		case int64:
			return left.(int64) >= right.(int64), nil
		case int:
			return left.(int64) >= int64(right.(int)), nil
		case float64:
			return float64(left.(int64)) >= right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return left.(int64) >= vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int64)) >= valf, nil
		case bool:
			if right.(bool) {
				return left.(int64) != int64(0), nil
			}
			return left.(int64) >= int64(0), nil
		}
	case int:
		switch right.(type) {
		case int64:
			return int64(left.(int)) >= right.(int64), nil
		case int:
			return left.(int) >= right.(int), nil
		case float64:
			return float64(left.(int)) >= right.(float64), nil
		case string:
			vali, err := strconv.ParseInt(right.(string), 10, 64)
			if err == nil {
				return int64(left.(int)) >= vali, nil
			}
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return float64(left.(int)) >= valf, nil
		case bool:
			if right.(bool) {
				return left.(int) != 0, nil
			}
			return left.(int) >= 0, nil
		}
	case float64:
		switch right.(type) {
		case int64:
			return left.(float64) >= float64(right.(int64)), nil
		case int:
			return left.(float64) >= float64(right.(int)), nil
		case float64:
			return left.(float64) >= right.(float64), nil
		case string:
			valf, err := strconv.ParseFloat(right.(string), 10)
			if err != nil {
				return nil, err
			}
			return left.(float64) >= valf, nil
		case bool:
			if right.(bool) {
				return left.(float64) != float64(0), nil
			}
			return left.(float64) >= float64(0), nil
		}
	case string:
		switch right.(type) {
		case string:
			return left.(string) >= right.(string), nil
		case int64:
			return left.(string) >= strconv.FormatInt(right.(int64), 10), nil
		case int:
			return left.(string) >= strconv.Itoa(right.(int)), nil
		case float64:
			return left.(string) >= strconv.FormatFloat(right.(float64), 'f', -1, 64), nil
		case bool:
			return left.(string) >= strconv.FormatBool(right.(bool)), nil
		}
	case nil:
		return nil, nil
	}
	return nil, fmt.Errorf("Unimplemented add for types  %s and %s", reflect.TypeOf(left), reflect.TypeOf(right))
}
