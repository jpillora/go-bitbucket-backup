package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"

	"github.com/codegangsta/cli"
)

const (
	BIT_BUCKET_URL   = "https://bitbucket.org/api/2.0/repositories"
	CREATE_REPO_BODY = `{"is_private":true,"fork_policy":"no_forks"}`
)

func main() {
	app := cli.NewApp()
	app.Name = "bitbucket-backup"
	app.Usage = "BitBucket Backup"
	app.Version = "0.1.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "directory",
			Usage: "Directory containing Git repositories (defaults to cwd)",
		},
		cli.StringFlag{
			Name:  "username",
			Usage: "BitBucket username",
		},
		cli.StringFlag{
			Name:  "password",
			Usage: "BitBucket password",
		},
		cli.StringFlag{
			Name:  "namespace",
			Usage: "BitBucket namespace (used to create repositories, defaults to username)",
		},
		cli.BoolFlag{
			Name:  "reset",
			Usage: "Resets the repositories on BitBucket before pushing (deletes then re-creates)",
		},
	}
	app.Action = run
	app.Run(os.Args)
}

func run(c *cli.Context) {
	//retrieve cli args and options
	baseDir := c.String("directory")
	if len(baseDir) == 0 {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal("Cant get working directory")
		}
		baseDir = wd
	}

	username := c.String("username")
	if len(username) == 0 {
		log.Fatal("Missing username")
	}

	password := c.String("password")
	if len(password) == 0 {
		log.Fatal("Missing password")
	}

	namespace := c.String("namespace")
	if len(namespace) == 0 {
		namespace = username
	}

	//bitbucket http helper
	request := func(method string, url string, bodyStr string) int {
		client := &http.Client{}

		body := bytes.NewBufferString(bodyStr)
		req, err := http.NewRequest(method, url, body)
		if err != nil {
			log.Fatalf("Error creating HTTP request: %s", err)
		}

		if body.Len() > 0 {
			req.Header.Add("Content-Type", "application/json")
		}

		req.SetBasicAuth(username, password)
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("Error performing HTTP request: %s", err)
		}
		return resp.StatusCode
	}

	//read all directories in cwd
	dirs, err := ioutil.ReadDir(baseDir)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %v files in '%s'\n", len(dirs), baseDir)

	for _, fileInfo := range dirs {

		//only process directories
		if !fileInfo.IsDir() {
			continue
		}

		name := fileInfo.Name()
		dir := path.Join(baseDir, name)

		//only process git directories
		if _, err := os.Stat(path.Join(dir, ".git")); err != nil {
			continue
		}

		fmt.Printf("Repo '%s/%s'\n", namespace, name)

		url := BIT_BUCKET_URL + "/" + namespace + "/" + name

		if c.Bool("reset") {
			fmt.Printf("  Deleting repository...\n")
			request("DELETE", url, "")
		}

		fmt.Printf("  Checking repository...\n")
		if checkStatus := request("HEAD", url, ""); checkStatus == 404 {
			//repo missing!
			//create it now...
			if status := request("POST", url, CREATE_REPO_BODY); status == 200 {
				fmt.Printf("  Created missing repository\n")
			} else {
				log.Fatalf("  Failed to create missing repository\n")
			}
		} else if checkStatus != 200 {
			//must report back ok!
			log.Fatalf("  Failed to check repository (status %v)\n", checkStatus)
		}

		cmd := exec.Command("git", "push", "--repo=https://"+username+":"+password+"@bitbucket.org/"+namespace+"/"+name+".git", "--all")
		cmd.Dir = dir

		fmt.Printf("  Git pushing repository...\n")
		err := cmd.Run()

		if err != nil {
			log.Fatalf("  Failed to push repository (%s)\n", err)
		}

		fmt.Printf("  Done\n")
	}
}
