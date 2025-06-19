package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/go-git/go-git/v6"
)

var infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
var debugLog = log.New(os.Stdout, "DEBUG\t", log.Ldate|log.Ltime|log.Lshortfile)
var errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)

func CheckErr(e error) {
	if e != nil {
		errorLog.Println(e)
	}
}

func main() {
	c := color.New(color.FgCyan)
	// loop ueber repo-pfade:
	infoLog.Println("looping ...")
	matches, _ := filepath.Glob("*")
	// var dirs []string
	for _, match := range matches {
		f, _ := os.Stat(match)
		if f.IsDir() {
			// dirs = append(dirs, match)
			c.Println("opening repo ", f.Name())
			//infoLog.Println("opening repo ...")
			r, err := git.PlainOpen(f.Name())
			CheckErr(err)
			remotes, err := r.Remotes()
			CheckErr(err)
			//fmt.Println(remotes)

			for i, remotename := range remotes {
				c.Println(i, "", remotename)
				/*
					r.Fetch(&git.FetchOptions{
						RemoteName: remotename.String(),
					})
				*/
			}

			// Get the working directory for the repository
			w, err := r.Worktree()
			CheckErr(err)
			err = w.Pull(&git.PullOptions{RemoteName: "origin"})
			CheckErr(err)

		}
	}
}
