package user

import (
	"math/rand"
	"os"
	"time"
)

type User struct {
	Name           string
	Password       string `json:"-"`
	Fix            string `json:""`
	WeeklyRequests WeeklyRequests
}

type WeeklyRequests struct {
	Montag     bool
	Dienstag   bool
	Mittwoch   bool
	Donnerstag bool
	Freitag    bool
}

func (wr WeeklyRequests) RequestForToday() bool {
	day := time.Now().Weekday()
	switch day {
	case time.Monday:
		return wr.Montag
	case time.Tuesday:
		return wr.Dienstag
	case time.Wednesday:
		return wr.Mittwoch
	case time.Thursday:
		return wr.Donnerstag
	case time.Friday:
		return wr.Freitag
	}
	return false
}

// GeneratePassword to hand the user for login
func GeneratePassword() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune(
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
			"abcdefghijklmnopqrstuvwxyz" +
			"0123456789")

	length := 8
	buf := make([]rune, length)
	for i := range buf {
		buf[i] = chars[rand.Intn(len(chars))]
	}
	str := string(buf)
	return str
}

func LoginAdmin(username string, password string) (ok bool) {
	adminUsername, ok := os.LookupEnv("SITOMAT_ADMIN_USERNAME")
	if !ok {
		adminUsername = "admin"
	}
	adminPassword, ok := os.LookupEnv("SITOMAT_ADMIN_PASSWORD")
	if !ok {
		adminPassword = "password"
	}
	return username == adminUsername && password == adminPassword
}
