#!/data/data/com.termux/files/usr/bin/bash

case "$1" in
    -v | --version )
        echo -e "$0 v1.1"
        echo -e "by PhateValleyman"
        echo -e "Jonas.Ned@outlook.com"
        exit 0
        ;;
    -h | --help )
        echo -e "Usage: $0 [folder/file]"
        echo -e "Rename folder/file.\n"
        echo -e "Options:"
        echo -e "  -h, --help\t\tShow this help message."
        echo -e "  -v, --version\t\tDisplay script version."
        exit 0
        ;;
esac

# Function to ask for a name if not provided
_ask_for_name() {
    read -p "Enter a name: " name
}

# Function to rename a file
_rename_file() {
    read -p "Enter the file name to rename: " file
    read -p "Enter the new file name: " new_file

    # Rename the file
    mv "$file" "$new_file"
    echo -e "File renamed to\n'\e[01;33m$new_file\e[0m'"
}

# Check if a name is provided as an argument
if [[ -n $1 ]]; then
    # Rename the file
    _rename_file
else
    # Ask for a name
    _ask_for_name
fi
