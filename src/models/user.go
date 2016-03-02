package models

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"
)

type Member struct {
	Username string
	Id int
	Password string
	FirstName string
}

type Session struct {
	Id int
	MemberId int
	SessionId string
}

//geters and setters

func GetMember(username string,  password string) (Member, error) {
	db, err := getDBConnection()
	
	if err == nil {
		defer db.Close()
		pwd := sha256.Sum256([]byte(password))
		
		row := db.QueryRow(`SELECT id, email, first_name
			FROM Member
			WHERE email = $1 AND password = $2`, username, hex.EncodeToString(pwd[:])) 
		
		result := Member{}
		err=row.Scan(&result.Id, &result.Username, &result.FirstName)
		
		if err == nil {
			return result, nil
		}else{
			return result, errors.New("Unable to find Member with that Username: "+err.Error())
		}
	}else{
		return Member{}, errors.New("Unable to get database connection")
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
		}else{
			return Session{}, errors.New("Unable to save session into the database")
		}
	}else{
		return result, errors.New("Unable to get a database connection to save the session")
	}
}
