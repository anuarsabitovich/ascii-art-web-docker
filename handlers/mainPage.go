package handlers

import (
	asciiArtFs "ascii-art-web/ascii-art-fs"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type ArtText struct {
	Text   string
	Style  string
	Result string
}

func ErrExec(r http.ResponseWriter, header int) {
	r.WriteHeader(header)
	tempErr, err := template.ParseFiles("./template/error.html")
	tempErr.ExecuteTemplate(r, "error", header)
	if err != nil {
		fmt.Println("err")
	}
	fmt.Println("check")
	return
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		if r.Method != "GET" {

			// http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			ErrExec(w, http.StatusMethodNotAllowed)
			return
		}
	}
	if r.URL.Path == "/ascii-art" {
		if r.Method != "POST" {
			// http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			ErrExec(w, http.StatusMethodNotAllowed)
			return
		}
	}

	if !(r.URL.Path == "/" || r.URL.Path == "/ascii-art") {
		ErrExec(w, http.StatusNotFound)

		// http.NotFound(w, r)

		return
	}
	temp, err := template.ParseFiles("./template/index.html")
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		return
	}

	if r.Method == "POST" {
		err = r.ParseForm()
		if err != nil {
			return
		}

		body := r.PostForm
		fmt.Println(body)
		TextContent, TextCheck := body["text"]
		_, StyleCheck := body["style"]

		fmt.Println(TextContent)
		fmt.Println(len(TextContent[0]))

		if len(TextContent[0]) < 1 {
			fmt.Println(TextContent)

			w.WriteHeader(http.StatusInternalServerError)
			ErrExec(w, http.StatusInternalServerError)

		}
		if TextCheck == false || StyleCheck == false {
			w.WriteHeader(http.StatusInternalServerError)
			ErrExec(w, http.StatusInternalServerError)

		}

	}

	var data ArtText
	data.Text = r.FormValue("text")
	data.Style = r.FormValue("style")
	data.Result, err = asciiArtFs.AsciiArtFs(data.Text, data.Style)
	if err != nil {
		switch err {
		case asciiArtFs.InvalidStyle:
			ErrExec(w, http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			return
		case asciiArtFs.InvalidInput:
			ErrExec(w, http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			return
		case asciiArtFs.BadRequestErr:
			w.WriteHeader(http.StatusBadRequest)

		}
	}

	err = temp.ExecuteTemplate(w, "index", data)
	if err != nil {
		log.Println(err)
		log.Fatalln("error in mainPage")
	}
}
