#!/bin/bash

convertToRoman () {
if   [ $(($1-1000)) -ge 0 ]; then v2=`convertToRoman $(($1-1000))`; echo "M$v2";
elif [ $(($1-900)) -ge 0 ];  then v2=`convertToRoman $(($1-900))`;  echo "CM$v2";
elif [ $(($1-500)) -ge 0 ];  then v2=`convertToRoman $(($1-500))`;  echo "D$v2";
elif [ $(($1-400)) -ge 0 ];  then v2=`convertToRoman $(($1-400))`;  echo "CD$v2";
elif [ $(($1-100)) -ge 0 ];  then v2=`convertToRoman $(($1-100))`;  echo "C$v2";
elif [ $(($1-90)) -ge 0 ];   then v2=`convertToRoman $(($1-90))`;   echo "XC$v2";
elif [ $(($1-50)) -ge 0 ];   then v2=`convertToRoman $(($1-50))`;   echo "L$v2";
elif [ $(($1-40)) -ge 0 ];   then v2=`convertToRoman $(($1-40))`;   echo "XL$v2";
elif [ $(($1-10)) -ge 0 ];   then v2=`convertToRoman $(($1-10))`;   echo "X$v2";
elif [ $(($1-9)) -ge 0 ];    then v2=`convertToRoman $(($1-9))`;    echo "IX$v2";
elif [ $(($1-5)) -ge 0 ];    then v2=`convertToRoman $(($1-5))`;    echo "V$v2";
elif [ $(($1-4)) -ge 0 ];    then v2=`convertToRoman $(($1-4))`;    echo "IV$v2";
elif [ $(($1-1)) -ge 0 ];    then v2=`convertToRoman $(($1-1))`;    echo "I$v2";
else echo ''; fi
}

convertToRoman 4154
