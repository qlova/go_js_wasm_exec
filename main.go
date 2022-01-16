package main

import (
	"go/build"
	"net/http"
	"os"

	"github.com/pkg/browser"

	_ "embed"
)

//go:embed wasm.html
var html []byte

func main() {

	wasm, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	js, err := os.ReadFile(build.Default.GOROOT + "/misc/wasm/wasm_exec.js")
	if err != nil {
		panic(err)
	}

	browser.OpenURL("http://localhost:8080")

	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			w.Write(html)
		case "/go.wasm":
			w.Header().Set("Content-Type", "application/wasm")
			w.Write(wasm)
		case "/wasm_exec.js":
			w.Header().Set("Content-Type", "application/javascript")
			w.Write(js)
		}
	}))
}
