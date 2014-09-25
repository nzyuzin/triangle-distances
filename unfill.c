#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

void printHelp(char*);
void printUsage(char*);

int main(int argc, char** argv) {
  if (argc != 3
      || strcmp(argv[1], "--help") == 0 || strcmp(argv[1], "-h") == 0) {
#ifdef DEBUG
    printf("incorrect input: argc = %d\n", argc);
#endif
    printHelp(argv[0]);
    return 1;
  }

  int arraySize = atoi(argv[1]);
  float fillingPercentage = atof(argv[2]);
#ifdef DEBUG
  printf("passed parameters = %d, %f\n", arraySize, fillingPercentage);
#endif

  int number;
  int amountOfNumbers = 0;

  srand(time(NULL));

  float fillChance;
  while (scanf("%d", &number) != EOF) {
    fillChance = rand() % 101 / 100.0;
    if (fillingPercentage < 1.0 - fillChance) {
      printf("?");
    } else {
      printf("%d", number);
    }
    amountOfNumbers++;
    if (amountOfNumbers % arraySize == 0) {
      printf("\n");
    } else {
      printf("\t");
    }
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

void printUsage(char* commandName) {
  printf("Usage: %s [array size] [filling percentage]\n", commandName);
}

void printHelp(char* commandName) {
  printUsage(commandName);
  printf("\nThis program reads input array of numbers from stdin and writes "
      "output to stdout.\n"
      "The output consists of elements of given array and '?' signs.\n"
      "'?' signs are placed instead of given numbers with probability of "
      "[filling percentage]\n");
}

