## Packages

Good code organization is essential for maintaining any project. Organizing code neatly into differen folders across multiple files makes it easier to test and maintain code very easily. Folders are called `packages` in Golang. There can be one or more files in a package. They are all bound by using the same package name at the beginning of the file. Importing a package is done by providing the path to the directory instead of the file in the package.

For using packages, a particular folder structure is expected as shown below:

```bash
Project Root folder
- bin
- pkg
- src
```

The `bin` folder contains binaries of the project that are created when `go install` is used.

The `pkg` folder contains compiled versions of the included packages with `.a` extension called package archives.

The `src` folder contains the source code of your project arranged under different folders and files.

---
To run the exercise, open the root folder for this exercise in VSCode and it will automatically infer the GOPATH for this project.

Running `GOPATH=$PWD go install app` will compile the code and if all goes well, will add an executable named `app` to the `bin` folder

Use of `GOPATH` is necessary as we are setting the value for the `GOPATH` environment variable just for running go install

Packages for the exercise should be created under the `src` folder without a `main` function as there should be only one `main` function for the whole application since only one entrypoint for an application is permitted.

---


### References

* [Golang standard library packages](https://golang.org/pkg/)
* [Everything you need to know about packages](https://medium.com/rungo/everything-you-need-to-know-about-packages-in-go-b8bac62b74cc)
