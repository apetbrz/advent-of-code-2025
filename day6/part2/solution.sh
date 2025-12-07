#!/usr/bin/env bash

vals=($(head -n 1 sequence.txt))

ops=()
read -a ops <<<"$(tail -n 1 sequence.txt)"

for index in "${!vals[@]}"; do
    ones=$((vals[index] % 10))
    tens=$(((vals[index] / 10) % 10))
    hundreds=$(((vals[index] / 100) % 10))
    thousands=$((vals[index] / 1000))
    while read line; do
        line=($line)

        # ahhhhhhh this doesnt work :P
        # i tried real hard but this problem looks really hard for bash
        # the input numbers are not right justified....
        #
        # i.e.
        #
        # how it needs to work:
        #
        #        1.   2.  3.
        # 328    |3| |2| |8|
        # 64  -> |6| |4|     -> 369 + 248 + 8
        # 98     |9| |8|
        #
        # what this script does:
        #
        #       hndrds tens ones
        # 328     3     2    8
        # 64  ->        6    4  -> 3 + 269 + 848
        # 98            9    8
        #
        # i fear i'm not quite sure how best to do this in pure bash...
        # a reminder as to why you pick the best tool for the job...
        # and not just whichever one you feel like picking...

        o=$((line[index] % 10))
        t=$(((line[index] / 10) % 10))
        h=$(((line[index] / 100) % 10))
        th=$((line[index] / 1000))
        [[ $o -eq 0 ]] || ones="${ones}${o}"
        [[ $t -eq 0 ]] || tens="${tens}${t}"
        [[ $h -eq 0 ]] || hundreds="${hundreds}${h}"
        [[ $th -eq 0 ]] || thousands="${thousands}${th}"
    done < <(tail -n +2 sequence.txt | head -n -1)
    case "${ops[$index]}" in
    "+")
        vals[$index]=$((ones + tens + hundreds + thousands))
        ;;
    "*")
        vals[$index]=$(((ones == 0 ? 1 : ones) * (tens == 0 ? 1 : tens) * (hundreds == 0 ? 1 : hundreds) * (thousands == 0 ? 1 : thousands)))
        ;;
    esac
    echo "${ones} ${tens} ${hundreds} ${thousands} ${ops[$index]} = ${vals[$index]}"
done

echo "$(($(echo ${vals[@]} | sed -e 's/ /+/g')))"
