Bash
-

````
bash --version

PATH=$PATH:~/bin

type cp
type /Users/k/web/kovpak/gh/ed/bash/examples/hw.sh
````

````bash
$1 # 1st script parameter
$* # all script parameters
$# # number of script parameters
$? # exit status for last command

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
