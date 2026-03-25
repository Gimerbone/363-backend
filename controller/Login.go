package controller

import (
	"363project/controller/service"
	"fmt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Ambil ID dri middleware
	userIdCtx := r.Context().Value("id")

	if userIdCtx == "" || userIdCtx == nil {
		user, err := service.CreateAnonymousUser()
		if err != nil {
			http.Error(w, "Gagal registrasi user baru", http.StatusInternalServerError)
			return
		}

		// Set cookie ke browser
		cookie := &http.Cookie{
			Name:     "id",
			Value:    fmt.Sprint(user.Id),
			Path:     "/",
			HttpOnly: true,
			MaxAge:   86400,
		}
		http.SetCookie(w, cookie)

		fmt.Fprintf(w, "Login Berhasil! Selamat datang user baru (ID: %d)", user.Id)
	} else {
		fmt.Fprintf(w, "Welcome Back! Anda login sebagai User ID: %v", userIdCtx)
	}
}
