Linux Tools
-

#### kill

````sh
pkill /var/www/x.php # Kill runned script x.php.
pkill -f php # Kill runned script with name like php.
pkill -f test.py # full argument lists, default is to match against process names

kill -9 `ps -aux|grep {{PROCESS_NAME}}|grep -v grep|awk '{print $2}'`
````sh

#### monitoring

````sh
dmesg | tail      # last 10 system messages
vmstat 1          # virtual memory stat
mpstat -P ALL 1   # CPU time breakdowns per CPU
pidstat 1         # top’s per-process summary
iostat -xz 1      # block devices (disks)
free -m           # block device I/O buffer cache & file systems page cache
sar -n DEV 1      # check network interface
sar -n TCP,ETCP 1 # TCP metrics
top               #
````

#### date

````sh
# date, in certain format:
date +'%d/%b/%Y:%H:%M' # 18/Jul/2018:11:48
date +%Y-%m-%dT%T%z    # 2018-07-18T11:49:03+0300
date +%s               # timestamp
date +%s%N             # seconds and current nanoseconds
````

#### ls

````sh
ls -A # except for . and ..
ls -h # sizes
ls -s # number of file system blocks
````

#### find

````sh
# delete all -name directories
find . -type d -name node_modules -exec rm -rf {} \;
find . -type d -name vendor -exec rm -rf {} \;

# md5 for directory
find src/ -type f -exec md5sum '{}' \; | md5sum

find -type f -mtime -20 | while read file; do modif=`git log -1 --format="%cd" $file`; echo "$modif - $file"; done
# Shows file types that present in foolder
find . -type f | perl -ne 'print $1 if m/\.([^.\/]+)$/' | sort -u

time find -name '*.php' -mtime -1 | xargs -l php -l | grep -v 'No syntax errors detected in'
find ./ | grep '.php' | xargs -l php -l | pv | grep -v 'No syntax errors detected in'
# skype Vdovin
cat /tmp/alcuda_tech.log | grep -Eo --color=never '^\[\S+\s+[^\*][^:]+' | sed -r 's/^\S+\s+//g' | sort | uniq -c | sort -nr | head -30
find -type f -name '*.php' -exec egrep -l 'class\s+ProfileManager' {} \;
find -name '*.htm' -exec touch {} \;
find -type f -name '*.php' -exec egrep -Hn --color=always 'is_failed' {} \; | grep profile
````

#### sort

````
sort -nk2 # sort as numbers by column 2
````

#### read

Reads user input (prompt).

````
# reads and prints input
read; echo $REPLY
read -r; echo $REPLY

# reads 1 character
read -r -n 1; echo $REPLY
````

#### printf

````sh
printf "one\ntwo\n"
printf "one\n \ttwo\n \tthree\n"
printf %s\\n {0..9}

printf "Hi %s \n" $USER
printf "p%s\n" 1 2 3
printf -v myVar "UserIs%s" $USER

printf "%s\n" *.{gif,jpg,png} # display only (some) image files

v="hello"
printf '%s\n' "${v^}" # Hello
printf '%s\n' "${v^^}" # HELLO

v="hello world"
declare -u string="$v" # HELLO WORLD

v="BYE"
printf '%s\n' "${v,}" # bYE
printf '%s\n' "${v,,}" # bye

v="HELLO WORLD"
declare -l string="$v" # hello world

v="Hello World"
echo "${v~~}" # hELLO wORLD
echo "${v~}" # hello World

var='0123456789abcdef'
printf '%s\n' "${var:3}" # 3456789abcdef
printf '%s\n' "${var:3:4}" # 3456
printf '%s\n' "${var:3:-5}" # 3456789a
printf '%s\n' "${var: -6}" # abcdef
printf '%s\n' "${var:(-6)}" # abcdef
printf '%s\n' "${var: -6:-5}" # a

set -- 0123456789abcdef
printf '%s\n' "${1:5}" # 56789abcdef

myarr[0]='0123456789abcdef'
printf '%s\n' "${myarr[0]:7:3}" # 789

set -- 1 2 3 4 5 6 7 8 9 0 a b c d e f
printf '%s\n' "${@:10}"
printf '%s\n' "${@:10:3}"
printf '%s\n' "${@:10:-2}"
printf '%s\n' "${@: -10:2}"
printf '%s\n' "${@:0:2}"

myarr=(0 1 2 3 4 5 6 7 8 9 a b c d e f)
printf '%s\n' "${myarr[@]:12}" c
printf '%s\n' "${myarr[@]:5:3}"
printf '%s\n' "${myarr[@]: -1}"
````

#### echo

````sh
echo 'Hello world'
echo "Hello world $USER"
echo -n "Hello world $USER" # no new line
echo n-{a,b,c}-b # fn-a-b n-b-b n-c-b
echo n#{1..3}, # n#1, n#2, n#3,

var='12345'
echo "${#var}" # length

myarr=(1 2 3)
echo "${#myarr[@]}" # number of array elements

v=foo-bar-baz
echo ${v%%-*} # foo
echo ${v%-*} # foo-bar
echo ${v##*-} # baz
echo ${v#*-} # bar-baz

echo "FOO is ${FOO:?EMPTY}"
echo "FOO is ${FOO?UNSET}"
echo "BAR is ${BAR:?EMPTY}"
echo "BAR is ${BAR?UNSET}"

echo {a..z}
echo {a..d}{1..3}
````

#### apachebench

````
ab -k -n 5000 -c 100 -t 2 "http://localhost"
# Where:
# -n   Number of requests.
# -c   Concurrency.
# -t   Timelimit in seconds.
````

#### less

````sh
# OPTIONS
-F or --quit-if-one-screen # exit if the entire file can be displayed on the first screen.
-i or --ignore-case
-M or --LONG-PROMPT
-m or --long-prompt        # prompt verbosely (like more).
-R or --RAW-CONTROL-CHARS  #
-S or --chop-long-lines    # chop lines longer than the screen width.
-X or --no-init            #
-xn,... or --tabs=n,...    # sets tab stops.
````

#### upstart

Is an event-based replacement for the traditional init daemon.

#### netstat

````sh
netstat -anp | grep LISTEN

# osx
sudo lsof -i -n -P | grep TCP.*80

# Active Internet connections
netstat -tpne
````

#### mount

Run `mount` command to see types of mounted filesystems.

On server machine:
````
echo '/tmp 52.48.208.239(rw,sync,no_subtree_check,no_root_squash)' >> /etc/exports
````

On client machine:
````
mount -t nfs {remote_machine_ip}:/remote/dir /local/dir
````

#### screen

````sh
screen -S sessionName # new named session
screen -ls            # list of active sessions
screen -x sessionName # attach to session
````
Ctrl+A followed by D = detach screen

#### image optimization

````sh
jpegoptim --size=12k --dest a/ -o origin.jpg
jpegoptim --max=30 --dest=res -o  bond.origin.jpg

convert origin.jpg -quality 95 result.jpg
convert origin.jpg -resize 100 result.jpg
convert origin.jpg -resize 50% result.jpg
convert bond.origin.jpg -resize 200x150 bond.res.jpg

jpegtran -copy none -optimize -outfile res.jpg origin.jpg

pngquant --force --quality=25-25 ~/Downloads/origin.png --ext .25.png

# resize image to particular size.
mogrify -resize 400x400! src.png
````

#### rsync

````sh
rsync -az --progress /source /target
rsync -az --progress --exclude=.git/* --rsh='ssh -p26' kovpak@trunk-web-php.pmmedia.priv:/usr/share/yii/ /home/volodymyr/web/kovpak/yii/framework/
rsync -az --progress --rsh='ssh -i /home/kovpak/web/storage/ziipr.pem' ./ ec2-user@ec2-52-210-246-232.eu-west-1.compute.amazonaws.com:/var/www/
````

#### sed

Sed was developed from 1973.

````sh
# Mode of operation
d - delete
q - quit
N - add the next line

echo 'Some text or regexp' | sed 's/regexp/replacement/g' # Some text or replacement
echo 'cat and dog' | sed -r "s/(cat|dog)s?/\1s/g" # cats and dogs
echo 'xxx and zzz' | sed 's/x/y/g' # yyy and zzz
echo '1 one; 2 two' | sed 's/1//g; s/2//g' #  one;  two

sed -i "s/admin_user/user/" /var/www/html/config.php

# delete 2nd line from file
sed -e '2d;' file.txt > res.file.txt

echo A_B_C_D_E_F_20180904_0038849_SUPER03.txt \
  |sed -e 's/_[0-9]\{8\}_[0-9]\{7\}_[A-Z0-9]*//g'
````

#### awk

AWK was created in the 1970s.

````sh
-F ':' # column separator

echo 'one and two' | awk '{print $1}' # will print one
awk 'BEGIN {print "Hello, world!"}'
ps aux|awk 'length($0) > 150' # Print lines longer than 150 characters
printf "one\n* two\n" | awk '{print ($1=="*" ? $2 : $1)}' # Print one \n two
printf "1\n 2\n 3\n" | awk 'ORS=NR?",":"\n"' # Replace new line with comma
````

#### ssh

`pssh` tool to run ssh command on few servers simultaneously

````sh
locate sshd_config

ssh-add ~/.ssh/id_rsa
vim /etc/ssh/sshd_config

sshfs -o nonempty -p22 root@host:/home/host/www /home/user/web/www
fusermount -u /home/user/web/www
ps aux | grep -i sftp | grep -v grep

sudo ssh-add ~/.ssh/id_rsa

ssh user@server.com

ssh -i key -N -L 9229:127.0.0.1:9229 root@server
# -L local_socket:remote_socket

scp -rp user@host:~/dir/ ~/dir/
````

`vim ~/.ssh/config`
````
Host ec2
    Hostname ec2-52-211-26-56.eu-west-1.compute.amazonaws.com
    User ec2-user
    IdentifyFile ~/path_to_ssh_key
````

#### ftp

````sh
ftp $hostname

# connect in passive mode
ftp -p $hostname

ncftpput -R -v -u {user} {host} remote_path ./local_path/*
````

````sh
status
system          # show remote system type
ascii           # set ascii transfer type
binary          # set Binary transfer type
dir             # list contents of remote directory
ls              # list contents of remote directory

size filename
get filename    # receive file
recv filename   # receive file
delete filename

bye             # exit
close
````

#### watch

execute a program periodically, showing output fullscreen.

````
-n, --interval <secs> # seconds to wait between updates
````

watch -n 1 'echo $(date +%s%N)'
watch -n .01 'echo $(date +%s%N)'
watch -n 2 'date'

#### ps

````sh
ps ax                             # Display all processess
ps aux                            # Detailed information about processess
ps -f -u www-data                 # Display process by user "www-data"
ps -C apache2 u                   # Show process by name
ps -f -p 1400                     # Show process by id
ps aux --sort=-pcpu,+pmem         # Sort process by cpu or memory usage ("-" or "+" symbol indicating descending or ascending sort)
ps -f --forest -C apache2         # Display process hierarchy in a tree style
ps -o pid,uname,comm -C apache2   # Display child processes of a parent process
ps -p 3150 -L                     # Display threads of a process
ps -e -o pid,uname,pcpu,pmem,comm # Change the columns to display
ps -e -o pid,uname=USERNAME,pcpu=CPU_USAGE,pmem,comm
ps -e -o pid,comm,etime           # Display elapsed time of processes

watch -n 1 'ps -e -o pid,uname,cmd,pmem,pcpu --sort=-pmem,-pcpu | head -15' # Turn ps into an realtime process viewer
````

#### chmod

````sh
-rwxrwxrwx
drwxr-xr-x
directory user group other

r (4) - read;
w (2) - write;
x (1) - execute;

Example:
u   g   o
7   5   3
rwx r-x -wx

chmod u-rx,g+x,o+w fileName

On server:
directory - drwxrwxr-x (775)
socket    - srw-rw---- (660)
file      - -rw-rw-r-- (664)
````
````
chown -R www-data:www-data /var/www/html/
````

````
sudo chown -R `whoami` /var/log/
````

#### crun (cron)

````
# dir with cron jobs
ls /etc/cron.d
````

#### mail

````
mail -s 'subject' mail@com.com < fileName
uuencode card.jpg card.jpg | mail mail@com.com
````

#### grep

````sh
-z, --null-data             строки разделяются байтом с нулевым значением, а не символом конца строки
-v, --revert-match          #
-m, --max-count=ЧИСЛО       #
-h, --no-filename           #
-L, --files-without-match   # print only file names
````

`grep pattern -Pz` # несколько строк

`grep "$DC"`

`grep match -wrni --color=always --include=*.php . | grep notMatch -v --color=always`

#### curl

````sh
-b, --cookie STRING/FILE             String or file to read cookies from (H)
-d, --data DATA                      HTTP POST data (H)
-F, --form CONTENT                   Specify HTTP multipart POST data (H)
-o, --output FILE                    Write output to <file> instead of stdout
-O, --remote-name                    Write output to file wiht origin name
-x, --proxy [PROTOCOL://]HOST[:PORT] Use proxy on given port
-e, --referer                        Referer URL (H)
-u, --user USER[:PASSWORD]           Server user and password
-v, --verbose

# (REST) JSON at POST.
# More examples available here: https://github.com/cn007b/my/blob/master/ed/php.symfony/generateRESTapi.md
curl -XPOST http://localhost:3000/dishes \
    -H 'Content-Type: application/json' -d '{"name": "newDish", "description": "newDesc"}'
curl -X POST -H 'application/json' -d '{"key":"val"}' http://url.com
curl -X POST -H 'Content-Type: text/plain' -d @/tmp/foo.txt http://url.com
curl http://url.com -d x=1 -d y=2

curl http://login:pass@base-http-auth.com/
curl -u login:pass http://base-http-auth.com/

// user == 'admin' && pass == 'password'
curl http://localhost:3000 -H 'Authorization: Basic YWRtaW46cGFzc3dvcmQ='

# upload file
curl http://localhost:8000 -F "file=@/home/kovpak/Downloads/download.jpg"
curl http://localhost:8000 -H "Content-Type: multipart/form-data" -F "file=@/Users/k/f.txt" -F "msg=MyFile"

curl -T firmware.bin http://0.0.0.48/cgi-bin/package_install?hash=017

# shows spent time (⏱)
time curl -si https://realtimelog.herokuapp.com/test
````
