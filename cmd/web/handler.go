package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/machilan1/go_prc/postgres"
)

// Returns HTML file of home page.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	files := []string{"./ui/html/pages/home.html", "./ui/html/base.html", "./ui/html/partials/nav.html"}
	tmp, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
	}

	err = tmp.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}

// Returns a snippet
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {

	// !這個值是暫時的，未來會刪掉。暫時先設為 1  .
	const _SNIPPET_ID int = 1

	snp, err := app.store.SnippetStore.FindSnippet(_SNIPPET_ID)
	if err != nil {
		app.serverError(w, r, err)
	}
	fmt.Println("Congratulations! You successfully fetch a snippet 😇", snp)

	v, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || v < 1 {
		app.clientError(w, http.StatusNotFound)
		return
	}

	msg := fmt.Sprintf("Snippet (ID = %d) is found", v)
	w.Write([]byte(msg))
}

// Create a snippet via GET method
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	err := app.store.SnippetStore.Create(postgres.CreateSnippetParams{Content: "Beautiful Content"})
	if err != nil {
		fmt.Println(err)
	}
	app.logger.Debug("A snippet is created successfully.")
	w.Write([]byte(r.PathValue("create")))
}

// Create a snippet via POST method
func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	app.logger.Debug("Trying to create a snippet via POST request. But it's not yet implemented. Have a nice day!")
	w.Header().Add("happy-birthday", "tom")
	w.WriteHeader(http.StatusCreated)
}
