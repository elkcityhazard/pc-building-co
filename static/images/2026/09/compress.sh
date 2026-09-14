#!/usr/bin/env bash

# Output directory
out_dir="output"
mkdir -p "$out_dir"

# Loop through all JPG files in the current working directory
for img in *.[jJ][pP][gG]; do
  # Skip if no matching files exist
  [ -e "$img" ] || continue

  # Extract base filename without extension or path
  filename="$(basename "$img")"
  base="${filename%.*}"
E
  # 1. Large (1920x1080 canvas, height max 1080)
  magick "$img" -auto-orient -resize x1080 \
    -background transparent -gravity center -extent 1920x1080 \
    -quality 80 -define webp:method=6 \
    "$out_dir/${base}_1920x1080.webp"

  # 2. Medium (1280x720 canvas, height max 720)
  magick "$img" -auto-orient -resize x720 \
    -background transparent -gravity center -extent 1280x720 \
    -quality 80 -define webp:method=6 \
    "$out_dir/${base}_1280x720.webp"

  # 3. Small / Thumbnail (640x360 canvas, height max 360)
  magick "$img" -auto-orient -resize x360 \
    -background transparent -gravity center -extent 640x360 \
    -quality 80 -define webp:method=6 \
    "$out_dir/${base}_640x360.webp"

  echo "Processed: $img -> $out_dir/"
done

echo "Done! All images saved to $out_dir/"