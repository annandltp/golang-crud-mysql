package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type databerita struct {
	name    string
	message string
	judul   string
	konten  string
}

func fikom(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Novanto - 170101005 - Fakultas Ilmu Komputer - Universitas Duta Bangsa")
}

func berita(w http.ResponseWriter, r *http.Request) {
	var tampilan, kesalahan = template.ParseFiles("template/berita.html")
	if kesalahan != nil {
		fmt.Println(kesalahan.Error())
		return
	}
	var data = map[string]string{
		"judul":  "Selamat Datang",
		"konten": "Konten Berita Ada Di Sini",
	}
	tampilan.Execute(w, data)
}

func main() {
	http.HandleFunc("/", routeIndexGet)
	http.HandleFunc("/process", routeSubmitPost)

	http.HandleFunc("/fikom", fikom)
	http.HandleFunc("/berita", berita)

	fmt.Println("Web berjalan di alamat -> http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
func routeIndexGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var tmpl = template.Must(template.New("form").ParseFiles("template/view.html"))
		var err = tmpl.Execute(w, nil)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func routeSubmitPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var tmpl = template.Must(template.New("result").ParseFiles("template/view.html"))

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var nim = r.FormValue("nim")
		var nama = r.Form.Get("name")
		var jurusan = r.FormValue("jurusan")
		var pesan = r.Form.Get("message")

		var data = map[string]string{"nim": nim, "name": nama, "jurusan": jurusan, "message": pesan}

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}
