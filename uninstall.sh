#!/bin/bash

# Remove the binary
rm -f ~/bin/toto

# Remove the database file
read -p "Do you want to remove all your todos (delete ~/.todos.db)? (y/N) " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    rm -f ~/.todos.db
    echo "Todos database removed"
else
    echo "Keeping todos database at ~/.todos.db"
fi

# Remove the PATH addition from .bashrc
# Using sed to remove the exact export PATH line
sed -i '/export PATH="\$HOME\/bin:\$PATH"/d' ~/.bashrc

echo "Uninstall complete! You may want to restart your terminal"
