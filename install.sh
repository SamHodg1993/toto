#!/bin/bash

# Build the binary with the tool name "toto" no, this isn't a typo lol!
go build -o toto

# Create ~/bin if it doesn't exist
mkdir -p ~/bin

# Move the binary to ~/bin
mv toto ~/bin/

# Move the todo database file to home directory if it exists
if [ -f "todos.db" ]; then
    mv todos.db ~/.todos.db
else
    # If no database exists, create an empty one in home directory
    touch ~/.todos.db
fi

# Check if PATH already contains ~/bin
if [[ ":$PATH:" != *":$HOME/bin:"* ]]; then
    # Add PATH update to ~/.bashrc if not already there
    echo 'export PATH="$HOME/bin:$PATH"' >> ~/.bashrc
    # Also update current session's PATH
    export PATH="$HOME/bin:$PATH"
    echo "Added ~/bin to PATH"
fi

# Source the bashrc
source ~/.bashrc
echo "Installation of the toto todo package has been completed! Enjoy!"
