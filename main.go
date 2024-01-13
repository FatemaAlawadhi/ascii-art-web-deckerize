package main

import (
	"ascii-art-web-dockerize/Ascii"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Result string
}

type CustomError struct {
	Code    int
	Message string
}

var err error

func handleMainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		err := &CustomError{
			Code:    http.StatusNotFound,
			Message: "Page Not Found",
		}
		renderErrorPage(w, err)
		return
	}

	if r.Method == http.MethodGet {
		renderMainPage(w)
	} else {
		err := &CustomError{
			Code:    http.StatusMethodNotAllowed,
			Message: "Method Not Allowed",
		}
		renderErrorPage(w, err)
	}

}

func renderMainPage(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("template/web.html")
	if err != nil {
		err := &CustomError{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}
		renderErrorPage(w, err)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		err := &CustomError{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}
		renderErrorPage(w, err)
	}
}

func handleAsciiArt(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			err := &CustomError{
				Code:    http.StatusBadRequest,
				Message: "Bad Request",
			}
			renderErrorPage(w, err)
			return
		}

		text := r.FormValue("Text")
		banner := r.FormValue("Font")

		result, err := Ascii.AsciiArt(text, banner)
		if err != nil {
			err := &CustomError{
				Code:    http.StatusInternalServerError,
				Message: "Internal Server Error",
			}
			renderErrorPage(w, err)
			return
		}

		data := PageData{
			Result: result,
		}

		renderMainPageWithData(w, data)
	} else {
		err := &CustomError{
			Code:    http.StatusMethodNotAllowed,
			Message: "Method Not Allowed",
		}

		renderErrorPage(w, err)
	}
}

func renderMainPageWithData(w http.ResponseWriter, data PageData) {
	tmpl, err := template.ParseFiles("template/web.html")
	if err != nil {
		err := &CustomError{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}
		renderErrorPage(w, err)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		err := &CustomError{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}
		renderErrorPage(w, err)
	}
}

func renderErrorPage(w http.ResponseWriter, er *CustomError) {
	w.WriteHeader(er.Code)
	tmpl, err := template.ParseFiles("template/Error.html")
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, er)
	if err != nil {
		http.Error(w, er.Message, er.Code)
		return
	}
}

func main() {
	http.HandleFunc("/", handleMainPage)
	http.HandleFunc("/ascii-art", handleAsciiArt)
	fs := http.FileServer(http.Dir("style"))
	http.Handle("/style/", http.StripPrefix("/style/", fs))
	fmt.Println("starting server at port 8120\n")
	err = http.ListenAndServe(":8120", nil)
	if err != nil {
		log.Fatal(err)
	}
}
