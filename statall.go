package main

import (
	"fmt"
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
	yellow := color.New(color.FgHiYellow)
	red := color.New(color.FgHiRed)
	bold := color.New(color.Bold)
	// loop ueber repo-pfade:
	infoLog.Println("looping ...")
	matches, _ := filepath.Glob("*")
	// var dirs []string
	for _, match := range matches {
		f, _ := os.Stat(match)
		if f.IsDir() {
			// dirs = append(dirs, match)
			c.Println("opening repo: ", f.Name())
			//infoLog.Println("opening repo ...")
			r, err := git.PlainOpen(f.Name())
			CheckErr(err)
			remotes, err := r.Remotes()
			CheckErr(err)
			//fmt.Println(remotes)

			//checkout master?
			// ... retrieving the branch being pointed by HEAD
			ref, err := r.Head()
			CheckErr(err)
			bold.Println("Ref: ", ref.Name())

			for i, remotename := range remotes {
				fmt.Println(i, ":", remotename)
				w, err := r.Worktree()
				CheckErr(err)

				state, _ := w.Status()
				fmt.Print("Status: ")
				red.Println(state)

				fmt.Println("trying to pull ...")
				err = w.Pull(&git.PullOptions{
					RemoteName: remotename.Config().Name,
				})
				yellow.Println(err)
				//CheckErr(err)
				fmt.Println("")

			}

			// Get the working directory for the repository

		}
	}
}
