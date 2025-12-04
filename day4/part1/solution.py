#!/usr/bin/env python3

file = open("sequence.txt", "r")
content = file.read()
content = content.split()
file.close()


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
print(accessible)
