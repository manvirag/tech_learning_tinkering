# Golang_learning Document: https://go.dev/doc/tutorial/getting-started


**GO111MODULE**

GO111MODULE is a setting in the Go programming language that controls the behavior of Go modules, introduced in Go version 1.11 for managing dependencies. Before Go modules, dependencies were managed using the GOPATH, which had issues with version conflicts and manual dependency management. Go modules provide a better way to handle dependencies by defining modules with unique paths and version numbers in a go.mod file.

The `GO111MODULE` environment variable has three possible values:
- `GO111MODULE=off`: Disables the module system, reverting to the GOPATH approach.
- `GO111MODULE=on`: Enables the module system and uses the go.mod file to manage dependencies.In this mode, the go command looks for the go.mod file in the project directory to determine the required dependencies and their versions. If the file exists, the command downloads the required dependencies and stores them in a local cache, which can be shared between projects
- `GO111MODULE=auto`: Automatically enables the module system if a go.mod file is present; otherwise, it falls back to GOPATH.

Developers set `GO111MODULE` as an environment variable to control the module system behavior. Using Go modules ensures consistent and efficient dependency management by specifying packages and versions in the go.mod file.


1. GO111MODULE=on
2. go mod init  <name of the module (github.com/manvirag982/tech_learning_tinkering/tree/main/golang)>   

If you plan to publish your module for others to use, the module path must be a location from which Go tools can download your module. https://go.dev/doc/modules/managing-dependencies#naming_module

3. touch main.go and add code
4. go run main.go


**External Package Using From pkg.go.dev**

commit 2

1. Add import 
2. go mod tidy (When you ran go mod tidy, it located and downloaded the rsc.io/quote module that contains the package you imported. By default, it downloaded the latest version)


**External Package Using From github**

commit 3




