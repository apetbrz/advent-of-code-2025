#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main() {
    FILE *fptr = fopen("sequence.txt", "r");

    char buff[128];

    int sum = 0;

    while (fgets(buff, 128, fptr)) {
        int tens = 0;
        int ones = 0;
        int temp = 0;
        char *largest = buff;
        char *ptr = buff;

        //loop through first n-1 digits
        while(*(ptr + 1) != '\n'){
            temp = (int)(*ptr - '0');
            if(temp > tens) {
                tens = temp;
                largest = ptr;
            }
            ptr = ptr + 1;
        }
        ptr = largest + 1;

        //loop from digit after ^^ to end
        while(*ptr != '\n'){
            temp = (int)(*ptr - '0');
            if(temp > ones) ones = temp;
            ptr = ptr + 1;
        }

        sum = sum + (tens * 10) + ones;
    }

    printf("%i\n", sum);

    return 0;
}
