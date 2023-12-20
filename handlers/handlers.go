package handlers

import (
	"encoding/base64"
	"html/template"
	rdb "main/ridership_db"
	"main/utils"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Get the selected chart from the query parameter
	selectedChart := r.URL.Query().Get("line")
	if selectedChart == "" {
		selectedChart = "red"
	}

	// instantiate ridershipDB
	var db rdb.RidershipDB = &rdb.SqliteRidershipDB{} // Sqlite implementation
	// var db rdb.RidershipDB = &rdb.CsvRidershipDB{} // CSV implementation

	// TODO: some code goes here
	// Get the chart data from RidershipDB
	db.Open("C:\\Users\\22057\\Documents\\Study\\MIT_s6583\\mbta.sqlite")
	// db.Open("C:\\Users\\22057\\Documents\\Study\\MIT_s6583\\lab\\lab0\\mbta.csv")
	chart_data, err := db.GetRidership(selectedChart)
	if err != nil {
		panic(err.Error())
	}

	// TODO: some code goes here
	// Plot the bar chart using utils.GenerateBarChart. The function will return the bar chart
	// as PNG byte slice. Convert the bytes to a base64 string, which is used to embed images in HTML.
	bar_chart, err := utils.GenerateBarChart(chart_data)
	if err != nil {
		panic(err.Error())
	}
	bar_chart = []byte(base64.StdEncoding.EncodeToString(bar_chart))

	// Get path to the HTML template for our web app
	_, currentFilePath, _, _ := runtime.Caller(0)
	templateFile := filepath.Join(filepath.Dir(currentFilePath), "template.html")

	// Read and parse the HTML so we can use it as our web app template
	html, err := os.ReadFile(templateFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl, err := template.New("line").Parse(string(html))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: some code goes here
	// We now want to create a struct to hold the values we want to embed in the HTML
	data := struct {
		Image string
		Chart string
	}{
		Image: string(bar_chart), // TODO: base64 string
		Chart: selectedChart,
	}

	// TODO: some code goes here
	// Use tmpl.Execute to generate the final HTML output and send it as a response
	// to the client's request.
	tmpl.Execute(w, data)
}
