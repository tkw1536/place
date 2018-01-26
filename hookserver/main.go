package main

import (
	"os"
	"log"
	"fmt"
	"./utils"
	"net/http"
	"flag"
	"io"
	"io/ioutil"
)

var (
	// the logger
	Info    *log.Logger

	// arguments on command line
	hookSecret  string
	bindAddress string
	listenPath  string
	hook		[]string
	quiet		bool
	timeout     int
)

func handler(res http.ResponseWriter, req *http.Request) {
	Info.Printf("[%s] %s\n", req.Method, req.URL)

	gitHubCheck := utils.IsValidGithubRequest(req, hookSecret)
	if gitHubCheck == nil {
		Info.Print("Received valid GitHub event\n")
		fmt.Fprintf(res, "Success\n")
		go utils.RunHook(hook, timeout, Info)
		return
	} else {
		Info.Printf("Not a valid GitHub event: %s\n", gitHubCheck.Error())
	}

	gitLabCheck := utils.IsValidGitLabRequest(req, hookSecret)
	if gitLabCheck == nil {
		Info.Print("Received valid GitLab event\n")
		fmt.Fprintf(res, "Success\n")
		go utils.RunHook(hook, timeout, Info)
		return
	} else {
		Info.Printf("Not a valid GitLab event: %s\n", gitLabCheck.Error())
	}


	http.Error(res,"Failure\n", 500)
}



func main() {
	// parse arguments
	var hookLine string

	flag.StringVar(&bindAddress, "bind", "127.0.0.1:3000", "address to bind to")
	flag.StringVar(&listenPath, "path", "/hook/", "path to listen to")
	flag.StringVar(&hookSecret, "secret", "", "hook secret")
	flag.BoolVar(&quiet, "quiet", false,"disable logging on stdout")
	flag.StringVar(&hookLine, "hook", "/bin/false", "executable to run for hook")
	flag.IntVar(&timeout, "timeout", 600, "timeout for hook script in seconds")
	flag.Parse()

	var err error
	hook, err = utils.SplitArguments(hookLine)
	if err != nil {
		panic(err.Error())
	}

	// setup the logger
	var logger io.Writer
	if quiet {
		logger = ioutil.Discard
	} else {
		logger = os.Stdout
	}
	Info = log.New(logger, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	// print some info
	Info.Printf("Bind Address: %s\n", bindAddress)
	Info.Printf(" Listen Path: %s\n", listenPath)
	Info.Printf(" Hook Script: %s\n", hookLine)
	Info.Printf("     Timeout: %d\n", timeout)
	Info.Printf("      Secret: %s\n", hookSecret)

	// and start the server
	http.HandleFunc(listenPath, handler)
	http.ListenAndServe(bindAddress, nil)
}
