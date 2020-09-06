# !/bin/bash
src=$1
dst=$2
find app -type f | xargs -I {} sed -i "" -e 's/$src/$dst/g' {}
