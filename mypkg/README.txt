// A go module is a directory tree containing go.mod file
// A go module contains a collection of go packages
// A go module is versioned to track releases

// A go package is a directory with a collection of .go source files
// each .go file in a package directory will have "package <name>" at the top



// initialize a go module
//      1) create a directory
//      2) change to the directory
//      3) run "go mod init <module_name>"
//      4) this creates go.mod file in the current directory

PS D:\Learning\Go\Projects\mypkg> go mod init mypkg
go: creating new go.mod: module mypkg
go: to add module requirements and sums:
        go mod tidy
PS D:\Learning\Go\Projects\mypkg>

// download a external go module
// run "go get <module>"
// example: "go get github.com/google/uuid"

PS D:\Learning\Go\Projects\mypkg> go get github.com/google/uuid
go: downloading github.com/google/uuid v1.6.0
go: added github.com/google/uuid v1.6.0
PS D:\Learning\Go\Projects\mypkg> 


// to add module requirements and sums
// run "go mod tidy"