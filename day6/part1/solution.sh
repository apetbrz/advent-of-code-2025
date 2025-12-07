#!/usr/bin/env bash

vals=($(head -n 1 sequence.txt))

ops=()
read -a ops <<<"$(tail -n 1 sequence.txt)"

while read line; do
    line=($line)
    for index in "${!line[@]}"; do
        case "${ops[$index]}" in
        "+")
            vals[$index]=$((vals[$index] + line[$index]))
            ;;
        "*")
            vals[$index]=$((vals[$index] * line[$index]))
            ;;
        esac
    done
done < <(tail -n +2 sequence.txt | head -n -1)

echo "$(($(echo ${vals[@]} | sed -e 's/ /+/g')))"
