package utils

import "strconv"

func ListContains(subList, list []string) bool {
	if len(subList) > len(list) {
		return false
	}
	for _, element := range subList {
		found := false
		for _, b := range list {
			if element == b {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func IsElementExists(list []string, element string) bool {
	for _, item := range list {
		if item == element {
			return true
		}
	}
	return false
}

func StringToInt64(str string) (int64, error) {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}
