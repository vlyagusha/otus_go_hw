package hw10programoptimization

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	return GetDomainStatNew(r, domain)
}

func GetDomainStatNew(r io.Reader, domain string) (DomainStat, error) {
	result := make(DomainStat)
	dotDomain := "." + strings.ToLower(domain)

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Bytes()
		if !bytes.Contains(line, []byte(domain)) {
			continue
		}

		var user User
		if err := jsoniter.Unmarshal(line, &user); err != nil {
			return nil, err
		}

		userDomain := strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])
		if !strings.HasSuffix(userDomain, dotDomain) {
			continue
		}

		result[userDomain]++
	}

	return result, nil
}

func GetDomainStatOld(r io.Reader, domain string) (DomainStat, error) {
	u, err := getUsers(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return countDomains(u, domain)
}

type users [100_000]User

func getUsers(r io.Reader) (result users, err error) {
	content, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		var user User
		if err = json.Unmarshal([]byte(line), &user); err != nil {
			return
		}
		result[i] = user
	}
	return
}

func countDomains(u users, domain string) (DomainStat, error) {
	result := make(DomainStat)

	for _, user := range u {
		matched, err := regexp.Match("\\."+domain, []byte(user.Email))
		if err != nil {
			return nil, err
		}

		if matched {
			num := result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])]
			num++
			result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])] = num
		}
	}
	return result, nil
}
