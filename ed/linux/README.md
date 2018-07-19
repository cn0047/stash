Linux
-

````
# nginx
telnet localhost 8080

# php-fpm
telnet localhost 9000

# websocket
telnet 0.0.0.0 12345
````

````
# HTTPS keys:
openssl genrsa 1024 > private.key
openssl req -new -key private.key -out cert.csr
openssl x509 -req -in cert.csr -signkey private.key -out certificate.pem
````

sudo dpkg -i {name}

On Linux systems, a kernel configuration parameter called `net.core.somaxconn`
provides an upper limit on the value of the backlog parameter passed to the listen function
that is used to create the servers listening socket.
If the backlog argument is greater than the value in /proc/sys/net/core/somaxconn,
then it is silently truncated to that value.
The default value in this file is 128.

````
/dev/null       # stream
/dev/stderr     # stream
/dev/stdin      # stream
/dev/stdout     # stream
/etc            # system configuration directory
/etc/os-release # about linux
/etc/ttys       # logged in users
/opt            # optional software directory
````

````
cat /proc/sys/net/core/somaxconn
sysctl -n net.core.somaxconn
# on server must be at least 1024
````

````
cat /etc/passwd
sar:x:205:105:Stephen Rago:/home/sar:/bin/ksh

colon-separated fields:
sar          - the login name,
x            - encrypted password,
205          - numeric user ID
105          - numeric group ID
Stephen Rago - a comment field
/home/sar    - home directory
/bin/ksh     - and shell program
````

Each directory contains subdirectories `.` and `..`.
<br>`dot` refers to the current directory,
<br>`dot-dot` refers to the parent directory,

````
cc  - the C compiler,
gcc - the GNU C compiler,
````

````
adduser rob
visudo
su rob
# on host machine
cat ~/.ssh/id_rsa.pub
# on remote machine
echo 'key from id_rsa.pub from host machine' >> ~/.ssh/authorized_keys
````

#### shell

````
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

# for standard output
echo ok 1> f
# for errors
echo no 2>| f
# for standard output & errors
echo kk &> f

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
````

````
# nohup
# When you execute a job in the background (using &, bg command),
# and logout from the session - your process will get killed. 

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

# md5 for directory
find src/ -type f -exec md5sum '{}' \; | md5sum

env # to see all ENV variables
sudo bash -c 'echo "APP_ENVIRONMENT=prod" > /etc/environment'
sudo sh -c 'echo "APP_ENVIRONMENT=prod" > /etc/environment'
sh -c 'echo 200'

cut -d' ' -f2 /tmp/file.txt # print column 2 from file using ' ' as delimer

df            # Show information about the file system.
df -h
du            # Summarize disk usage of each FILE.
du -sh IraFoto/* # Summarize disk usage of each dir in foto dir.

uptime # CPU load average

pkill /var/www/x.php # Kill runned script x.php.
pkill -f php # Kill runned script with name like php.

kill -9 `ps -aux|grep {{PROCESS_NAME}}|grep -v grep|awk '{print $2}'`

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
````

````
cd -       # go to previous dir
pushd path # remember path (save it at stack)
popd       # got to pushed path (and delete it from stack)

history
!207 # run from history

~/.bash_history
````

#### SNIPPETS

````
fwrite(STDOUT, __METHOD__ . "\n");
prompt \u@\h [\d]>
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

# date, in certain format:
date +'%d/%b/%Y:%H:%M' # 18/Jul/2018:11:48
date +%Y-%m-%dT%T%z    # 2018-07-18T11:49:03+0300
````

#### H

````sh
# scan ports:
nmap --script=http-headers www.ziipr.com

# osx proxy

sudo networksetup -setwebproxy "Wi-Fi" 54.174.16.166 80
sudo networksetup -setwebproxystate "Wi-Fi" off

netstat -an

nc -zv 10.0.2.2 22

nc -zv 78.140.186.238 1-65535
nc -zv 10.0.2.2 22

Public DNS IP addresses:
8.8.8.8
8.8.4.4
````
