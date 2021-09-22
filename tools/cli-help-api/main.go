package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
)

const (
	prysmaticWebsite      = "https://prysmaticlabs.com"
	latestReleaseEndpoint = "/releases/latest"
	releaseEndpoint       = "/releases/"
	arch                  = "amd64"
	system                = "linux"
	execPerms             = 0700
)

var (
	portFlag       = flag.String("port", "3000", "port for server")
	hostFlag       = flag.String("host", "127.0.0.1", "host for server")
	componentsFlag = flag.String(
		"components",
		"beacon-chain,validator,client-stats",
		"comma-separated list of Prysm components for which to serve help text",
	)
	releaseFormat = "%s-%s-%s-%s"
)

func main() {
	flag.Parse()
	port := *portFlag
	host := *hostFlag
	componentsList := *componentsFlag
	components := strings.Split(componentsList, ",")
	version, err := latestReleaseVersion()
	if err != nil {
		log.Fatal(err)
	}
	for _, comp := range components {
		if err := downloadRelease(comp, version); err != nil {
			log.Fatal(err)
		}
	}
	r := mux.NewRouter()
	for _, comp := range components {
		tmp := comp
		r.HandleFunc("/"+comp, func(w http.ResponseWriter, r *http.Request) {
			if err := showHelpText(w, tmp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		})
	}
	log.Printf("Running API server on %s:%s", host, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), r))
}

func showHelpText(w io.Writer, component string) error {
	version, err := latestReleaseVersion()
	if err != nil {
		return err
	}
	binName := fmt.Sprintf(releaseFormat, component, version, system, arch)
	releasePath := filepath.Join("dist", binName)
	if _, err := os.Stat(releasePath); os.IsNotExist(err) {
		log.Printf("Downloading release %s, not found locally", releasePath)
		if err := downloadRelease(component, version); err != nil {
			return err
		}
	}
	cmd := exec.Command(releasePath, "--help")
	cmd.Stdout = w
	cmd.Stderr = w
	return cmd.Run()
}

func latestReleaseVersion() (string, error) {
	resp, err := http.Get(prysmaticWebsite + latestReleaseEndpoint)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	enc, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return strings.Trim(string(enc), "\n"), nil
}

func downloadRelease(component, version string) error {
	binName := fmt.Sprintf(releaseFormat, component, version, system, arch)
	resp, err := http.Get(prysmaticWebsite + releaseEndpoint + binName)
	if err != nil {
		return err
	}
	if err := os.MkdirAll("dist", execPerms); err != nil {
		return err
	}
	releasePath := filepath.Join("dist", binName)
	out, err := os.Create(releasePath)
	if err != nil {
		return err
	}
	defer out.Close()
	if _, err := io.Copy(out, resp.Body); err != nil {
		return err
	}
	// Make executable.
	return os.Chmod(releasePath, execPerms)
}
