CC=gcc
CFLAGS=-Wall
DEBUGFLAGS=-D DEBUG

all:
	$(CC) $(CFLAGS) triangle_distance.c -o tdist

debug:
	$(CC) $(DEBUGFLAGS) $(CFLAGS) triangle_distance.c -o tdist

clean:
	rm -rf tdist
