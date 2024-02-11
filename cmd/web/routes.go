package main

import (
	"github.com/justinas/alice"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

// Update the signature for the routes() method so that it returns a
// http.Handler instead of *http.ServeMux.
func (app *application) routes() http.Handler {

	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})
		
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	router.Handler(http.MethodGet,"/static/*filepath", http.StripPrefix("/static", fileServer))

	router.HandlerFunc(http.MethodGet,"/", app.home)
	router.HandlerFunc(http.MethodGet,"/snippet/view/:id", app.snippetView)
	router.HandlerFunc(http.MethodGet,"/snippet/create", app.snippetCreate)
	router.HandlerFunc(http.MethodPost,"/snippet/create", app.snippetCreatePost)
	/// Create a middleware chain containing our 'standard' middleware
	// which will be used for every request our application receives.
	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	// Return the 'standard' middleware chain followed by the servemux.
	return standard.Then(router)

}
