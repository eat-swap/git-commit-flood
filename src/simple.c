#include <stdio.h>

// this file cannot be executed directly.

int main(void) {
    int x = %d;
    int i;
    for (i = 0; i < x; i++) {
        printf("%s\n");
    }
    return 0;
}