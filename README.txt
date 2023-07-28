
Review

    A first type abstraction fails to compile.

    The Go programming language represents an interesting
    perspective on the tradition and history of programming
    languages.

    Categorical type theory becomes tedious in the GOLANG
    freedoms of convenience (i.e. r/t libc, type anonymity)
    and expression (i.e. struct and interface distinction).
    Features embellishing the logical abstraction with
    practical effects.

    The present moment of Go 1.20 seems encumbered with the
    tedium.


Building

    Running

        make

    fails with

	$ make
	go build -o list main/main.go
	# command-line-arguments
	main/main.go:46:15: undefined: Abstract
	make: *** [Makefile:7: list] Error 1


Debugging

    The shell script

        ./db.sh

    depends on the implicit configuration of the local go
    runtime.

    Generally, it illustrates the debugging of the compiler.

