package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing"
)

var infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
var errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)

func CheckErr(e error) {
	if e != nil {
		errorLog.Println(e)
	}
}

var cyan = color.New(color.FgCyan)
var yellow = color.New(color.FgHiYellow)
var red = color.New(color.FgHiRed)
var bold = color.New(color.Bold)
var green = color.New(color.FgHiGreen)

func main() {

	if len(os.Args) <= 1 {
		errorLog.Println("Missing branchname")
		os.Exit(1)
	}
	branch := os.Args[1]

	// loop over directories of current path:
	infoLog.Println("getting directories ...")
	matches, _ := filepath.Glob("*")
	// var dirs []string
	for _, match := range matches {
		f, _ := os.Stat(match)
		if f.IsDir() {
			// dirs = append(dirs, match)
			cyan.Println("opening repo: ", f.Name())
			r, err := git.PlainOpen(f.Name())
			CheckErr(err)
			remotes, err := r.Remotes()
			CheckErr(err)

			/* retrieving the branch being pointed by HEAD
			ref, err := r.Head()
			CheckErr(err)
			bold.Println("Ref: ", ref.Name()) */

			// for every remote:
			for i, remotename := range remotes {
				fmt.Println(i, ":", remotename.Config().URLs)
				// Get the working directory for the repository
				w, err := r.Worktree()
				CheckErr(err)
				fmt.Println("Worktree: ", w.Filesystem)

				// ... checking out branch
				yellow.Println("git checkout ", branch)
				branchRefName := plumbing.NewBranchReferenceName(branch)
				branchCoOpts := git.CheckoutOptions{
					Branch: plumbing.ReferenceName(branchRefName),
					Force:  false,
				}
				if err := w.Checkout(&branchCoOpts); err != nil {
					red.Println("checkout failed.", err)
				} else {
					green.Println("OK")
				}

				// git status:
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
						red.Println("Ups: ", err)
						yellow.Println("(please check manually if you need to.)")
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
