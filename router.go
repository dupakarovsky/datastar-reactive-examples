package main

import "net/http"

// router returns a http.Handler for the server to acts as the server's Handler
func (app *application) router() http.Handler {

	mux := http.NewServeMux()

	//= STATIC FILE SERVER ==========================================
	static := http.FileServer(http.Dir("public"))
	mux.Handle("/public/", http.StripPrefix("/public", static))

	//#=============================================================================
	//# POOLING SIGNALS
	//#=============================================================================
	mux.HandleFunc("GET /{$}", app.homePageHandler)
	mux.HandleFunc("GET /pooling", app.poolingPage)
	mux.HandleFunc("GET /pooling/data", app.pooling)

	//#=============================================================================
	//# IGNORE SIGNALS
	//#=============================================================================
	// mux.HandleFunc("GET /ingore", app.ignoreHandler)
	mux.HandleFunc("GET /ignore", app.ignoreHandler)

	//#=============================================================================
	//# BIND KEYS
	//#=============================================================================
	mux.HandleFunc("GET /bind-keys", app.bindKeysHandler)

	//#=============================================================================
	//# TOGGEL CLASSES
	//#=============================================================================
	mux.HandleFunc("GET /toggle-classes", app.toggleClassesHandler)

	//#=============================================================================
	//# ON LOAD
	//#=============================================================================
	mux.HandleFunc("GET /onload", app.onLoadHandler)
	mux.HandleFunc("GET /onload/data", app.onLoadDataHandler)

	//#=============================================================================
	//# MODEL BIND
	//#=============================================================================
	mux.HandleFunc("GET /model-bind", app.modelBindHandler)

	//#=============================================================================
	//# DISABLE BUTTON
	//#=============================================================================
	mux.HandleFunc("GET /disable-button", app.disableButtonHandler)
	mux.HandleFunc("GET /disable-button/data", app.disableButtonDataHandler)

	return app.sessionManager.LoadAndSave(mux)
}
