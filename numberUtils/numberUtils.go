package numberUtils

import "strconv"


func ToInt64(param string) (int64, error) {
	return strconv.ParseInt(param, 10, 64)
}


func ToInt(param string)(int, error) {
	int64Value, error := ToInt64(param)
	if error != nil {
		return 0,error
	}
	return int(int64Value), nil
}