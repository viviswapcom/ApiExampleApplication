package main

// import (
// 	"fyne.io/fyne"
// 	"fyne.io/fyne/app"
// 	"fyne.io/fyne/widget"

// 	"github.com/tangleMesh/OmokuApiExampleApplication/packages/config"
// )

// func main() {
// 	configuration := config.GetConfig()
// 	// currencyPairs := http.GetCurrencyPairs()

// 	a := app.New()

// 	w := a.NewWindow(configuration.ApplicationName)
// 	w.Resize(fyne.Size{600, 800})

// 	selectEntry := widget.NewSelectEntry([]string{"1", "2", "3"})
// 	w.SetContent(selectEntry)

// 	w.ShowAndRun()

// }

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	fmt.Println("Server is listening on http://localhost:3000 or http://127.0.0.1:3000")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := string(r.URL.Path)
		if path == "/" || path == "" {
			path = "index"
		}
		tmplFile, error := template.ParseFiles("./static/" + path + ".html")
		if error != nil {
			tmplFile, error = template.ParseFiles("./static/404.html")
		}

		tmpl := template.Must(tmplFile, error)
		tmpl.Execute(w, "") // , data
	})
	http.ListenAndServe(":3000", nil)
}
