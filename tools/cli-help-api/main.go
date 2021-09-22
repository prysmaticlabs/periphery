package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gorilla/mux"
)

const (
	prysmSh     = "https://raw.githubusercontent.com/prysmaticlabs/prysm/develop/prysm.sh"
	prysmShPath = "prysm.sh"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT env var not found")
	}
	if err := downloadScript(prysmShPath); err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/beacon-chain", func(w http.ResponseWriter, r *http.Request) {
		if err := showHelpText(w, "beacon-chain"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	r.HandleFunc("/validator", func(w http.ResponseWriter, r *http.Request) {
		if err := showHelpText(w, "validator"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	r.HandleFunc("/client-stats", func(w http.ResponseWriter, r *http.Request) {
		if err := showHelpText(w, "client-stats"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	log.Fatal(http.ListenAndServe(fmt.Sprintf("127.0.0.1:%s", port), r))
}

func showHelpText(w io.Writer, component string) error {
	cmd := exec.Command("/bin/sh", "./prysm.sh", component, "--help")
	cmd.Stdout = w
	cmd.Stderr = w
	return cmd.Run()
}

func downloadScript(filepath string) error {
	resp, err := http.Get(prysmSh)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	if _, err := io.Copy(out, resp.Body); err != nil {
		return err
	}
	// Make executable.
	return os.Chmod(filepath, 0700)
}
