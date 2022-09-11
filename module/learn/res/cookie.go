package res

import (
	"fmt"
	"net/http"
)

func CookieSet(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Go Web",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "JApplication",
		HttpOnly: true,
	}
	w.Header().Set("Set-Cookie", c1.String())
	w.Header().Add("Set-Cookie", c2.String())
	c3 := http.Cookie{
		Name: "cookie3",
		Value: "cook three",
		HttpOnly: false,
	}
	http.SetCookie(w, &c3)
	fmt.Fprintln(w, r.RequestURI)
	h := r.Header["Cookie"]
	fmt.Fprintln(w, h)
	getC3, err  := r.Cookie("cookie3")
	if err != nil {
		fmt.Println(w, "can not get cookie3")
	}
	cs := r.Cookies()
	fmt.Fprintln(w, getC3)
	fmt.Fprintln(w, cs)
}
