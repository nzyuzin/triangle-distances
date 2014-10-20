CC=gcc
GO=go build
CFLAGS=-Wall
DEBUGFLAGS=-D DEBUG
UNFILL=unfill
FILL=fill
COMPARE=compare
NORMALIZE=normalize
EXECUTABLES=$(FILL) $(UNFILL) $(COMPARE) $(NORMALIZE)

all: c go

cdebug:
	$(CC) $(DEBUGFLAGS) $(CFLAGS) $(UNFILL).c -o $(UNFILL)

c:
	$(CC) $(CFLAGS) $(UNFILL).c -o $(UNFILL)

go:
	$(GO) -o $(FILL) $(FILL).go
	$(GO) -o $(COMPARE) $(COMPARE).go
	$(GO) -o $(NORMALIZE) $(NORMALIZE).go

clean:
	rm -rf $(EXECUTABLES)
