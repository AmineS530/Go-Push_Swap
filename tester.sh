#!/bin/bash

if [ $# -eq 0 ]; then
echo "Usage: $0 <numbers>"
exit 1
fi

NUMBERS="$*"
MOVES_FILE="moves.txt"

# Generate instructions

./push-swap "$NUMBERS" > "$MOVES_FILE"

# Count instructions

COUNT=$(wc -l < "$MOVES_FILE")

# Run checker using the generated instructions

RESULT=$(cat "$MOVES_FILE" | ./checker "$NUMBERS")

# Append summary to file

{
echo ""
echo "----- RESULT -----"
echo "Instructions: $COUNT"
echo "Checker: $RESULT"
} >> "$MOVES_FILE"

echo "Instructions: $COUNT"
echo "Checker: $RESULT"
echo "Saved to $MOVES_FILE"
