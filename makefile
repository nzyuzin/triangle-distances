CC=gcc
GO=go build
CFLAGS=-Wall
DEBUGFLAGS=-D DEBUG
UNFILL=unfill
FILL=fill
COMPARE=compare
NORMALIZE=normalize
BADNESS=badness
EXECUTABLES=$(FILL) $(UNFILL) $(COMPARE) $(NORMALIZE) $(BADNESS)

all: c go

cdebug:
	$(CC) $(DEBUGFLAGS) $(CFLAGS) $(UNFILL).c -o $(UNFILL)

c:
	$(CC) $(CFLAGS) $(UNFILL).c -o $(UNFILL)

go:
	$(GO) -o $(FILL) $(FILL).go
	$(GO) -o $(COMPARE) $(COMPARE).go
	$(GO) -o $(NORMALIZE) $(NORMALIZE).go
	$(GO) -o $(BADNESS) $(BADNESS).go

clean:
	rm -rf $(EXECUTABLES)
