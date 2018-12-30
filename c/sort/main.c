#include <stdio.h>
#include "sort-in-one-loop.c"

void hr();
void measure();
void sort_in_one_loop();

int main()
{
  printf("\t\tCode Geek\n\n");
  printf("\t\tLanguage: C\n");
  printf("\t\tCategory: Sorting\n");
  hr();

  hr();
  measure(sort_in_one_loop);
  hr();
  return 0;
}

void hr()
{
  printf("============================================================\n");
}

void measure(void *f())
{
  (*f)();
}
