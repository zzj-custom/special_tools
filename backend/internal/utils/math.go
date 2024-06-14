package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// GenerateRandomNumber 生成随机数字
func GenerateRandomNumber(length int) int {
	rand.NewSource(time.Now().UnixNano())
	var str string
	for i := 0; i < length; i++ {
		str += fmt.Sprint(rand.Intn(10))
	}

	i, _ := strconv.Atoi(str)
	return i
}
