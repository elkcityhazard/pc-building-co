#!/usr/bin/env bash

# Output directory
mkdir -p output

# Loop through all .jpg files (case-insensitive)
for img in *.[jJ][pP][gG]; do
  # Skip if no matching files are found
  [ -e "$img" ] || continue

  filename="${img%.*}" # Extracts filename without extension
  ext="${img##*.}"     # Keeps original extension

  # Large (1920x1080)
  magick "$img" -auto-orient -resize x1080 \
    -background transparent -gravity center -extent 1920x1080 \
    "output/${filename}_1920x1080.${ext}"

  # Medium (1280x720)
  magick "$img" -auto-orient -resize x720 \
    -background transparent -gravity center -extent 1280x720 \
    "output/${filename}_1280x720.${ext}"

  # Small (640x360)
  magick "$img" -auto-orient -resize x360 \
    -background transparent -gravity center -extent 640x360 \
    "output/${filename}_640x360.${ext}"

  echo "Processed: $img"
done

echo "Done! Check the output/ folder."