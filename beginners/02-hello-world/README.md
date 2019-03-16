# Hello World

We are going to start with the traditional `Hello World` program. The goal is to write code to print `Hello World` on the terminal.

## Steps

* Create a folder named `hello-world`
* Assign the environment variable `GOPATH` to point to the absolute path of the above folder. Here is an example: `export GOPATH="/home/users/allen/workspace/hello-world"`
* Open this folder in VSCode `File > Open`. You should be able to see a side pane with the folder opened
* Create a file called `main.go` from `File > New File`
* After the file is created, VSCode Go extension will kick in and start downloading few packages. Below is an example on my machine:

```bash
Installing 10 tools at /Users/allen/go/bin
  gocode
  gopkgs
  go-outline
  go-symbols
  guru
  gorename
  dlv
  godef
  goreturns
  golint

Installing github.com/mdempsky/gocode SUCCEEDED
Installing github.com/uudashr/gopkgs/cmd/gopkgs SUCCEEDED
Installing github.com/ramya-rao-a/go-outline SUCCEEDED
Installing github.com/acroca/go-symbols SUCCEEDED
Installing golang.org/x/tools/cmd/guru SUCCEEDED
Installing golang.org/x/tools/cmd/gorename SUCCEEDED
Installing github.com/go-delve/delve/cmd/dlv SUCCEEDED
Installing github.com/rogpeppe/godef SUCCEEDED
Installing github.com/sqs/goreturns SUCCEEDED
Installing golang.org/x/lint/golint SUCCEEDED

All tools successfully installed. You're ready to Go :).
```

The above tools help to write code that conforms to the Golang code standards.

* Add the following snippet to the file

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}
```

* Execute the able file with the command `go run main.go`
* `Hello World` is printed on the terminal

## Deep Dive

In this section, we are going to evaluate what each line means in the `Hello World` program.

Line 1 => `package main`

Packages are how you group common functionality in Golang. The default package is `main` and is usually used to signify that this file will be an entrypoint for an executable. Since we are using only one file, we set the package name to be `main`

Line 3 => `import "fmt"`

This line indicates that the package named `fmt` that is shipped as part of Golang, is required for execution of this program. How this is used will be explained later

Line 5 => `func main() {`

Functions in Golang are defined with the keyword `func`. Here a function named `main` is defined that takes no input arguments. The function body is defined between curly braces `{ }`

To run file a file in Golang as an executable, there should be a `main` function and the package should be `main`. There can be only one `main` function in a Go program even if there are multiple `.go` files i.e. there can be only one entry point.

Line 6 => `fmt.Println("Hello World")`

This is the line that displays `Hello World` on the terminal. `Println` is a function in the `fmt` package which takes one string argument `"Hello World"`. The `.` is used to define the path for Golang to find the function. So `fmt.Println` means the function `Println` is defined in the `fmt` package.

The function `Println` does two things:

* Displays the input as provided by the user
* Adds a new line after the string is displayed
