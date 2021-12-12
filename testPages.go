package main

import (
	"fmt"
	"net/http"
)

const loginPage = `
<h1>Login Page</h1>
<br>
<form method="post" action="/login">
    <label for="name">User name</label>
    <input type="text" id="name" name="name">
    <label for="password">Password</label>
    <input type="password" id="password" name="password">
    <button type="submit">Login</button>
</form>

<form method="post" action="/forgotPassword">
    <label for="name">User name</label>
    <input type="text" id="name" name="name">
    <button type="submit">Forgot Password</button>
</form>
`

func loginPageHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, loginPage)
}

const userPage = `
<h1>User Page</h1>
<br>
<hr>
<br>
<small>User: %s</small>
<form method="post" action="/logout">
    <button type="submit">Logout</button>
</form>
<br>
<hr>
<br>
<form method="post" action="/updatePassword">
    <label for="password">Password</label>
    <input type="text" id="password" name="password">
    <button type="submit">Update Password</button>
</form>
`

func userPageHandler(response http.ResponseWriter, request *http.Request) {
	userName := getUserName(request)
	if userName != "" {
		fmt.Fprintf(response, userPage, userName)
	} else {
		http.Redirect(response, request, "/", 302)
	}
}
