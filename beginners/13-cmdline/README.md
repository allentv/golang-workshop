# Command Line

## Args

Interacting with the command line is a very common operation and is very useful when executing scripts or programs without a GUI. The `os` package is used to get access to the command line arguments when executing a Go program.

```go
import "os"

os.Args // Complete list including program name
os.Args[1:] // Just the arguments
```

The command line arguments are available as a Golang array in `os.Args` for easy access.

## Flags

Providing the arguments as space separated values in a terminal is very basic and makes it hard to extend. Using command line flags provide flexibility in how arguments should be passed especially the optional ones. Golang provides the `flag` package for supporting `int`, `string` and `bool` types.

```go
import "flag"

flag.<type>(name of flag, default value, description)
```

A flag is defined in the above syntax. The `<type>` can be one of `String`, `Int` or `Bool`.

A flag can be required by checking if value of the flag variable is the same as the default value. This necessarily does not mean that the user cannot provide the flag with default value during runtime.

A default value can be provided when defining a flag which is quite useful when sensible defaults are available. So if the flag is skipped in the command line, then the default value will be assigned for that flag.

Also all the flags should appear before positional arguments else they will be interpreted as such.

## Environment Variables

Environment variables are used to configure systems in mostly UNIX systems but is also available on Windows machines. These are key-value pairs that are usually available to all programs running from the terminal/shell. Custom environment variables can also be defined that is valid only during the execution of a shell script.

The `os` package provides two functions - `Setenv` and `Getenv` to set and get the value of an environment variable. An empty string is returned if the key is not present.

```go
import "os"

os.Setenv("FOO", "1")
fmt.Println("FOO:", os.Getenv("FOO"))
fmt.Println("BAR:", os.Getenv("BAR"))

> FOO: 1
  BAR:
```

The `FOO` variable is given a value of `1` at the start of the program while the value of `BAR` is looked up in the existing list of environment variables.

The `Environ` function gives the list of all the available environment variables.

```go
import "strings"

for _, e := range os.Environ() {
    pair := strings.Split(e, "=")
    fmt.Println(pair[0])
}
```

The existing environment variables are returned in the format `KEY=VALUE`. So the string has to be split at `=` to get the key and value. This is done through `strings.Split` function which returns an array.