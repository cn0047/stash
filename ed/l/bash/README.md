Bash
-

[source code](http://git.savannah.gnu.org/cgit/coreutils.git/tree/src/comm.c)

Variables in bash are untyped, but bash has:
* number
* string
* array
* function

````sh
bash --version
echo $BASH_VERSION
echo There are ${#BASH_ALIASES[*]} aliases defined.

bash -c 'echo ok' # cmd
bash -ce 'ok'     # cmd and fail if any error
bash -i           # interactive
bash -l           # login
bash --norc       # don't read file ~/.bashrc
bash --noprofile  # don't read files ~/.bash_profile, ~/.bash_login, or ~/.profile
bash --init-file  # file instead of ~/.bashrc

zero exit code     - ok
non zero exit code - error
exit 0 # success
exit 1 # fail

command1 && command2 # command2 is executed if, and only if,
                     # command1 returns an exit status of zero.
command1 || command2 # command2 is executed if and only if
                     # command1 returns a non-zero exit status.
&& # &
|| # or

: # do nothing

PATH=$PATH:~/bin

# shabang:
#!/bin/bash
#!/usr/bin/env php
#!/usr/bin/env python3

# set Input Field Separator, by default ` ` (space)
IFS=:

$0   # name of called script
$1   # 1st script parameter
$2   # 2nd
!$   # last argument
"$@" # all script parameters (+ quotes)
$*   # all script parameters
$#   # number of script parameters
$?   # exit status for last command
$!   # PID of last background task
$$   # PID of shell

# shift:
$2 -> $1
$3 -> $2
$4 -> $3

# Redirects:
echo yes 0> f   # for standard input
echo ok 1> f    # for standard output
echo no 2>| f   # for errors
echo okay &> f  # for standard output & errors
echo OK >| f # overwrite file

# also
echo "test" 1>/dev/null 2>&1 &
1>&2 # output to stderr
2>&1 # stderr to stdout
>&2 # output to stderr

cd "$(dirname "$(readlink -f "$0")")"

cmd || echo 'cmd failed'
docker info 2>/dev/null || echo 'fail'
test-d$HOME/.kube || mkdir$HOME/.kube

read -p "Your note: " note # prompt

# process with pid
if [[ ! -e /tmp/test.py.pid ]]; then
  python test.py &
  echo $! > /tmp/test.py.pid
else
  echo -n "ERROR: The process is already running with pid "
  cat /tmp/test.py.pid
  echo
fi

check_cmd_exists() {
  if ! which $1 &>/dev/null; then
    error "$1 command not found, you must install it."
  fi
}
````

````bash
echo '3^2 * 10 / 23' | bc -l # math

  if [[ -f $filename ]]; then echo "$filename is a regular file"
elif [[ -e $filename ]]; then echo "$filename is exists"
elif [[ -d $filename ]]; then echo "$filename is a directory"
elif [[ -p $filename ]]; then echo "$filename is a named pipe"
elif [[ -S $filename ]]; then echo "$filename is a named socket"
elif [[ -b $filename ]]; then echo "$filename is a block device"
elif [[ -c $filename ]]; then echo "$filename is a character device"
elif [[ -L $filename ]]; then echo "$filename is a symbolic link (to any file type)"
elif [[ -r $filename ]]; then echo "$filename is a readable file"
elif [[ -w $filename ]]; then echo "$filename is a writable file"
elif [[ -x $filename ]]; then echo "$filename is an executable file"

if [[ -n "$string" ]];      then echo "$string is not empty"
if [[ -z "${string// }" ]]; then echo "$string is empty or contains only spaces"
if [[ -z "$string" ]];      then echo "$string is empty"

-e "$file"        # file exists
-d "$file"        # directory exists
-f "$file"        # regular file exists
-h "$file"        # symbolic link exists
-z "$str"         # length is zero
-n "$str"         # length is non-zero
"$str" = "$str2"  # equal strings (NOT for integer)
"$int1" -eq "$int2"
"$int1" -ne "$int2"
"$int1" -gt "$int2"
"$int1" -ge "$int2"
"$int1" -lt "$int2"
"$int1" -le "$int2"

[[ hello = h*o ]] && echo yes
[[ heeello =~ (e+) ]] && echo "yes, because: ${BASH_REMATCH[1]}"
[[ $1 ]] || { echo "missing argument" >&2; exit 1; }
{ cat x.txt || echo "file x.txt not found"; } 2>/dev/null

local val=${1:?Must provide an argument}

# default value
declare y=${myDefVar:-"nil"}
echo $y # nil
myDefVar=null
declare y=${myDefVar:-"nil"}
echo $y # null

declare -i # interger
declare -r # readonly
declare -x # export

# end of options:
touch -a # error
touch -- -a # ok

set -e           # exit whenever a command fails
set -n           # validate but not exec script
set -o           # display options
set -o noclobber # to enable option
set +o noclobber # to disable option
set -o allexport #
set -u           # error when using uniinitialized var
set -v           #
set -v           # print each command
set -x           # to start debug

echo "${v1:-7}" # 7
v2=22
echo "${v2:-7}" # 22
echo "${v3:=7}" # 7

# v4="${v4:?EMPTY}"; echo $v4

# echo "FOO is ${FOO:?EMPTY}"
# echo "FOO is ${FOO?UNSET}"

# ${param:-word}
# ${param-word}
# ${param:=word}
# ${param=word}
# ${param:?word}
# ${param?word}
# ${param:+word}
# ${param+word}

cat <<TXT
---------
204
---------
TXT

cat << EOF >> /tmp/yourfilehere
These contents will be written to the file.
        This line is indented.
EOF
````

#### Debug:

````bash
#!/bin/bash -x

# or
set -x # to enable debug
set +x # to disable debug

# or
bash -x /Users/k/web/kovpak/gh/ed/bash/examples/hw.sh
````

#### Strings:

````bash
v="hello world"
declare -u s="$v"; echo $s # HELLO WORLD
declare -l s="$v"; echo $s # hello world

var='12345'
echo "${#var}" # length

s="http://host/json_path.json"
echo ${s#htt?}        # ://host/json_path.json
echo ${s#*/}          # /host/json_path.json
echo ${s##*/}         # json_path.json
echo ${s%.*}          # http://host/json_path
echo ${s%/*}          # http://host
echo ${s%%/*}         # http:
echo ${s/json/yaml}   # http://host/yaml_path.json
echo ${s/%json/yaml}  # http://host/json_path.yaml
echo ${s/.json/}      # http://host/json_path
echo ${s%.json}       # http://host/json_path
echo ${s//[o]/X}      # http://hXst/jsXn_path.jsXn

str='I am a string'
echo "${str/a/A}"  # I Am a string
echo "${str//a/A}" # I Am A string
echo "${str/#I/=}" # = am a string
echo "${str/%g/N}" # I am a strinN
echo "${str/g/}"   # I am a strin # replace with nothing
echo "${str%a*}"   # I am
echo "${str#*a}"   # m a string
echo "${str##*a}"  # string

FILENAME="/tmp/example/myfile.txt"
echo "${FILENAME%/*}"    # /tmp/example
echo "${FILENAME##*/}"   # myfile.txt
BASENAME="${FILENAME##*/}"
echo "${BASENAME%%.txt}" # myfile

A=(hello world)
echo "${A[@]/#/R}" # Rhello Rworld

v="hello"
printf '%s\n' "${v^}" # Hello
printf '%s\n' "${v^^}" # HELLO

v="BYE"
printf '%s\n' "${v,}" # bYE
printf '%s\n' "${v,,}" # bye

v="Hello World"
echo "${v~}" # hello World
echo "${v~~}" # hELLO wORLD

v=foo-bar-baz
echo ${v%%-*} # foo
echo ${v%-*}  # foo-bar
echo ${v##*-} # baz
echo ${v#*-}  # bar-baz

var='0123456789abcdef'
printf '%s\n' "${var:3}"      # 3456789abcdef
printf '%s\n' "${var:3:4}"    # 3456
printf '%s\n' "${var:3:-5}"   # 3456789a
printf '%s\n' "${var: -6}"    # abcdef
printf '%s\n' "${var: -6:-5}" # a
printf '%s\n' "${var: -6:3}"  # abc

set -- 0123456789abcdef
printf '%s\n' "${1:5}" # 56789abcdef

if [[ $str ]];           # str isn't empty
if [[ $str = "txt" ]];   # str equals "txt"
if [[ $str="txt" ]];     # always true
if [[ $str = [Yy] ]];    # Y || y
if [[ $str == *.txt ]];  #
if [[ ! $1 ]];           # $1 is empty
if [[ $1 =~ ^[0-9]+$ ]]; # is number
````

#### Numbers:

````bash
-eq
-ne
-le
-gt
# ‼️ don't use =,<,> for numbers

declare -i p
p="4+5"
echo $p # 9
p="ok"
echo $p # 0
$((++p))
echo $p # 1

declare -ir const=$(( 1 + 1))
echo $const
const=3 # -bash: const: readonly variable

declare -x var="outer" # export var for inluded scripts
````

#### Array:

````bash
declare -a ar=("element1" "element2" "element3")
declare -p ar # declare -a ar='([0]="element1" [1]="element2" [2]="element3")'
array=([3]='fourth element' [4]='fifth element')

array+=('fourth element' 'fifth element')   # append
array=("new element" "${array[@]}")         # preppend
arr=("${arr[@]:0:$i}" 'new' "${arr[@]:$i}") # insert

arr=(1 2 3 a b)
echo ${#arr[@]} # array length
echo ${!arr[@]} # array indices

echo "${arr[@]: -1 }" # last el
echo "${array[-1]}" # last el
echo "${array[@]:1:3}"
echo "${array[@]/*[aA]*/}" # elements which match [aA] pattern
echo "${array[@]/[A-Z]/=}" # replace [A-Z] to =

arr=(a b c)
echo "${arr[@]}"  # outputs: a b c
echo "${!arr[@]}" # outputs: 0 1 2
unset -v 'arr[1]'
echo "${arr[@]}"  # outputs: a c
echo "${!arr[@]}" # outputs: 0 2

unset array[10]

stringVar="Apple+Orange+Banana+Mango"
arrayVar=(${stringVar//+/ })

myarr[0]='0123456789abcdef'
printf '%s\n' "${myarr[0]:7:3}" # 789

myarr=(0 1 2 3 4 5 6 7 8 9 a b c d e f)
printf '%s\n' "${myarr[@]:10}"
printf '%s\n' "${myarr[@]:5:3}"
printf '%s\n' "${myarr[@]: -1}"

set -- 1 2 3 4 5 6 7 8 9 0 a b c d e f
printf '%s\n' "${@:10}"
printf '%s\n' "${@:10:3}"
printf '%s\n' "${@:10:-2}"
printf '%s\n' "${@: -10:2}"
printf '%s\n' "${@:0:2}"
````

#### Function:

````bash
# in case no return in function
# return value will be return value from last function's line of code
startsWithA() {
  [[ $1 == [Aa]* ]];
}

error() {
  echo "Error: $1"
  exit 1
} >&2

foo() { bar; }
export -f foo
env | grep -A1 foo

export foo='() { echo "Inside function"; }'
bash -c 'foo'

export foo='() { echo "Inside function" ; }; echo "Executed echo"'
bash -c 'foo'

export dummy='() { echo "hi"; }; echo "pwned"'
$ bash
pwned
````

#### Loop:

````bash
break
continue

while true; do echo -n . ; done
while true; do echo -n . ; sleep 999; done
until false; do echo -n . ; done

# wait dir
while true ; do [[ -d '/tmp/x' ]] && break ; echo -n '.' ; sleep 1 ; done
# wait file by pattern
while true; do f=`find /tmp/ -name '*.txt.*' -type f`; [[ -n '$f' ]] && break ; echo -n '.' ; sleep 1 ; done

ls | grep $ptrn | while read f; do echo $f; done

for el in arr; do
    ;; code
done
for f in *"$1"; do
    echo "$f $1"
done
for i in $(seq 1 $end); do echo $i; done
for (( init; test; update )); do
    ;; code
done

break 2 # Break multiple loop
````
