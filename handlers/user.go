package handlers

import "net/http"

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./html/user.html")
}
