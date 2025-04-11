package main

import (
	"datastar-reactive/internal/data"
	"datastar-reactive/ui"
	"errors"
	"net/http"
	"time"

	datastar "github.com/starfederation/datastar/sdk/go"
)

// #=============================================================================
// # HOME
// #=============================================================================
// homePageHandler renders the page when the /route is requested
func (app *application) homePageHandler(w http.ResponseWriter, r *http.Request) {
	ui.Page("Home Page").Render(r.Context(), w)
}

// #=============================================================================
// # POOLING SIGNALS
// #=============================================================================
// poolingPage will render the page for /pooling
func (app *application) poolingPage(w http.ResponseWriter, r *http.Request) {
	// when the Pooling component loads, it'll send a GET request to the /pooling/interval route
	ui.Pooling().Render(r.Context(), w)
}

// runs when ui.Pooling loads.
func (app *application) pooling(w http.ResponseWriter, r *http.Request) {
	// start defining the signals struct to receive the signals from the frontend
	signals := data.PoolingSignals{
		Count: 0,
	}

	// read the signals from the front-end into the struct
	if err := datastar.ReadSignals(r, &signals); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// update the coutner
	signals.Count++
	// app.infoLog.Println("📶 signals: ", signals)

	// start a new event
	sse := datastar.NewSSE(w, r)

	// send the udpated signal back to the front-end
	sse.MarshalAndMergeSignals(signals)

	//INFO: OPTION #2
	// use data-on-delay, then  merge a new component back to the DOM.
	// sse.MergeFragmentTempl(ui.PoolingDelay())
}

// #=============================================================================
// # IGNORE SIGNALS
// #=============================================================================
// bindKeysHandler will render the page for /ingore
func (app *application) ignoreHandler(w http.ResponseWriter, r *http.Request) {
	ui.Ignore().Render(r.Context(), w)
}

// #=============================================================================
// # BIND KEYS
// #=============================================================================
// bindKeysHandler will render the page for /bind-keys
func (app *application) bindKeysHandler(w http.ResponseWriter, r *http.Request) {
	ui.BindKeys().Render(r.Context(), w)
}

// #=============================================================================
// # TOGGLE CLASSES
// #=============================================================================
// toggleClassesHandler will render the page for /toggle-clases
func (app *application) toggleClassesHandler(w http.ResponseWriter, r *http.Request) {
	ui.ToggleClasses().Render(r.Context(), w)
}

// #=============================================================================
// # ON LOAD
// #=============================================================================
// onLoadHandler will render the page for /on-load
func (app *application) onLoadHandler(w http.ResponseWriter, r *http.Request) {
	//define a key to retrieve the value from context
	sessionKey := "datastar-on-load-example"

	// put values in the session manager
	app.sessionManager.Put(r.Context(), sessionKey, data.SessionMap{"foo": "bar", "baz": 42})

	ui.OnLoad().Render(r.Context(), w)
}

func (app *application) onLoadDataHandler(w http.ResponseWriter, r *http.Request) {
	// session key
	sessionKey := "datastar-on-load-example"

	// try retrieving a session value
	sessionValues, ok := app.sessionManager.Get(r.Context(), sessionKey).(*data.SessionMap)
	if !ok {
		err := errors.New("❌ unknown datatype from session")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// generate a new sse
	sse := datastar.NewSSE(w, r)

	// send the session data to the front end with a new component
	sse.MergeFragmentTempl(ui.OnLoadReplacer(sessionKey, sessionValues))

}

// #=============================================================================
// # MODEL BIND
// #=============================================================================
// modelBindHandler will render the page for /model-bind on first-load
func (app *application) modelBindHandler(w http.ResponseWriter, r *http.Request) {
	ui.ModelBind().Render(r.Context(), w)
}

// #=============================================================================
// # DISABLE BUTTON
// #=============================================================================
// disableButtonHandler renders the page on first load
func (app *application) disableButtonHandler(w http.ResponseWriter, r *http.Request) {
	ui.DisableButton().Render(r.Context(), w)
}

// disableButtonDataHandler will run once the /disabel
func (app *application) disableButtonDataHandler(w http.ResponseWriter, r *http.Request) {
	// struct to receive the signal from frontend
	signal := struct {
		ShouldDisable bool `json:"shouldDisable"`
	}{ShouldDisable: false}

	// read the data into the signal struct
	if err := datastar.ReadSignals(r, &signal); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// the frontend changes the signal to 'true'
	// app.infoLog.Printf("📶 signal: %#v", signal)

	// create a new event
	sse := datastar.NewSSE(w, r)

	// replace the "result" div
	// merge a new child div into the 'container' div
	now := time.Now().UTC().Format(time.RFC3339)
	sse.MergeFragmentTempl(ui.DisableButtonAppender(now), datastar.WithSelectorID("container"), datastar.WithMergeAppend())

	// sleep for 2 seconds to simulate server working
	time.Sleep(2 * time.Second)

	// update the signal back to false and merge it with the front-end
	signal.ShouldDisable = false
	sse.MarshalAndMergeSignals(signal)
}
