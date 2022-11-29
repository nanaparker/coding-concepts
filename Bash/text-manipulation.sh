#!/bin/bash

# Usage: 
# ./text-manipulation.sh <cipher shift>

var1="Text Manipulation"
echo -e "String: ${var1}\n" 


echo "---Task 1: Substrings---"
echo "Character Count: ${#var1}" # Checking number of characters
echo "First 3 Letters: ${var1:0:3}"
echo "Second Word: $(echo ${var1} | cut -d ' ' -f 2)"



echo -e "\n---Task 2: Character Cases---"
echo "Lower Case: $(echo ${var1} | tr -s '[:upper:]' '[:lower:]' )"
echo "Upper Case: $(echo ${var1} | tr -s '[:lower:]' '[:upper:]' )"

var2=""
echo "Mixed Case: $(for i in $(seq 0 ${#var1})
	do
		if [[ ${var1:$i:1} == ' ' ]]
		then
			var2="${var2} "
		fi
	
		if [ $(expr ${i} % 2) == 1 ]
		then
			var2="${var2}$(echo ${var1:$i:1} | tr -s '[:upper:]' '[:lower:]')"
		else
			var2="${var2}$(echo ${var1:$i:1} | tr -s '[:lower:]' '[:upper:]')"
		fi
	done 
	echo ${var2}
)"



echo -e "\n---Task 3: Ciphers---"

var3=""
shift=$(expr $1 % 26)
echo "Caesar Cipher: $(for ((i=0; i<${#var1}; i++))
	do
		Char=${var1:$i:1}
		if [[ ${Char} == [A-Z] ]]
		then
			var3=$var3$(echo $Char | tr $(printf %${shift}s | tr " " ".")A-Z A-ZA-Z)
		elif [[ ${Char} == [a-z] ]]
		then 
			var3=$var3$(echo $Char | tr $(printf %$1s | tr " " ".")a-z a-za-z)
		else
			var3=${var3}${Char}
		fi
	done
	echo ${var3}
)"
