.PHONY: all clean

all: example.so

%.so: % %/*.pyx %/*.go %/go.mod %/*.c go.mod *.go
	go get $*
	go build -buildmode=c-shared -o $@

%.c: %.pyx
	cython $<

go.mod:
	go mod init main

%/go.mod: % go.mod
	cd $* && go mod init $*
	go mod edit -replace $*=./$*

clean:
	-rm go.mod example/go.mod example/example.c
