One should always remember differences between module, package, and library.

Module is a single file.
Package is directory containing modules.
Library is a collection of packages.

In Golang, each module includes package name. For example package main

All the modules contained in one package should have access to the properties (considering these properties to be public) and methods of other modules in that package (directory)

In case we want to import something from another package, we first need to determine its path determined by go.mod file

For example, we have our project main directory inside which we have basics directory. There is also another directory inside
basics folder called mypck. To determine its path we need to go to go.mod file and see how we determined virtual path for
the project. We see that it's determined as `module github.com/Yalchin403/goFun`. So github.com/Yalchin403/goFun is a virtual
path for our most outer directory. Now we can say that if it is main directory, path to the mypck folder should theoretically
be `github.com/Yalchin403/goFun/basics/mypck`. This way you can import any package and start to use exported properties.