#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include <unistd.h>

void printHelp(char*);
void printUsage(char*);

int main(int argc, char** argv) {
  int fFilled = 0, sFilled = 0;
  int triangular = 0;
  int arraySize;
  float fillingPercentage;
  int c;
  while ((c = getopt (argc, argv, "s:f:ht")) != -1) {
    switch (c) {
      case 's':
        arraySize = atoi(strdup(optarg));
        sFilled = 1;
        break;
      case 'f':
        fillingPercentage = atof(strdup(optarg));
        fFilled = 1;
        break;
      case 't':
        triangular = 1;
        break;
      case 'h':
        printHelp(argv[0]);
        return 0;
      case '?':
      default:
          printUsage(argv[0]);
          return 1;
    }
  }

  if (!sFilled || !fFilled) {
    printUsage(argv[0]);
    return 1;
  }

#ifdef DEBUG
  printf("passed parameters = %d, %f\n", arraySize, fillingPercentage);
#endif

  int number;
  int amountOfNumbers = 0;

  srand(time(NULL));

  float fillChance;
  int row, col;
  while (scanf("%d", &number) != EOF) {
    row = amountOfNumbers / arraySize;
    col = amountOfNumbers % arraySize;

    fillChance = rand() % 101 / 100.0;
    if ((fillingPercentage < 1.0 - fillChance) || (triangular && row > col)) {
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
  printf("Usage: %s [-t] -s [array size] -f [filling percentage]\n", commandName);
}

void printHelp(char* commandName) {
  printUsage(commandName);
  printf("\nThis program reads input array of numbers from stdin and writes "
      "output to stdout.\n"
      "The output consists of elements of given array and '?' signs.\n"
      "'?' signs are placed instead of given numbers with probability of "
      "[filling percentage]\n");
}

