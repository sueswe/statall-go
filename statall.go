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
	green := color.New(color.FgHiGreen)

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

			// try for every remote:
			for i, remotename := range remotes {
				fmt.Println(i, ":", remotename)
				// Get the working directory for the repository
				w, err := r.Worktree()
				CheckErr(err)

				state, _ := w.Status()
				fmt.Print("Status: ")

				if state.IsClean() {
					green.Print("clean")
					fmt.Println("\ntrying to pull ...")
					err = w.Pull(&git.PullOptions{
						RemoteName: remotename.Config().Name,
					})
					if err == nil {
						green.Println("Success!")
					} else {
						yellow.Println(err)
					}
					//CheckErr(err)
					// Print the latest commit that was just pulled
					ref, err := r.Head()
					CheckErr(err)
					commit, err := r.CommitObject(ref.Hash())
					CheckErr(err)

					bold.Println("--- Last commit: ---")
					fmt.Println(commit)
					fmt.Println("")
				} else {
					red.Print(state)
					red.Println("(pull denied)")
					fmt.Println("")
				}

			}

		}
	}
}
