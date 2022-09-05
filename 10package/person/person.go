package person

import (
	"errors"
	"strconv"
)

// 首字母大写 可以被外部调用
func PersonInfo(name string, age int) (string, error) {
	if name == "" {
		return "", errors.New("name can not be empty")
	} else if age <= 0 || age > 100 {
		return "", errors.New("wrong age")
	}
	// 这里强转了一下 age
	return "name:" + name + ",age:" + strconv.Itoa(age), nil
}
