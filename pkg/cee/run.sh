#!/bin/sh

FILE=$1

if [ -z "$FILE" ]; then
    echo "error: file not found"
    exit 1
fi

EXT="${FILE##*.}"

case "$EXT" in
    py)
        python3 "$FILE"
        ;;
    c)
        gcc "${FILE}" -o /sandbox/code/main.out && /sandbox/code/main.out
        ;;
    cpp)
        gcc "${FILE}" -o /sandbox/code/main.out && /sandbox/code/main.out
        ;;
    *)
        echo "unsupported file type: $EXT"
        exit 1
        ;;
esac