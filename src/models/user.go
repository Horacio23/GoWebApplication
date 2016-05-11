package models

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"
)

type Member struct {
	Username  string
	Id        int
	Password  string
	FirstName string
}

type Session struct {
	Id        int
	MemberId  int
	SessionId string
}

//geters and setters

func GetMember(username string, password string) (Member, error) {
	db, err := getDBConnection()

	if err == nil {
		defer db.Close()
		pwd := sha256.Sum256([]byte(password))

		row := db.QueryRow(`SELECT id, email, first_name
			FROM Member
			WHERE email = $1 AND password = $2`, username, hex.EncodeToString(pwd[:]))

		result := Member{}
		err = row.Scan(&result.Id, &result.Username, &result.FirstName)

		if err == nil {
			return result, nil
		} else {
			return result, errors.New("Unable to find Member with that Username: " + err.Error())
		}
	} else {
		return Member{}, errors.New("Unable to get database connection")
	}

}

func InsertMember(firstName string, email string, password string) error {
	db, err := getDBConnection()

	if err == nil {
		defer db.Close()
		pwd := sha256.Sum256([]byte(password))

		db.QueryRow(`INSERT INTO member( id, email, password, first_name) VALUES (DEFAULT, $1, $2, $3);`, email, hex.EncodeToString(pwd[:]), firstName)

		return nil
	} else {
		return errors.New("Couldn't connect to the database")
	}
}

func GetMemberById(id int) (Member, error) {
	db, err := getDBConnection()
	result := Member{}

	if err == nil {
		defer db.Close()

		row := db.QueryRow(`SELECT id, email, first_name FROM member WHERE id = $1;`, id)

		err = row.Scan(&result.Id, &result.Username, &result.FirstName)

		if err == nil {
			return result, nil
		} else {
			return result, errors.New("Unable to find Member with info provided")
		}
	} else {
		return result, errors.New("Couldn't connect to the database")
	}
}

func CreateSession(member Member) (Session, error) {
	result := Session{}
	result.MemberId = member.Id
	sessionId := sha256.Sum256([]byte(member.Username + time.Now().Format("12:00:00")))
	result.SessionId = hex.EncodeToString(sessionId[:])

	db, err := getDBConnection()
	if err == nil {
		defer db.Close()
		err := db.QueryRow(`INSERT INTO Session
			(member_id, session_id)
			VALUES ($1, $2)
			RETURNING id`, member.Id, result.SessionId).Scan(&result.Id)
		if err == nil {
			return result, nil
		} else {
			return Session{}, errors.New("Unable to save session into the database")
		}
	} else {
		return result, errors.New("Unable to get a database connection to save the session")
	}
}

func RemoveSession(member_id string) bool {
	db, err := getDBConnection()

	if err == nil {
		defer db.Close()

		nerr := db.QueryRow(`DELETE FROM session WHERE member_id = $1;`, member_id)

		if nerr == nil {
			return true
		}
		return true
	}

	return false
}
