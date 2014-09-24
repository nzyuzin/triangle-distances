#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void printHelp(char*);

int main(int argc, char** argv) {
  if (argc != 2 || strcmp(argv[1], "--help") == 0 || strcmp(argv[1], "-h") == 0) {
#ifdef DEBUG
    printf("incorrect input: argc = %d\n", argc);
#endif
    printHelp(argv[0]);
    return 1;
  }

  int arraySize = atoi(argv[1]);
  int array[arraySize][arraySize];
#ifdef DEBUG
  printf("passed parameter = %d\n", arraySize);
#endif

  int number;
  int amountOfNumbers = 0;
  while (scanf("%d", &number) != EOF) {
    array[amountOfNumbers / arraySize][amountOfNumbers % arraySize] = number;
    amountOfNumbers++;
  }

#ifdef DEBUG
  printf("read %d numbers\n", amountOfNumbers);
#endif

  if (amountOfNumbers != arraySize * arraySize) {
    fprintf(stderr, "Error: Input size doesn't match given array size ^ 2");
    return 2;
  }
  return 0;
}

void printHelp(char* commandName) {
  printf("Usage: %s [array size]\n", commandName);
}

