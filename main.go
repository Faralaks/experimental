package main

import (
	"experemental/code/super_functions"
	"html/template"
	"net/http"
	"strconv"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	res := struct {
		X string
		Y string
		Z string
	}{Z: ""}
	res.X = r.FormValue("x")
	res.Y = r.FormValue("y")
	if res.X != "" || res.Y != "" {
		x, _ := strconv.Atoi(res.X)
		y, _ := strconv.Atoi(res.Y)
		s := super_functions.SumObj{}
		res.Z = strconv.Itoa(s.Sum(x, y))
	}

	w.Header().Set("Content-Type", "text/html")
	main := "./templates/index.html"

	tmpl, err := template.ParseFiles(main)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	err = tmpl.ExecuteTemplate(w, "index", res)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

}

func main() {
	println("\n\t\t\tGo experemental Is starting")
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)

	err := http.ListenAndServe(":80", mux) // Запускаем сервер
	if err != nil {
		panic(err)
	}
}
