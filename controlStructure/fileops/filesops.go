package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(balance float64, fileName string) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceText), 0644)
}

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err == nil {
		balanceText := string(data)
		balance, err := strconv.ParseFloat(balanceText, 64)
		if err != nil {
			return 1000, errors.New("failed to parse stored value")
		} else {
			return balance, nil
		}
	} else {
		return 1000, errors.New("failed to find " + fileName)
	}

}
