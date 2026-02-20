package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func GetInput(year int, day int) (string, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	req, errReq := http.NewRequest("GET", "https://adventofcode.com/2025/day/1/input", nil)
	if errReq != nil {
		return "", errReq
	}

	errENV := godotenv.Load("../../../.env")
	if errENV != nil {
		return "", errENV
	}

	cookie, ok := os.LookupEnv("COOKIE")

	if !ok {
		return "", fmt.Errorf("no cookie")

	}

	req.Header.Add("Cookie", cookie)

	res, errRes := client.Do(req)
	if errRes != nil {
		return "", errRes
	}

	defer res.Body.Close()

	body, errB := io.ReadAll(res.Body)
	if errB != nil {
		return "", errB
	}

	return string(body), nil
}
