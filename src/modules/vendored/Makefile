#
# This is a Makefile specific to this vendored module.
#

#Include the core Makefile and all its targets
include ../../../Makefile.core
include ../../../Makefile.goapp

#now let's change the GOPATH to have a `vendor` prefix.
GOPATH := $(GOPATH)/vendor:$(GOPATH)
export GOPATH

#Add our vendored GOPATH bin directory to the path as well
PATH := $(GOPATH)/vendor/bin:$(PATH)
export PATH