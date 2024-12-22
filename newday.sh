#!/bin/sh

# Determine the day number
if [ -z "$1" ]; then
  DAY=$(date +"%d") # Use today's day if no argument is provided
else
  DAY=$(printf "%02d" "$1") # Ensure day is zero-padded
fi
year=$(date +"%Y")

# Create necessary directories
mkdir -p "internal/day$DAY" "puzzles/day$DAY"

# check if the day already exists
ALREADY_EXISTS=$(ls internal/day$DAY 2>/dev/null)

# Create files with safety checks
touch "internal/day$DAY/day${DAY}.go" "puzzles/day$DAY/example1.txt" "puzzles/day$DAY/input1.txt"

# Populate the Go file with a template if it doesn't exist
if [ -z "$ALREADY_EXISTS" ]; then
NEW_DAY="internal/day$DAY/day${DAY}.go"
cp internal/day00/day00.go $NEW_DAY
sed -i "" "s/00/${DAY}/g" $NEW_DAY

# Format the Go file
gofmt -w "internal/day$DAY/day${DAY}.go"
fi


# Update cmd/main.go
MAIN_FILE="cmd/main.go"
IMPORT_STATEMENT="\\\"github.com/afonsocraposo/advent-of-code-${year}/internal/day${DAY}\\\""
DAYS_ENTRY="${DAY}: day${DAY}.Main,"

# Add import statement if not already present
if ! grep -q "$IMPORT_STATEMENT" "$MAIN_FILE"; then
  printf "%s\n" "$IMPORT_STATEMENT" | sed -i '' "/import (/a\\"$'\n'"$IMPORT_STATEMENT;" "$MAIN_FILE"
fi

# Add day entry to the map if not already present
if ! grep -q "$DAYS_ENTRY" "$MAIN_FILE"; then
  printf "%s\n" "$DAYS_ENTRY" | sed -i '' "/var days = map\\[int\\]func()/a\\"$'\n'"$DAYS_ENTRY"$'\n'" " "$MAIN_FILE"
fi

gofmt -w "$MAIN_FILE"

# Notify the user
echo "Setup completed for Day $DAY"
