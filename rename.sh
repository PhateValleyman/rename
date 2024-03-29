#!/data/data/com.termux/files/usr/bin/bash

B="\e[01;34m" # Blue color
G="\e[32m" # Green color
R="\e[31m" # Red color
Y="\e[01;33m" # Yellow color
N="\e[0m" # Reset colors

# Function to display usage information
show_help() {
	echo -e "Usage: rename <${Y}filename${N}>"
}

# Function to display version information
show_version() {
	echo -e "${Y}rename${N} v1.1"
	echo -e "by ${Y}PhateValleyman${N}"
	echo -e "${G}Jonas.Ned@outlook.com${N}"
}

# Function to rename a file
rename_file() {
	local old_filename="$1"
	local new_filename="$2"

	if [ -e "$new_filename" ]; then
		echo -e "\"${Y}${new_filename}${N}\""
		read -p "already exists. Do you want to overwrite it? (y/n): " overwrite
		case "$overwrite" in
			[Yy]*)
				if mv -f -- "$old_filename" "$new_filename"; then
					echo -e "File \"${Y}${old_filename}${N}\" renamed to \"${G}${new_filename}${N}\""
				else
					echo -e "Failed to rename file \"${R}${old_filename}${N}\""
					return 1
				fi
				;;
			*)
				echo -e "Rename ${Y}aborted${N}..."
				return 1
				;;
		esac
	else
		if mv -f -- "$old_filename" "$new_filename"; then
			echo -e "File \"${Y}${old_filename}${N}\" renamed to \"${G}${new_filename}${N}\""
		else
			echo -e "Failed to rename file \"${R}${old_filename}${N}\""
			return 1
		fi
	fi
}

# Function to process command line arguments
process_arguments() {
	case "$1" in
		--help | -h | -?)
			show_help
			;;
		--version | -v)
			show_version
			;;
		*)
			if [ "$#" -lt 1 ] || [ ! -e "$1" ]; then
				show_help
				return 1
			fi
			local old_filename="$1"
			local new_filename
			read -ei "$old_filename" new_filename
			rename_file "$old_filename" "$new_filename"
			;;
	esac
}

# Call the function to process command line arguments
process_arguments "$@"
