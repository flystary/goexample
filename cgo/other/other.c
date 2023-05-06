#include <stdio.h>
#include <stdlib.h>

int add(int a, int b)
{
    return a + b;
}

struct Point
{
    int x;
    int y;
};

int arean(struct Point a, struct Point b)
{
    int length = a.x - b.y;
    int width = a.y - b.y;

    return length * width;
}

int arean_point(struct Point *a, struct Point *b)
{
    int length = a->x - b->y;
    int width = a->y - b->y;

    return length * width;
}
