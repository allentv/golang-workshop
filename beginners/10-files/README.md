# Files

Using files is one of the most common tasks in any programming language and is one of the ways of persisting information for future retrieval. Information can be read from a file and written to a file

## Reading from files

Information is read from a file using the `ReadFile` function from the `ioutil` package. Assume that there is a text file called `sample.txt` in the same folder as the go program which has some lines of random text in it.

The below code snippet reads from the example file, checks if there is an error during the reading process and then displays the content of the file.

```go
import (  
    "fmt"
    "io/ioutil"
)

func main() {  
    data, err := ioutil.ReadFile("sample.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    fmt.Println("Contents of file:", string(data))
}
```

The `ReadFile` takes the path to the input file as the argument. Since this file is in the same folder as the `main.go` file, only the file name is required. If the location is different, then the full path to the file is required for the go program to find the file.

The `ReadFile` returns two values - the actual information and an error information object which has details of any possible error that occurred while reading the file.

Checking if the error object (`err`) is empty (`nil`) is idomatic Go or standard practice in Golang. If there is an error, then print an error message and quit the program or halt further execution of the program.

The information is read from the file as an array of bytes. Converting them into a string is done using the `string` function.

## Writing to files

Data can be written to a file using the `WriteFile` function from the `ioutil` package. The information should be converted to an array of bytes before writing to the file.

```go
data := []byte("This is some information!")
err := ioutil.WriteFile("write_data.txt", data, 0666)
if err != nil {
    fmt.Println("There has been an error:", err)
    return
}
```

The data is a string that is converted into a `byte` array with the `byte` function. The `WriteFile` requires the file mode to be used to signify the permissions that this new file should have in the file system. Golang compiler will complain if the permissions is not set. The return value of `WriteFile` is an error object.

If it is empty, that means there is no error and a file named `write_data.txt` should be created in the same folder as the golang program.

If there is an error, then the error is printed out and the program quits.
