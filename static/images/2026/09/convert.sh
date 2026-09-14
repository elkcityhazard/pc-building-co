#!/usr/bin/env bash

# Output directory
out_dir="output"
mkdir -p "$out_dir"

# Loop through all JPG files in the current working directory
for img in *.[jJ][pP][gG]; do
  # Skip if no matching files exist
  [ -e "$img" ] || continue

  # Extract base filename without extension
  filename="$(basename "$img")"
  base="${filename%.*}"

  # 1. Large (Max 1920x1080, preserves aspect ratio, padded with transparency)
  magick "$img" -auto-orient -resize 1920x1080 \
    -background transparent -gravity center -extent 1920x1080 \
    -quality 80 -define webp:method=6 \
    "$out_dir/${base}_1920x1080.webp"

  # 2. Small / Thumbnail (360x360 square, crop-to-fill)
  magick "$img" -auto-orient -resize 360x360^ \
    -gravity center -extent 360x360 \
    -quality 80 -define webp:method=6 \
    "$out_dir/${base}_360x360.webp"

  echo "Processed: $img -> $out_dir/"
done

echo "Done! All images processed and saved to $out_dir/"