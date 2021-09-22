package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/gorilla/mux"
)

const (
	prysmSh     = "https://raw.githubusercontent.com/prysmaticlabs/prysm/develop/prysm.sh"
	prysmShPath = "prysm.sh"
	execPerms   = 0700
)

var (
	portFlag       = flag.String("port", "3000", "port for server")
	hostFlag       = flag.String("host", "127.0.0.1", "host for server")
	componentsFlag = flag.String(
		"components",
		"beacon-chain,validator,client-stats",
		"comma-separated list of Prysm components for which to serve help text",
	)
)

func main() {
	flag.Parse()
	port := *portFlag
	host := *hostFlag
	componentsList := *componentsFlag
	components := strings.Split(componentsList, ",")
	if err := downloadScript(prysmShPath); err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	for _, comp := range components {
		r.HandleFunc("/"+comp, func(w http.ResponseWriter, r *http.Request) {
			if err := showHelpText(w, comp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		})
	}
	log.Printf("Running API server on %s:%s", host, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), r))
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
	return os.Chmod(filepath, execPerms)
}
