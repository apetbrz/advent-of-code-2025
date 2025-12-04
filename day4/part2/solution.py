#!/usr/bin/env python3

import os

# i really liked how this looked with each iter
# wanted to keep the prints and add a delay per step
# run with SHOW_OUTPUT env var to true to see
display = os.environ.get( 'SHOW_OUTPUT' )

if(display):
    import time

file = open("sequence.txt", "r")
content = file.read()
if(display):
    print(content)
content = [list(string) for string in content.split()]
file.close()


removed = 0

while True:
    accessible = 0
    for i in range(0, len(content)):
        for j in range(0, len(content[i])):
            if(content[i][j] == '.'):
                continue
            around = 0
            for k in range(0,9):
                if(k==4):
                    continue
                x = j + int(k%3) - 1
                y = i + int(k/3) - 1
                if(x < 0 or y < 0):
                    continue
                try:
                    if(content[y][x] == "@"):
                        around += 1
                except IndexError:
                    continue
                    #do nothing
            if(around < 4):
                accessible += 1
                content[i][j] = '.'
    if(display):
        print("\n".join(["".join(line) for line in content]))
        time.sleep(0.1)
    if(accessible == 0):
        break
    removed += accessible

print(removed)
