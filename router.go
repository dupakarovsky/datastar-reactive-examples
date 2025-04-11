package main

import "net/http"

// router returns a http.Handler for the server to acts as the server's default Handler
func (app *application) router() http.Handler {

	mux := http.NewServeMux()

	//= STATIC FILE SERVER ==========================================
	static := http.FileServer(http.Dir("public"))
	mux.Handle("/public/", http.StripPrefix("/public", static))

	//#=============================================================================
	//# START PAGE
	//#=============================================================================
	mux.HandleFunc("GET /{$}", app.homePageHandler)

	//#=============================================================================
	//# POOLING SIGNALS
	//#=============================================================================
	// Handlers for the /pooling example
	mux.HandleFunc("GET /pooling", app.poolingPage)
	mux.HandleFunc("GET /pooling/data", app.pooling)

	//#=============================================================================
	//# IGNORE SIGNALS
	//#=============================================================================
	// Renders the page for the /ignore exemaple
	mux.HandleFunc("GET /ignore", app.ignoreHandler)

	//#=============================================================================
	//# BIND KEYS
	//#=============================================================================
	// Renders the page for the /bind-keys exemaple
	mux.HandleFunc("GET /bind-keys", app.bindKeysHandler)

	//#=============================================================================
	//# TOGGEL CLASSES
	//#=============================================================================
	// Renders the page for the /toggle-class exemaple
	mux.HandleFunc("GET /toggle-classes", app.toggleClassesHandler)

	//#=============================================================================
	//# ON LOAD
	//#=============================================================================
	// Renders the page for the /onload example
	mux.HandleFunc("GET /onload", app.onLoadHandler)
	// route handler for /onload/data
	mux.HandleFunc("GET /onload/data", app.onLoadDataHandler)

	//#=============================================================================
	//# MODEL BIND
	//#=============================================================================
	// Renders the page for the /model-bind example
	mux.HandleFunc("GET /model-bind", app.modelBindHandler)

	//#=============================================================================
	//# DISABLE BUTTON
	//#=============================================================================
	// Renders the page for the /disable-button example
	mux.HandleFunc("GET /disable-button", app.disableButtonHandler)
	// route handler for the /disable-buton example
	mux.HandleFunc("GET /disable-button/data", app.disableButtonDataHandler)

	return app.sessionManager.LoadAndSave(mux)
}
