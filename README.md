Examples of Go and App Engine Structure
========================================

I wrote this repository, because every time I went to start a new Go project on Google App Engine
I couldn't remember an optimal way to structure my application and how to manage my GOPATH in relation
to my application code.

This repository aims to show several options for managing GOPATHs when working with a Go App Engine projects.

It assumes that you already have the [Google Cloud Platform SDK](https://cloud.google.com/sdk/) installed.

If you have a different way that you feel is appropriate, pull requests are more than welcome.

It should be noted that while this repository makes relatively heavy use of [Make](http://www.gnu.org/software/make/), you only need to be able to 
run Make commands, not understand the files (but it is a useful thing to know!)

For a full description of what this repository is showing read the **[full blog post](http://www.compoundtheory.com/configuring-your-gopath-with-go-and-google-app-engine)**.

Except for the GB tasks, the GOPATH is controlled within the Makefiles that are in this project. 
This means you can run any of the Make tasks in this project to test them out, 
without having to change a GOPATH you already have set.

## Steps for Exploring This Repository

### Setup

1. Clone the repository: `git clone https://github.com/markmandel/appengine-golang-gopath.git`
1. `cd appengine-golang-gopath`
1. Install the go tooling we use for this: `make install-tools` 

### The Basic GOPATH Module

1. Enter the directory: `cd src/modules/basic`
1. Review the `routes.go` file to see what is being executed: `cat routes.go` 
1. Have a look at the current GOPATH: `make debug-env`
1. Install all dependencies `make deps`
1. Have a look where those dependencies end up in the directory structure: `(cd ../.. && pwd && ls)`
1. Run the module: `make serve`
1. Browse to http://localhost:8080 to see the results.
1. (Optional) - Deploy the module: `make deploy`
1. Have a look at the local `Makefile` and `Makefile.goapp` to see how the steps worked: `cat Makefile` and `cat ../../../Makefile.goapp`

### The Vendored GOPATH Module

1. Enter the directory: `cd ../vendored`
1. Clean the previous dependencies: `make clean`
1. Review the `routes.go` file to see what is being executed: `cat routes.go` 
1. Have a look at the current GOPATH: `make debug-env`. Notice the two part GOPATH.
1. Install all dependencies `make deps`
1. Have a look where those dependencies end up in the directory structure: `(cd ../../../vendor/src && pwd && ls)`
1. Run the module: `make serve`
1. Browse to http://localhost:8080 to see the results.
1. (Optional) - Deploy the module: `make deploy`
1. Have a look at the `Makefile.goapp` to see how the steps worked: `cat ../../../Makefile.goapp`

### The GB Module

1. Enter the directory: `cd ../gb`
1. Clean the previous dependencies: `make clean`
1. Review the `routes.go` file to see what is being executed: `cat routes.go` 
1. Have a look at the current GOPATH: `make debug-env`. Notice there is no GOPATH. Crazy.
1. Install all dependencies `make deps`
1. Have a look where those dependencies end up in the directory structure: `(cd ../../../vendor/src && pwd && ls)` 
1. Have a look at the `manifest` file: `cat ../../../vendor/manifest`.
1. Run the module: `make serve`
1. Browse to http://localhost:8080 to see the results.
1. (Optional) - Deploy the module: `make deploy`
1. Have a look at the local `Makefile` and `Makefile.gb` to see how the steps worked: `cat Makefile` and `cat ../../../Makefile.gb` 

## Licence
Apache 2.0

This is not an official Google product.
