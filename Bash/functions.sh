#!/bin/bash

echo "---Task 1: Hello World---"
function hi(){
    echo "Hello World!"
}

hi



echo -e "\n---Task 2: Percentages---"

# Usage: percent <number>
function percent(){
    if [ $(echo $1 | grep "^-\?[0-9]*$") ]
    then
        echo 'Percentage: '$(echo "scale=3;$1/100" | bc)'%'
    else
        echo "Invalid Number Used"
    fi
}

percent 238



echo -e "\n---Task 3: Caesar Cipher---"

# Usage: caesar <shift number> <word>
function caesar(){

    var1=""
    shift=$(expr $1 % 26)
    input=$2

    echo "Caesar Cipher: $(for ((i=0; i<${#input}; i++))
        do
            Char=${input:$i:1}
            if [[ ${Char} == [A-Z] ]]
            then
                var1=$var1$(echo $Char | tr $(printf %${shift}s | tr " " ".")A-Z A-ZA-Z)
            elif [[ ${Char} == [a-z] ]]
            then 
                var1=$var1$(echo $Char | tr $(printf %${shift}s | tr " " ".")a-z a-za-z)
            else
                var1=${var1}${Char}
            fi
        done
        echo ${var1}
    )"
}

caesar 1 "Text Manipulation"