go_sources := $(shell ls *.go)
go_targets := list run

gopl: $(go_targets)

list: main/main.go $(go_sources)
	go build -o $@ $<

run: main/run.go
	go build -o $@ $<

clean:
	$(RM) $(go_targets)

