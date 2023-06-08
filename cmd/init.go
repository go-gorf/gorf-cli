/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
)

var projectTemplateName = "template"

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new Gorf project",
	Long: `Initialize a new Gorf project
Usage example:

gorf init myproject
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Creating project %v :)", args[0])
		CreateNewGorfProject(args[0])
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

type Project struct {
	Name string
}

func (p *Project) Create() {
	specUrl := "https://github.com/go-gorf/template/archive/refs/heads/main.zip"
	resp, err := http.Get(specUrl)
	if err != nil {
		fmt.Printf("err: %s", err)
	}

	defer resp.Body.Close()
	fmt.Println("status", resp.Status)
	if resp.StatusCode != 200 {
		return
	}

	// Create the file
	out, err := os.Create(fmt.Sprintf("%s.zip", p.Name))
	if err != nil {
		fmt.Printf("err: %s", err)
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	fmt.Printf("err: %s", err)
}

func CreateNewGorfProject(name string) {
	project := &Project{name}
	project.Create()
	fmt.Println("Successfully created project!")
	fmt.Printf(`To run the project
cd %v
go run .
`,
		project.Name)
}
