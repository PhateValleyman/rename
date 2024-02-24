#!/data/data/com.termux/files/usr/bin/bash

G="\e[32m" # Green color
R="\e[31m" # Red color
Y="\e[01;33m" # Yellow color
N="\e[0m" # Reset colors

# Function to display usage information
show_help() {
	echo -e "Usage: rename <${Y}filename${N}/${Y}s${N}>"
}

# Function to display version information
show_version() {
	echo "rename v1.1"
	echo "by PhateValleyman"
	echo "Jonas.Ned@outlook.com"
}

# Function to rename a file
rename_file() {
	local old_filename="$1"
	local new_filename="$2"

	if [ -e "$new_filename" ]; then
		dialog --backtitle "Rename File" --title "File Exists" --yes-label "Overwrite" --no-label "Rename" --cancel-label "Cancel" --yesno "\"${new_filename}\" already exists. Do you want to:" 10 40

		case $? in
			0) # Overwrite
				if mv -f -- "$old_filename" "$new_filename"; then
					echo -e "File \"${Y}${old_filename}${N}\" renamed to \"${G}${new_filename}${N}\""
				else
					echo -e "Failed to rename file \"${R}${old_filename}${N}\""
					return 1
				fi
				;;
			1) # Rename
				local new_name
				dialog --backtitle "Rename File" --title "Rename File" --inputbox "Enter new filename:" 10 40 2> /data/data/com.termux/files/usr/tmp/rename.tmp
				new_name=$(cat /data/data/com.termux/files/usr/tmp/rename.tmp)
				rm /data/data/com.termux/files/usr/tmp/rename.tmp
				if [ -n "$new_name" ]; then
					rename_file "$old_filename" "$new_name"
				else
					echo -e "Rename ${R}aborted${N}..."
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
			dialog --backtitle "Rename File" --title "Rename File" --inputbox "Enter new filename:" 10 40 2> /data/data/com.termux/files/usr/tmp/rename.tmp
			new_filename=$(cat /data/data/com.termux/files/usr/tmp/rename.tmp)
			rm /data/data/com.termux/files/usr/tmp/rename.tmp
			if [ -n "$new_filename" ]; then
				rename_file "$old_filename" "$new_filename"
			else
				echo -e "Rename ${R}aborted${N}..."
				return 1
			fi
			;;
	esac
}

# Call the function to process command line arguments
process_arguments "$@"
