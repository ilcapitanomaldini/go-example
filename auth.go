package main

import (
	"crypto/rand"
	"encoding/base32"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/securecookie"
)

// cookie handling

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func getUserName(request *http.Request) (userName string) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}

func setSession(userName string, response http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

//checking test data for now
func loginHandler(response http.ResponseWriter, request *http.Request) {
	username := request.FormValue("name")
	password := request.FormValue("password")
	if obtainedInfo, obtained := userInfo[username]; obtained && obtainedInfo.Password == password {
		setSession(username, response)
		redirectTarget := "/userPage"
		http.Redirect(response, request, redirectTarget, 302)
	} else {
		http.Error(response, "Invalid User", 401)
	}

}

func logoutHandler(response http.ResponseWriter, request *http.Request) {
	clearSession(response)
	http.Redirect(response, request, "/", 302)
}

func updatePasswordHandler(response http.ResponseWriter, request *http.Request) {
	username := getUserName(request)
	if username != "" {
		newPassword := request.FormValue("password")
		updatedInfo := userInfo[username]
		updatedInfo.Password = newPassword
		userInfo[username] = updatedInfo
		//Sending the profile info back in plain text just for demo purposes.
		js, err := json.Marshal(updatedInfo)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Header().Set("Content-Type", "application/json")
		response.Write(js)
		return
		//http.ResponseWriter()
	}
	http.Error(response, "Forbidden Action", 403)
}

//Should be sending the token to alternate means of communication and associating the generated token with the account.
//Then allowing the user to reset password. Setting token as password for demo purpose.
func forgotPasswordHandler(response http.ResponseWriter, request *http.Request) {
	username := request.FormValue("name")
	if userdata, present := userInfo[username]; present {
		newPassword := getToken(10)
		userdata.Password = newPassword
		userInfo[username] = userdata
		//TODO: Remove, debugging log only
		fmt.Println(newPassword, ": Password, data: ", userdata)
		//Sending the profile info back in plain text just for demo purposes.
		js, err := json.Marshal(userdata)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Header().Set("Content-Type", "application/json")
		response.Write(js)
	} else {
		http.Error(response, "Forbidden Action", 403)
	}
}

//Get a random token, something to replace the password with
func getToken(length int) string {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	return base32.StdEncoding.EncodeToString(randomBytes)[:length]
}
