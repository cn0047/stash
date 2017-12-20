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
cat /proc/sys/net/core/somaxconn
sysctl -n net.core.somaxconn
# on server must be at least 1024
````

#### about linux

````
cat /etc/os-release
````

#### upstart

Is an event-based replacement for the traditional init daemon.

#### centos

````
yum --showduplicates list available elasticsearch\*
````

#### netstat

````
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

````
screen -S sessionName # new named session
screen -ls            # list of active sessions
screen -x sessionName # attach to session
````
Ctrl+A followed by D = detach screen

#### image optimization

````
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

````
rsync -az --progress --exclude=.git/* --rsh='ssh -p26' kovpak@trunk-web-php.pmmedia.priv:/usr/share/yii/ /home/volodymyr/web/kovpak/yii/framework/
rsync -az --progress --rsh='ssh -i /home/kovpak/web/storage/ziipr.pem' ./ ec2-user@ec2-52-210-246-232.eu-west-1.compute.amazonaws.com:/var/www/
````

#### sed

Sed was developed from 1973.

````
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
````

#### awk

AWK was created in the 1970s.

````
echo 'one and two' | awk '{print $1}' # will print one
awk 'BEGIN {print "Hello, world!"}'
ps aux|awk 'length($0) > 150' # Print lines longer than 150 characters
printf "one\n* two\n" | awk '{print ($1=="*" ? $2 : $1)}' # Print one \n two
printf "1\n 2\n 3\n" | awk 'ORS=NR?",":"\n"' # Replace new line with comma
````

#### ssh
````
ssh-add ~/.ssh/id_rsa
vim /etc/ssh/sshd_config

sshfs -o nonempty -p22 root@host:/home/host/www /home/user/web/www
fusermount -u /home/user/web/www
ps aux | grep -i sftp | grep -v grep

sudo ssh-add ~/.ssh/id_rsa

ssh user@server.com

ssh -i key -N -L 9229:127.0.0.1:9229 root@server
# -L local_socket:remote_socket
````

`vim ~/.ssh/config`
````
Host ec2
    Hostname ec2-52-211-26-56.eu-west-1.compute.amazonaws.com
    User ec2-user
    IdentifyFile ~/path_to_ssh_key
````

#### ftp
````
ncftpput -R -v -u {user} {host} remote_path ./local_path/*
````

#### ps
````
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

````
adduser rob
visudo
su rob
# on host machine
cat ~/.ssh/id_rsa.pub
# on remote machine
echo 'key from id_rsa.pub from host machine' >> ~/.ssh/authorized_keys
````

#### chmod
````
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

#### mail
````
mail -s 'subject' mail@com.com < fileName
uuencode card.jpg card.jpg | mail mail@com.com
````

#### grep
````
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
````
-b, --cookie STRING/FILE             String or file to read cookies from (H)
-d, --data DATA                      HTTP POST data (H)
-F, --form CONTENT                   Specify HTTP multipart POST data (H)
-o, --output FILE                    Write output to <file> instead of stdout
-x, --proxy [PROTOCOL://]HOST[:PORT] Use proxy on given port
-e, --referer                        Referer URL (H)
-u, --user USER[:PASSWORD]           Server user and password

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
````

#### shell
````
# watch changes in directory
fswatch ./src | while read f; do echo $f; done

# prettify json output
echo '{"foo": "f", "bar": "b", "items": [1, 2, 3] }' | jq
echo '{"foo": "f", "bar": "b", "items": [1, 2, 3] }' | jq '.items | length'

# copy data into clipboard buffer.
echo 200 | pbcopy

# md5 for directory
find src/ -type f -exec md5sum '{}' \; | md5sum

sudo bash -c 'echo "APP_ENVIRONMENT=prod" > /etc/environment'
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

host github.com # Shows ip by host.
dig domain
whois ip

colordiff -u file1 file2

scp -rp access@host:~/dir/ ~/dir/

~/.bash_history

ln -s {file} {symbolic-name}

which

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

# date, in next format [02/жов/2016:00:21
date +'[%d/%b'/%Y:%H:%M
````

#### H
````sh
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
