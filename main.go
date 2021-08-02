package main

import (
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
		res.Z = strconv.Itoa(Alinka(x, y))
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

func Alinka(x, y int) int {
	// Суммирует два числа и возвращает результат сложения
	// Паникует, если результат больше 100

	sum := x + y
	if sum > 100 {
		panic("AAAA! это число слишком большое!")
	}
	return sum
}

func main() {
	println("\t\t\tGo Index Page Is starting")
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)

	_ = http.ListenAndServe(":80", mux) // Запускаем сервер
}
