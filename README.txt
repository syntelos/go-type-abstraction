
Review

    A first type abstraction fails to compile.

    The test case converted from "build main" file structure
    to "run" file structure works.


Building

    Running

        make list

    fails with

	$ make
	go build -o list main/main.go
	# command-line-arguments
	main/main.go:46:15: undefined: Abstract
	make: *** [Makefile:7: list] Error 1

    See
        ./main/main.go
	./types.go


    Running

        make run

    succeeds.  See './main/run.go'.


Version

    go 1.19.11
    go 1.20.5
    go 1.20.6
    go 1.21.0


References

    [go] https://go.dev/
    [devel] https://github.com/golang/go
    [devel] https://go.googlesource.com/go
    [delve] https://github.com/go-delve/delve
    [play] https://go.dev/play/p/5mr5M0luZ9k

