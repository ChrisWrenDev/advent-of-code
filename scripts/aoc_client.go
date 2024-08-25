package scripts

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const baseURL = "https://adventofcode.com"

func getAocSession() (string, error) {
	file, err := os.Open(".session")
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text(), scanner.Err()
}

func aocRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	cookie, err := getAocSession()

	if err != nil {
		return nil, err
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: cookie,
	})

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Sprintf("failed to retrieve prompt: status code %v", resp.StatusCode)
		return nil, errors.New(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if strings.HasPrefix(string(body), "Please don't repeatedly") {
		return nil, errors.New("AOC repeated request error")
	}

	return body, nil
}

func GetAocInput(day, year int) []byte {
	url := fmt.Sprintf("%s/%d/day/%d", baseURL, year, day)

	body, err := aocRequest(url)

	if err != nil {
		fmt.Println("Error getting input: ", err)
	}

	return body
}

func GetAocPrompt(day, year int) []byte {
	url := fmt.Sprintf("%s/%d/day/%d", baseURL, year, day)

	body, err := aocRequest(url)

	if err != nil {
		fmt.Println("Error getting prompt: ", err)
	}

	return body
}
