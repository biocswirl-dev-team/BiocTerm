#!/usr/bin/env bash

####################################################

#FILE         : yaml_parse.sh
#DESCRIPTION  : This bash script parses biocswirl course yaml files and imports them as bash variables

#OPTIONS      : ---
#REQUIREMENTS : ---
#BUGS         : ---
#NOTES        : ---

#AUTHOR       : Lisa N. Cao 
#CONTACT      : ---
#DATE CREATED : 05/23/2020
#LAST REVISION: 

####################################################

# source scripts
source yaml.sh


# debug purposes
set -a -v -e


# pulled from jasperes test.sh
DEBUG="$1"

function is_debug() {
    [ "$DEBUG" = "--debug" ] && return 0 || return 1
}

if is_debug; then
    parse_yaml file.yml && echo
fi


# select lesson 
read -p "Specify Lesson YAML: " lesson
echo $lesson


# Execute 
create_variables $lesson

