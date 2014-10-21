package main

import (
	"fmt"
	"strings"
	"bytes"
	"io/ioutil"
	"net/http"
	"html/template"
	"crypto/sha1"
	"encoding/base64"
	"github.com/gorilla/mux"
)


var (
	ROOT_PATH = "/home/kyle/apps/earths.center/pastes/"
	PASTE_PATH = ROOT_PATH + "pastes/"
	TMPLT_PATH = ROOT_PATH + "static/"

	SITE_URL = "http://earths.center/"
	LISTEN_PORT = "8001"

	LANGS = []string{"markup", "html", "css", "clike", "javascript", "java",
		"php", "scss", "bash", "c", "cpp", "python", "sql", "ruby", "csharp",
		"go", "haskell", "objectivec", "apacheconf"}
)

type Template struct {
	Key string
	Body []byte
	Lang string
}


/*-------------------------------------
	Main
-------------------------------------*/

func main() {

	r := mux.NewRouter()

	// Landing on homepage
	r.HandleFunc("/", handleLand).
		Methods("GET")

	// Posting a paste
	r.HandleFunc("/", handlePaste).
		Methods("POST")
	
	// Reading a paste
	r.HandleFunc("/{pasteId}", handleView).
		Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":"+LISTEN_PORT, nil)
}


/*-------------------------------------
	Landing Handler
-------------------------------------*/

func handleLand(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to the jungle")
}


/*-------------------------------------
	Paste Handler
-------------------------------------*/

func handlePaste(w http.ResponseWriter, r *http.Request) {
	paste := r.FormValue("paste")

	// Generate hash to use as filename/key
	// hash is base64 encoding of the first 72 bits of sha1(paste)
	h := sha1.New()
	h.Write([]byte(paste))
	keyHash := h.Sum(nil)
	key := base64.URLEncoding.EncodeToString(keyHash[:9])

	// Save our paste
	f := PASTE_PATH + key + ".paste"
	err := ioutil.WriteFile(f, []byte(paste), 0600)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}

	u := SITE_URL + key
	fmt.Fprintf(w, "%s\n", u)
}


/*-------------------------------------
	View Handler	
-------------------------------------*/

func handleView(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["pasteId"]
	file := PASTE_PATH + key + ".paste"

	paste, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Fprintf(w, "[%s] not found", key)
	}

	// If no lang query is set, just sent back plain text
	lang, ok := r.URL.Query()["lang"]
	if !ok {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "%s", paste)
		return
	}

	// Try to load the template. Send plain text if err
	tmpl, err := template.ParseFiles(TMPLT_PATH + "template.html")
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "%s\n\n%s", err, paste)
		return
	}

	// If the Lang is a valid lang, then load the template
	l := strings.ToLower(lang[0])
	for _, validLang := range LANGS {
		if l == validLang {
			t := Template{Key: key, Body: paste, Lang: l}
			tmpl.Execute(w, t)
			return
		}
	}

	// Else just return plain text with error
	var errbuf bytes.Buffer
	errbuf.WriteString("########################################\n\n")
	errbuf.WriteString("INVALID LANG! valid langs are:\n")
	fmt.Fprintf(&errbuf, "%v\n\n", LANGS)	
	errbuf.WriteString("########################################\n")

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "%s\n\n%s", errbuf.String(), paste)
}


