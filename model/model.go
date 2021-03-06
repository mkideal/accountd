package model

import (
	"crypto/sha256"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 字符串慢比较函数,字节比较次数等于 got 的长度
func slowEq(expected, got string) bool {
	l1, l2 := len(expected), len(got)
	eq := l1 == l2
	for i := 0; i < l1 && i < l2; i++ {
		if expected[i] != got[i] {
			eq = false
		}
	}
	// 这个for循环没有改变结果的实际意义,因为如果 l1 != l2, eq 已经为 false
	// 这个for循环的目的在于保证字符串的字节比较次数等于必定等于 l2,即参数 got 的长度
	for i := l1; i < l2; i++ {
		if got[i] == byte(0) {
			eq = false
		}
	}
	return eq
}

func ValidateClient(client *Client, clientSecret string) bool {
	return slowEq(clientSecret, client.Secret)
}

func ValidatePassword(user *User, password string) bool {
	return slowEq(user.EncryptedPassword, EncryptPassword(password, user.PasswordSalt))
}

func EncryptPassword(password, salt string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(password+":"+salt)))
}

func JoinAccount(accountType AccountType, account string) string {
	if accountType <= AccountType_Email {
		return account
	}
	return fmt.Sprintf("%d#%s", accountType, account)
}

func SplitAccount(account string) (AccountType, string) {
	var (
		accountType = AccountType_Normal
		index       = strings.Index(account, "#")
	)
	if index > 0 {
		t, err := strconv.Atoi(account[:index])
		if err != nil {
			return accountType, ""
		}
		accountType = AccountType(t)
		account = account[index+1:]
	} else {
		if IsNormalUsername(account) {
			accountType = AccountType_Normal
		} else if IsAutoUsername(account) {
			accountType = AccountType_Auto
		} else if IsEmail(account) {
			accountType = AccountType_Email
		} else if IsTelno(account) {
			accountType = AccountType_Telno
		} else {
			account = ""
		}
	}
	return accountType, account
}

var (
	regNormalUsername = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9]{1,31}$")
	regAutoUsername   = regexp.MustCompile("^[_][a-zA-Z][a-zA-Z0-9]{1,30}$")
	regEmail          = regexp.MustCompile("^[a-zA-Z0-9_-]{1,64}@.+\\..+$")
	regTelno          = regexp.MustCompile("^[+]?[0-9]{6,15}$")
	regPassword       = regexp.MustCompile("^[a-zA-Z0-9!@#$%^&*_]{6,32}$")
)

func IsNormalUsername(account string) bool { return regNormalUsername.MatchString(account) }
func IsAutoUsername(account string) bool   { return regAutoUsername.MatchString(account) }
func IsEmail(account string) bool          { return regEmail.MatchString(account) }
func IsTelno(account string) bool          { return regTelno.MatchString(account) }
func IsPassword(password string) bool      { return regPassword.MatchString(password) }

const RFC3339Milli = "2006-01-02T15:04:05.999Z07:00"

func ParseTime(s string) (time.Time, error) {
	return time.Parse(RFC3339Milli, s)
}

func ToUnix(s string) int64 {
	t, err := ParseTime(s)
	if err != nil {
		return 0
	}
	return t.Unix()
}

func ToUnixMilli(s string) int64 {
	t, err := ParseTime(s)
	if err != nil {
		return 0
	}
	return t.UnixNano() / 1000000
}

func FormatTime(t time.Time) string {
	return t.Format(RFC3339Milli)
}

func DurationFrom(s string, from time.Time) time.Duration {
	to, err := ParseTime(s)
	if err != nil {
		return 0
	}
	d := to.Sub(from)
	if d < 0 {
		d = 0
	}
	return d
}

func DurationFromNow(s string) time.Duration {
	return DurationFrom(s, time.Now())
}

func IsExpired(s string) bool {
	t, err := ParseTime(s)
	if err != nil {
		return true
	}
	return t.Before(time.Now())
}
