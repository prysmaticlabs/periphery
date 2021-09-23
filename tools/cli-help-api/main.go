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
	binOutputFlag  = flag.String("binary-output-path", "", "output directory for binaries")
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
	binOutputPath := *binOutputFlag
	if binOutputPath == "" {
		log.Fatal("-binary-output-path flag not specified")
	}
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
	dirEntries, err := os.ReadDir(binOutputPath)
	if err != nil {
		log.Fatal(err)
	}
	fileNames := make([]string, 0)
	for _, entry := range dirEntries {
		fileNames = append(fileNames, entry.Name())
	}
	log.Printf("Downloaded releases into directory %s: %v", binOutputPath, fileNames)
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
	releasePath := filepath.Join(*binOutputFlag, binName)
	if _, err := os.Stat(releasePath); os.IsNotExist(err) {
		log.Printf("Downloading release %s, not found locally", releasePath)
		if err := downloadRelease(component, version); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	log.Print(os.Stat(releasePath))
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
	log.Printf("Downloading release for component %s", binName)
	resp, err := http.Get(prysmaticWebsite + releaseEndpoint + binName)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(*binOutputFlag, execPerms); err != nil {
		return err
	}
	releasePath := filepath.Join(*binOutputFlag, binName)
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
