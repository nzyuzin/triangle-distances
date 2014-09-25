CC=gcc
CFLAGS=-Wall
DEBUGFLAGS=-D DEBUG
SOURCE=unfill.c
EXECUTABLE=unfill

all:
	$(CC) $(CFLAGS) $(SOURCE) -o $(EXECUTABLE)

debug:
	$(CC) $(DEBUGFLAGS) $(CFLAGS) $(SOURCE) -o $(EXECUTABLE)

clean:
	rm -rf $(EXECUTABLE)
