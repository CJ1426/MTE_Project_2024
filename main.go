package main

import (
	f "fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("src/as"));
	http.Handle("/assets/", http.StripPrefix("/assets/", fs));
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil);
}

func handler(w http.ResponseWriter, r *http.Request) {
	theCookie, err := r.Cookie("IsThatACookie");
	if err == http.ErrNoCookie {
		if (r.Method == "POST" && r.URL.Path == "/varl") {
			id := CheckAccount(r.FormValue("uname") ,r.FormValue("passd"));
			if (id != -1) {
				cookie := http.Cookie{
					Name: "IsThatACookie",
					Value: f.Sprintf("%d",id),
					Path: "/",
					MaxAge: 180,
					HttpOnly: true,
					Secure: true,
					SameSite: http.SameSiteLaxMode,
				}
				http.SetCookie(w, &cookie);
			}
			Redirect("/?err=1").Render(r.Context(), w);
		} else if (r.URL.Path == "/reg") {
			if (r.Method == "POST") {
				id := CreateUser(r.FormValue("uname") ,r.FormValue("passd"));
				if (id != -1) {
					cookie := http.Cookie{
						Name: "IsThatACookie",
						Value: f.Sprintf("%d",id),
						Path: "/",
						MaxAge: 180,
						HttpOnly: true,
						Secure: true,
						SameSite: http.SameSiteLaxMode,
					}
					http.SetCookie(w, &cookie);
					Redirect("/").Render(r.Context(), w);
				} else {
				Redirect("/reg?err=1").Render(r.Context(), w);
				}
			} else {
				isErr := false;
				if r.URL.Query().Get("err") != "" {
					isErr = true;
				}
				Regis(isErr).Render(r.Context(), w);
			}
		} else {
			isErr := false;
			if r.URL.Query().Get("err") != "" {
				isErr = true;
			}
			EnterPW(isErr).Render(r.Context(), w);
		}
	} else {
		switch r.Method {
			case "GET":
				switch r.URL.Path {
					case "/":
						Index(getAllNote(theCookie.Value)).Render(r.Context(), w);
					case "/del":
						DeleteNote(r.URL.Query().Get("id"));
						Redirect("/").Render(r.Context(), w);
					case "/edit":
						EditNote(r.URL.Query().Get("id"), GetNoteById(r.URL.Query().Get("id"), theCookie.Value)).Render(r.Context(), w);		
					case "/logout":
						cookie := http.Cookie{
							Name: "IsThatACookie",
							Value: "",
							Path: "/",
							MaxAge: -1,
							HttpOnly: true,
							Secure: true,
							SameSite: http.SameSiteLaxMode,
						}
						http.SetCookie(w, &cookie);
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
	theCookie, _ := r.Cookie("IsThatACookie");
	switch r.URL.Path {
		case "/add":
			AddNote(r.FormValue("note"), theCookie.Value);
			Redirect("/").Render(r.Context(), w);
		case "/edit":
			EditNoteQ(r.FormValue("note"), r.URL.Query().Get("id"));
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
