## Hello World

We are going to start with the traditional `Hello World` program. The goal is to write code to print `Hello World` on the terminal.

### Steps
* Create a folder named `hello-world`
* Assign the environment variable `GOPATH` to point to the absolute path of the above folder. Here is an example: `export GOPATH="/home/users/allen/workspace/hello-world"`
* Open this folder in VSCode `File > Open`. You should be able to see a side pane with the folder opened
* Create a file called `hello.go` from `File > New File`
* After the file is created, VSCode Go extension will kick in and start downloading few packages. Below is an example on my machine:

```
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

```
package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}
```

* Execute the able file with the command `go run hello.go`
* `Hello World` is printed on the terminal

