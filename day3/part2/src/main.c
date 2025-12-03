#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main() {
    FILE *fptr = fopen("sequence.txt", "r");

    char buff[128];

    u_int64_t sum = 0;

    while (fgets(buff, 128, fptr)) {
        u_int64_t num = 0;
        int digits_remaining = 12;
        char *largest = buff;
        char *ptr = buff;

        while(digits_remaining > 0){
            //loop through first n-(digits_remaining) digits
            while(*(ptr + digits_remaining - 1) != '\n'){
                if((int)(*ptr - '0') > (int)(*largest - '0')) largest = ptr;
                ptr++;
            }
            //push digit onto final number
            num = num * 10 + (int)(*largest - '0');
            //start next iter from after previous digit
            ptr = largest + 1;
            largest = ptr;

            digits_remaining--;
        }

        sum = sum + num;
    }

    printf("%li\n", sum);

    return 0;
}
