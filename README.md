# gorf-cli

## Installation using go install

```bash
go install github.com/go-gorf/gorf-cli@latest
```

## Installation using binary  

* Download the latest binary executable from the latest release page for your machine (https://github.com/go-gorf/gorf-cli/releases/latest)
* Extract the file
* Run the executable with required permissions or move the file to your $PATH


## Quickstart

```bash
gorf-cli init hello
```

## Run Project

```bash
cd hello
go run .
```
Now look at localhost:8080/health on your browser

## Add new app  

CD to your apps directory 

```bash
gorf-cli app newApp
```

Add the app to the settings.go in the project root.
