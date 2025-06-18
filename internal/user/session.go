package user

import (
	"os"
	"time"

	"github.com/Blaze5333/todo-cli/internal/storage"
	"github.com/Blaze5333/todo-cli/utils"
)

type Session struct {
	Username string    `json:"username"`
	Time     time.Time `json:"time"`
}

var sessionStorage = storage.NewStorage[Session]("session.json")

func SaveSession(username string) error {
	session := Session{Username: username, Time: time.Now()}
	return sessionStorage.Save(session)
}
func LoadSession() (*Session, error) {
	session := &Session{}
	err := sessionStorage.Load(session)
	if err != nil {
		return nil, err
	}
	return session, nil
}
func CheckSession() string {
	session, err := LoadSession()
	if err != nil {
		utils.ShowErrorMessage("Error loading session: " + err.Error())
		os.Exit(1)
	}
	if session.Username == "" {
		utils.ShowErrorMessage("No user is currently logged in.")
		os.Exit(1)
	}
	utils.ShowSuccessMessage("User " + session.Username + " is currently logged in.")
	return session.Username
}
func ClearSession() error {
	session := &Session{}
	return sessionStorage.Save(*session)
}
