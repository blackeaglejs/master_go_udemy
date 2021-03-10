package main

func main() {
	// GOPATH is not a friendly approach - it requires all projects to be in the same package.
	// Modules are a fix for this. This was finalized in v1.13
	// A module is a collection of related Go packages with a go.mod file at their root.
	// go.mod defines the module's path.
	// the go command has direct support for module.
	// allows you to use different versions of the same package.
}
