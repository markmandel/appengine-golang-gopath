#
# This root makefile exists primarily to install the tools
# you will need for this sample code, but you can run any of the
# core tasks from here as well.
#

include Makefile.core
include Makefile.goapp

#install any go tools we will need for this project
#and then clean up again afterwards, so we just have
#the binaries.
install-tools:
	go get github.com/constabulary/gb/...
	go get code.palmstonegames.com/gb-gae
	goapp get github.com/golang/lint/golint
	goapp get golang.org/x/tools/cmd/goimports

	#cleaning up source code, and leaving behind binaries...
	rm -rf $(ROOT)/src/github.com/constabulary
	rm -rf $(ROOT)/src/code.palmstonegames.com
	rm -rf $(ROOT)/src/golang.org
	rm -rf $(ROOT)/src/github.com/golang
	-rmdir $(ROOT)/src/github.com #if the dir is empty, kill it