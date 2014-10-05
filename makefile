CC=gcc
GO=go build
CFLAGS=-Wall
DEBUGFLAGS=-D DEBUG
UNFILL=unfill
FILL=fill
COMPARE=compare
EXECUTABLES=$(FILL) $(UNFILL) $(COMPARE)

all: c go

cdebug:
	$(CC) $(DEBUGFLAGS) $(CFLAGS) $(UNFILL).c -o $(UNFILL)

c:
	$(CC) $(CFLAGS) $(UNFILL).c -o $(UNFILL)

go:
	$(GO) -o $(FILL) $(FILL).go
	$(GO) -o $(COMPARE) $(COMPARE).go

clean:
	rm -rf $(EXECUTABLES)
