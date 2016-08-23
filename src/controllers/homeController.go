package controllers

import (
	"controllers/util"
	"fmt"
	"models"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"
	"viewmodels"
)

type homeController struct {
	template       *template.Template
	loginTemplate  *template.Template
	signupTemplate *template.Template
}

func (this *homeController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetHome()
	println("Getting Home controller")

	w.Header().Add("Content-Type", "text/html")
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	_, err := req.Cookie("sessionId")
	if err == nil {
		// get the member cookie and set the displayed name to the member
		if cookie, err := req.Cookie("user"); err == nil {

			vm.User = cookie.Value

		} else {
			fmt.Println("Error retrieving the member cookie:", err.Error())
		}

	} else {
		// if there is no session cookie then redirect to login
		http.Redirect(responseWriter, req, "/login", http.StatusFound)
	}

	this.template.Execute(responseWriter, vm)

}

func (this *homeController) login(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	responseWriter.Header().Add("Content-Type", "text/html")
	vm := viewmodels.GetLogin()
	if req.Method == "POST" {
		username := req.FormValue("username")
		password := req.FormValue("password")

		member, err := models.GetMember(username, password)

		if err == nil {
			sCookie, mCookie, cErr := setupSession(member)
			if cErr == nil {
				responseWriter.Header().Add("Set-Cookie", sCookie.String())
				responseWriter.Header().Add("Set-Cookie", mCookie.String())
			}

			http.Redirect(responseWriter, req, "/home", http.StatusFound)
		} else {
			fmt.Println("Unable to get member in Login:", err.Error())
			vm.Error = "Invalid Username or Password. Please try again"
		}

	}

	this.loginTemplate.Execute(responseWriter, vm)
}

func (this *homeController) signup(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	responseWriter.Header().Add("Content-Type", "text/html")
	vm := viewmodels.GetSignup()

	if req.Method == "POST" {
		firstName := req.FormValue("name")
		username := req.FormValue("username")
		password := req.FormValue("password")

		err := models.InsertMember(firstName, username, password)

		if err == nil {

			member, mErr := models.GetMember(username, password)

			if mErr == nil {
				sCookie, mCookie, cErr := setupSession(member)
				if cErr == nil {
					responseWriter.Header().Add("Set-Cookie", sCookie.String())
					responseWriter.Header().Add("Set-Cookie", mCookie.String())
				}

				http.Redirect(responseWriter, req, "/home", http.StatusFound)
			} else {
				fmt.Println("There was a problem in Signup:", mErr.Error())
				vm.Error = "There was an error, please try again"
			}
		} else {
			fmt.Println("There was an error creating the Member:", err.Error())
		}
	}

	this.signupTemplate.Execute(responseWriter, vm)
}

func (this *homeController) logout(w http.ResponseWriter, req *http.Request) {

	ck, err := req.Cookie("sessionId")

	if err == nil {
		removedS := models.RemoveSession(ck.Value)
		if removedS {
			cookieMonster := &http.Cookie{
				Name:    "sessionId",
				Expires: time.Now(),
				Value:   strconv.FormatInt(time.Now().Unix(), 10),
			}
			cookieMonster2 := &http.Cookie{
				Name:    "user",
				Expires: time.Now(),
				Value:   strconv.FormatInt(time.Now().Unix(), 10),
			}

			http.SetCookie(w, cookieMonster)
			http.SetCookie(w, cookieMonster2)

			http.Redirect(w, req, "/login", http.StatusFound)
		}
	}
}

func setupSession(member models.Member) (sCookie http.Cookie, mCookie http.Cookie, retErr error) {

	session, err := models.CreateSession(member)
	if err == nil {
		sCookie.Name = "sessionId"
		sCookie.Value = session.SessionId
		sCookie.Expires = time.Now().Add(1 * time.Hour)

		mCookie.Name = "user"
		mCookie.Value = strings.Title(member.FirstName)
		mCookie.Expires = time.Now().Add(1 * time.Hour)

	} else {
		retErr = err
	}

	return

}
