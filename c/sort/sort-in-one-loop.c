#include <stdio.h>
#include <stdlib.h>

void sort_in_one_loop()
{
  printf("sort in one loop\n");

  int array_length = 50;
  int array_number[array_length];
  int i;
  int x;

  printf("random number:");
  for (i = 0; i < array_length; i++)
  {
    array_number[i] = rand() % 10;
    printf(" %d", array_number[i]);
  }
  printf("\n");

  i = 0;

  while (i < array_length)
  {
    if (i != (array_length - 1) && i >= 0 && array_number[i] > array_number[i + 1])
    {
      x = array_number[i];
      array_number[i] = array_number[i + 1];
      array_number[i + 1] = x;
      i -= 2;
    }
    i++;
  }

  printf("sorted number: ");
  for (i = 0; i < array_length; i++)
  {
    printf(" %d", array_number[i]);
  }
  printf("\n");
}
