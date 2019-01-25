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
# shabang:
#!/bin/bash
#!/usr/bin/env php
#!/usr/bin/env python3

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
echo $RANDOM
echo $UID

read -p "Your note: " note

if [[ $str ]];           # str isn't empty
if [[ $str = "txt" ]];   # str equals "txt"
if [[ $str="txt" ]];     # always true
if [[ $str = [Yy] ]];    # Y || y
if [[ $str == *.txt ]];  #
if [[ ! $1 ]];           # $1 is empty
if [[ -e $file ]];       # file exists
if [[ -d $dir ]];        # is directory
if [[ $1 =~ ^[0-9]+$ ]]; # is number
&& # &
|| # or

if [[ -f $filename ]]; then
  echo "$filename is a regular file"
elif [[ -d $filename ]]; then
  echo "$filename is a directory"
elif [[ -p $filename ]]; then
  echo "$filename is a named pipe"
elif [[ -S $filename ]]; then
  echo "$filename is a named socket"
elif [[ -b $filename ]]; then
  echo "$filename is a block device"
elif [[ -c $filename ]]; then
  echo "$filename is a character device"
fi
if [[ -L $filename ]]; then
  echo "$filename is a symbolic link (to any file type)"
fi

if [[ -r $filename ]]; then
  echo "$filename is a readable file"
fi
if [[ -w $filename ]]; then
  echo "$filename is a writable file"
fi
if [[ -x $filename ]]; then
  echo "$filename is an executable file"
fi

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

if [[ -n "$string" ]]; then
  echo "$string is non-empty"
fi
if [[ -z "${string// }" ]]; then
  echo "$string is empty or contains only spaces"
fi
if [[ -z "$string" ]]; then
  echo "$string is empty"
fi

exit 0 # success
exit 1 # fail

# set Input Field Separator, by default ` ` (space)
IFS=:

[[ hello = h*o ]] && echo yes
[[ heeello =~ (e+) ]] && echo "yes, because: ${BASH_REMATCH[1]}"
[[ $1 ]] || { echo "missing argument" >&2; exit 1; }
{ cat x.txt || echo "file x.txt not found"; } 2>/dev/null

# default value
declare y=${myDefVar:-"nil"}
echo $y # nil
myDefVar=null
declare y=${myDefVar:-"nil"}
echo $y # null

# end of options:
touch -a # error
touch -- -a # ok

set -e # exit whenever a command fails
set -n # validate but not exec script
set -o #
set -u # error when using uniinitialized var
set -v #
set -v # print each command
set -x # to start debug

declare -i # interger
declare -r # readonly
declare -x # export

````

#### Debug:

````bash
#!/bin/bash -x

# or
set -x # to start debug
set +x # to end debug

# or
bash -x /Users/k/web/kovpak/gh/ed/bash/examples/hw.sh
````

#### Strings:

````bash
${#var} # string length

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
echo "${str/a/A}" # I Am a string
echo "${str//a/A}" # I Am A string
echo "${str/#I/=}" # = am a string
echo "${str/%g/N}" # I am a strinN
echo "${str/g/}" # I am a strin # replace with nothing
echo "${str%a*}"  # I am
echo "${str#*a}" # m a string
echo "${str##*a}" #  string

FILENAME="/tmp/example/myfile.txt"
echo "${FILENAME%/*}"    # /tmp/example
echo "${FILENAME##*/}"   # myfile.txt
BASENAME="${FILENAME##*/}"
echo "${BASENAME%%.txt}" # myfile

A=(hello world)
echo "${A[@]/#/R}" # Rhello Rworld
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
array=([3]='fourth element' [4]='fifth element')

arr=(1 2 3 a b)
echo ${#arr[@]} # array length
echo ${!arr[@]} # array indices

echo "${arr[@]: -1 }" # last el
echo "${array[-1]}" # last el
echo "${array[@]:1:3}"

array+=('fourth element' 'fifth element') # append
array=("new element" "${array[@]}") # preppend
arr=("${arr[@]:0:$i}" 'new' "${arr[@]:$i}") # insert

arr=(a b c)
echo "${arr[@]}" # outputs: a b c
echo "${!arr[@]}" # outputs: 0 1 2
unset -v 'arr[1]'
echo "${arr[@]}" # outputs: a c
echo "${!arr[@]}" # outputs: 0 2

unset array[10]

stringVar="Apple+Orange+Banana+Mango"
arrayVar=(${stringVar//+/ })
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

break 2 # Break multiple loop
````

#### X

````sh
fg # send command to foreground
bg # send command to background

# replace 1 to o
echo b1nd | tr 1 o # bond

# generates uuid
uuid -n 1

# tool to disable redundant background services on machine
sysv-rc-conf

# CPU usage info
# apt-get install sysstat
mpstat

# Memory usage info
free
cat /proc/meminfo

# Disc usage info
iotop

# tool for net traffic
cbm

# general stat
dstat

# shows OS limits
ulimit -a

# last argument
!$

id
who
echo $USER
hostname
tty

# directory mode
mkdir -m 777 test

which git
# /usr/local/bin/git

# Redirects:
#
# for standard input
echo yes 0> f
# for standard output
echo ok 1> f
# for errors
echo no 2>| f
# for standard output & errors
echo okay &> f
#
echo "test" 2>&1 1>/dev/null &
>&2 # output to stderr
1>&2 # output to stderr
2>&1 # stderr to stdout

# overwrite file
echo OK >| f

set -o
# to enable
set -o noclobber
# to disable
set +o noclobber

cut -f7 -d: /etc/passwd

# create named pipe
mkfifo mypipe

# when permission denied
echo '127.0.0.1 trick' | sudo tee -a /etc/hosts

stringInBase64=`echo "shaken not stirred" | base64`
echo $stringInBase64 | base64 -D

file ed/bash/README.md # prints file type

adduser rob
visudo
su rob
# on host machine
cat ~/.ssh/id_rsa.pub
# on remote machine
echo 'key from id_rsa.pub from host machine' >> ~/.ssh/authorized_keys
````

````sh
# nohup
# When you execute a job in the background (using &, bg command),
# and logout from the session - your process will get killed,
# to prevent this run:
nohup myscript &
# to run with low priority:
nice myscript
nohup nice myscript &

apt-get install -y --force-yes
# search package
apt-cache search htop

cat <<TXT
---------
204
---------
TXT

cat << EOF >> /tmp/yourfilehere
These contents will be written to the file.
        This line is indented.
EOF

ping 8.8.8.8 -c 15

# watch changes in directory
fswatch ./src | while read f; do echo $f; done

# copy data into clipboard buffer.
echo 200 | pbcopy

env # to see all ENV variables
sudo bash -c 'echo "APP_ENVIRONMENT=prod" > /etc/environment'
sudo sh -c 'echo "APP_ENVIRONMENT=prod" > /etc/environment'
sh -c 'echo 200'

cut -d' ' -f2 /tmp/file.txt # print column 2 from file using ' ' as delimer

df            # Show information about the file system.
df -h
df -T         # Show filesystem type
du            # Summarize disk usage of each FILE.
du -sh IraFoto/* # Summarize disk usage of each dir in foto dir.

uptime # CPU load average

uname -a # Shows 32 or 64 bit OS.
uname -r # Show the kernel version.
cat /etc/*release # all ablut ubuntu

nslookup git.mdm.comodo.od.ua
host github.com # Shows ip by host.
dig domain
whois ip

colordiff -u file1 file2

ln -s {file} {symbolic-name}

DISPLAY=:7
echo $DISPLAY

# print screen from url
cutycapt --url=cn007b.tumblr.com --out=out.png
wkhtmltoimage cn007b.tumblr.com i.jpg
webkit-image-gtk http://cn007b.tumblr.com > i.png

fwrite(STDOUT, __METHOD__ . "\n");
prompt \u@\h [\d]>

cd "$(dirname "$(readlink -f "$0")")"

echo There are ${#BASH_ALIASES[*]} aliases defined.

cmd || echo 'cmd failed'
docker info 2>/dev/null || echo 'fail'
test-d$HOME/.kube || mkdir$HOME/.kube

local val=${1:?Must provide an argument}

echo "${var:-XX}" # XX
var=23
echo "${var:-XX}" # 23

# process with pid
if [[ ! -e /tmp/test.py.pid ]]; then
  python test.py &
  echo $! > /tmp/test.py.pid
else
  echo -n "ERROR: The process is already running with pid "
  cat /tmp/test.py.pid
  echo
fi
````

````sh
cd -       # go to previous dir
pushd path # remember path (save it at stack)
popd       # got to pushed path (and delete it from stack)

# history:
# ⬆, ⬇ # keys to navigat through history
# Ctrl-p, Ctrl-n
# Ctrl-r # search in history
!207 # run from history

~/.bash_history
````

#### H

````sh
# scan ports:
nmap --script=http-headers www.ziipr.com

# osx proxy

sudo networksetup -setwebproxy "Wi-Fi" 54.174.16.166 80
# '71.4.186.100:21288', '198.7.96.243:21239', '104.128.17.224:21237', '209.48.175.196:21325', '172.102.207.55:21279', '198.7.97.209:21297', '162.248.134.54:21326', '23.239.219.244:21325', '143.191.19.7:21238', '173.44.12.201:21259', '173.44.5.52:21310', '173.44.12.68:21305', '173.44.26.80:21269', '107.152.144.66:21291', '208.73.73.206:21257', '204.14.87.85:21326', '144.168.128.88:21244', '204.14.87.149:21271', '45.57.195.33:21232', '173.44.5.185:21247', '173.44.12.141:21280', '173.44.26.220:21318', '107.152.144.219:21274', '208.73.73.9:21278', '143.191.19.89:21263', '143.191.31.123:21304', '69.174.99.149:21322', '50.117.15.33:21318', '173.44.12.16:21297', '216.172.144.220:21285',
sudo networksetup -setwebproxystate "Wi-Fi" off

netstat -an

nc -zv 10.0.2.2 22

nc -zv 78.140.186.238 1-65535
nc -zv 10.0.2.2 22

Public DNS IP addresses:
8.8.8.8
8.8.4.4
````
