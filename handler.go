package main

import (
	"html/template"
	"net/http"
	"path"
	"strconv"
)

func IndexPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func LengthPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pages/length.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ConvertLengthHandler(w http.ResponseWriter, r *http.Request) {
	filePath := path.Join("pages", "length.html")
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func WeightPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pages/weight.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ConvertWeightHandler(w http.ResponseWriter, r *http.Request) {
	filePath := path.Join("pages", "weight.html")
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func TemperaturePageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pages/temperature.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ConvertTemperatureHandler(w http.ResponseWriter, r *http.Request) {
	filePath := path.Join("pages", "temperature.html")
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.FormValue("unit") != "" && r.FormValue("from") != "" && r.FormValue("to") != "" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		unit, err := strconv.ParseFloat(r.FormValue("unit"), 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		from := r.FormValue("from")
		to := r.FormValue("to")

		data := map[string]interface{}{
			"Unit": unit,
			"From": from,
			"To":   to,
			"Res":  ConvertTemperature(unit, from, to),
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
