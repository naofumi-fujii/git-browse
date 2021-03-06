package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/browser"
)

func parsedURL(gitRemoteURL string) (*url.URL, error) {
	return url.Parse(getFormattedGitRemoteURL(gitRemoteURL))
}

func main() {
	u, err := parsedURL(getGitRemoteURL())
	if err != nil {
		log.Fatal(err)
	}

	args := os.Args
	browser.OpenURL(getTargetURL(args, u))
}

func getTargetURL(args []string, u *url.URL) string {
	if len(args) < 2 {
		return strings.TrimSuffix("https://"+u.Hostname()+u.Path, "\n")
	} else {
		return getCommitHashURLByHostingService(u, args)
	}
}

func getCommitHashURLByHostingService(u *url.URL, args []string) string {
	commitHash := args[1]
	if strings.Contains(u.Hostname(), "bitbucket") {
		return strings.TrimSuffix("https://"+u.Hostname()+u.Path, "\n") + "/commits/" + commitHash
	}
	return strings.TrimSuffix("https://"+u.Hostname()+u.Path, "\n") + "/commit/" + commitHash
}

func getFormattedGitRemoteURL(gitRemoteURL string) string {
	r1 := strings.Replace(gitRemoteURL, "git@github.com:", "ssh://git@github.com/", -1)
	r2 := strings.Replace(r1, ".git", "", -1)
	return r2
}

func getGitRemoteURL() string {
	out, err := exec.Command("git", "remote", "get-url", "origin").Output()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	url := string(out)
	return url
}
