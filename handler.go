package main

import (
	"html/template"
	"log/slog"
	"net/http"
	"path"
	"strconv"
	"strings"
)

func IndexPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	slog.Info("Index page is success loaded")
}

func LengthPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("view/length.html")
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	slog.Info("Length page is success loaded")
}

func WeightPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("view/weight.html")
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	slog.Info("Weight page is success loaded")
}

func TemperaturePageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("view/temperature.html")
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	slog.Info("Temperature page is success loaded")
}

func ConvertHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pageName := r.FormValue("converter") + ".html"
	filePath := path.Join("view", pageName)
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	unit := r.FormValue("unit")
	from := r.FormValue("from")
	to := r.FormValue("to")

	if unit != "" && from != "" && to != "" {
		unitParse, err := strconv.ParseFloat(r.FormValue("unit"), 64)
		if err != nil {
			slog.Error(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var res float64
		if r.FormValue("converter") == "length" {
			res = ConvertLength(unitParse, from, to)
		} else if r.FormValue("converter") == "weight" {
			res = ConvertWeight(unitParse, from, to)
		} else {
			res = ConvertTemperature(unitParse, from, to)
		}

		data := map[string]interface{}{
			"Unit": unit,
			"From": strings.ToUpper(from[:1]) + from[1:],
			"To":   strings.ToUpper(to[:1]) + to[1:],
			"Res":  res,
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			slog.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		slog.Info("Convert is success")
	} else {
		var nameInputs []string
		if unit == "" {
			nameInputs = append(nameInputs, "unit")
		}

		if from == "" {
			nameInputs = append(nameInputs, "from")
		}

		if to == "" {
			nameInputs = append(nameInputs, "to")
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			slog.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		slog.Error(strings.Join(nameInputs, ", ") + " is empty")
		http.Error(w, strings.Join(nameInputs, ", "), http.StatusBadRequest)
	}
}
