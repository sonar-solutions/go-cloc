#!/bin/bash

# clean build directory
rm -rf builds/

# remove unused packages
go mod tidy

# Set the output directory and binary name
OUTPUT_DIR="builds"
BINARY_NAME="go-cloc" # Change this to your program's name
mkdir -p $OUTPUT_DIR

# Define platforms and architectures
platforms=(
    "linux/amd64"
    "linux/arm64"
    "darwin/amd64"
    "darwin/arm64"
    "windows/amd64"
    "windows/386"
)

# Build for each platform and create a zip for each
for platform in "${platforms[@]}"; do
    IFS="/" read -r os arch <<< "$platform"
    
    if [ "$os" == "windows" ]; then
        output_file="$OUTPUT_DIR/$BINARY_NAME.exe"
    else
        output_file="$OUTPUT_DIR/$BINARY_NAME"
    fi
    
    echo "Building for $os/$arch..."
    GOOS=$os GOARCH=$arch go build -o "$output_file" main.go
    
    if [ $? -ne 0 ]; then
        echo "Build failed for $os/$arch"
        exit 1
    fi
  
    # Create a zip file for the current build
    zip_file="$OUTPUT_DIR/${BINARY_NAME}-${os}-${arch}.zip"
    zip -j "$zip_file" "$output_file"
    
    # Optionally remove the binary after zipping
    rm "$output_file"
    
    echo "Created zip: $zip_file"
done

echo "All builds are zipped in the $OUTPUT_DIR directory."
