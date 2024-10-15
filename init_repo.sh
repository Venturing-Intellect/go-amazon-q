#!/bin/bash

# Check if a directory name was provided
if [ $# -eq 0 ]; then
    echo "Please provide a directory name for your Go project."
    exit 1
fi

# Create the main project directory
mkdir -p "$1"
cd "$1" || exit

# Initialize git repository
git init

# Create .gitignore file with Go template
cat << EOF > .gitignore
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with 'go test -c'
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Dependency directories (remove the comment below to include it)
# vendor/

# Go workspace file
go.work

# IDE-specific files
.idea/
.vscode/

# OS-specific files
.DS_Store
Thumbs.db
EOF

# Create main.go file
cat << EOF > main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
EOF

# Create go.mod file
go mod init "$(basename "$(pwd)")"

# Create directories
mkdir -p middleware controller repository service model

# Create placeholder files in each directory
touch middleware/.gitkeep controller/.gitkeep repository/.gitkeep service/.gitkeep model/.gitkeep

echo "Golang repository created successfully with the following structure:"
tree -L 2
echo "Don't forget to run 'go mod tidy' to manage your dependencies."
