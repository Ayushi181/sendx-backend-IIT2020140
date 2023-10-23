package handlers

import "net/http"

func AdminLoginHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./html/admin.html")
}
