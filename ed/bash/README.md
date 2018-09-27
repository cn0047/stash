Bash
-

````
bash --version
echo $BASH_VERSION

PATH=$PATH:~/bin

type cp
type /Users/k/web/kovpak/gh/ed/bash/examples/hw.sh
````

````bash
$0   # name of called script
$1   # 1st script parameter
$2   # 2nd
"$@" # all script parameters (+ quotes)
$*   # all script parameters
$#   # number of script parameters
$?   # exit status for last command

# shift:
$2 -> $1
$3 -> $2
$4 -> $3

echo $(date)
echo '$myVar' # not evaluate vars
echo "$myVar" # evaluate
echo $USER

read -p "Your note: " note

if [[ $str ]];           # str isn't empty
if [[ $str = "txt" ]];   # str equals "txt"
if [[ $str="txt" ]];     # always true
if [[ $str = [Yy] ]];    # Y || y
if [[ ! $1 ]];           # $1 is empty
if [[ -e $file ]];       # file exists
if [[ -d $dir ]];        # is directory
if [[ $1 =~ ^[0-9]+$ ]]; # is number
&& # &
|| # or

exit 0 # success
exit 1 # fail

# set Input Field Separator, by default ` ` (space)
IFS=:

[[ $1 ]] || { echo "missing argument" >&2; exit 1; }
{ cat x.txt || echo "file x.txt not found"; } 2>/dev/null
````

#### Debug:

````bash
#!/bin/bash -x

# or
set -x # to start debug
set +x # to end debug
````

#### Strings:

````bash
${#var} # string length
````

#### Numbers:

````bash
-eq
-ne
-le
-gt
# don't use =,<,> for numbers

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

arr=(1 2 3 a b)
echo ${#arr[@]} # array length
echo ${!arr[@]} # array indices
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
````

#### Loop:

````bash
break
continue

while test; do
    ;; code
done

until test; do
    ;; code
done

for el in arr; do
    ;; code
done
for f in *"$1"; do
    echo "$f $1"
done
for (( init; test; update )); do
    ;; code
done
````
