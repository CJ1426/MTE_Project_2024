package main

import (
	f "fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("src/as"));
	http.Handle("/assets/", http.StripPrefix("/assets/", fs));
	http.HandleFunc("/", handler)
	http.ListenAndServe(":6969", nil);
}

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("IsThatACookie");
	if err == http.ErrNoCookie {
		//f.Println(cookie);
		if (r.Method == "POST" && r.URL.Path == "/varl") {
			if (r.FormValue("passd") == "test") {
				cookie := http.Cookie{
					Name: "IsThatACookie",
					Value: "token",
					Path: "/",
					MaxAge: 180,
					HttpOnly: true,
					Secure: true,
					SameSite: http.SameSiteLaxMode,
				}
				http.SetCookie(w, &cookie);
				Redirect("/").Render(r.Context(), w);
			}
			Redirect("/").Render(r.Context(), w);
		} else {
			EnterPW().Render(r.Context(), w);
		}
	} else {
		switch r.Method {
			case "GET":
				switch r.URL.Path {
					case "/":
						Index(getAllNote()).Render(r.Context(), w);
					case "/del":
						DeleteNote(r.URL.Query().Get("id"));
						Redirect("/").Render(r.Context(), w);
					default:
						notFound(w, r);		
				}
			case "POST":
				PostHandle(w,r);
			default:
				notFound(w, r);		
		}
	}
}

func PostHandle(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
		case "/add":
			AddNote(r.FormValue("note"));
			Redirect("/").Render(r.Context(), w);
		default:
			notFound(w, r);		
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	f.Println(r.URL.Path, "not found");
	w.WriteHeader(http.StatusNotFound); //ref : https://go.dev/src/net/http/status.go
	f.Fprint(w, "404 not found :(");
}
