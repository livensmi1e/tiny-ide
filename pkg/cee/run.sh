#!/bin/sh

FILE=$1

if [ -z "$FILE" ]; then
    echo "error: file not found"
    exit 1
fi

EXT="${FILE##*.}"

START_TIME=$(date +%s%N)
TIME_OUTPUT=""
MEMORY_OUTPUT=""
STDOUT_OUTPUT=""
STDERR_OUTPUT=""

case "$EXT" in
    py)
        {
            STDOUT_OUTPUT=$(python3 "$FILE"; printf '.'; exit "$?")
        } 2>/sandbox/code/main.stderr
        STDOUT_OUTPUT=${STDOUT_OUTPUT%.}
        STDERR_OUTPUT=$(cat /sandbox/code/main.stderr)
        ;;
    c)
        gcc "$FILE" -o /sandbox/code/main.out 2>/sandbox/code/main.stderr
        if [ $? -eq 0 ]; then
            {
                STDOUT_OUTPUT=$(/sandbox/code/main.out; printf '.'; exit "$?")
            } 2>/sandbox/code/main.stderr
            STDOUT_OUTPUT=${STDOUT_OUTPUT%.}
            STDERR_OUTPUT=$(cat /sandbox/code/main.stderr)
        else
            STDERR_OUTPUT=$(cat /sandbox/code/main.stderr)
        fi
        ;;
    cpp)
        g++ "$FILE" -o /sandbox/code/main.out 2>/sandbox/code/main.err
        if [ $? -eq 0 ]; then
            {
                STDOUT_OUTPUT=$(/sandbox/code/main.out; printf '.'; exit "$?")
            } 2>/sandbox/code/main.err
            STDOUT_OUTPUT=${STDOUT_OUTPUT%.}
            STDERR_OUTPUT=$(cat /sandbox/code/main.err)
        else
            STDERR_OUTPUT=$(cat /sandbox/code/main.err)
        fi
        ;;
    *)
        STDERR_OUTPUT="unsupported file type: $EXT"
        ;;
esac

END_TIME=$(date +%s%N)
TIME_DIFF=$((($END_TIME - $START_TIME) / 1000000))

MEMORY_USAGE=$(ps -o rss= -p $$)

printf "stdout: %s\n" "$STDOUT_OUTPUT"
printf "stderr: %s\n" "$STDERR_OUTPUT"
printf "time: %d ms\n" "$TIME_DIFF"
printf "memory: %s kb\n" "$MEMORY_USAGE"