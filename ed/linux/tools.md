Linux Tools
-

#### apachebench

````
ab -k -n 5000 -c 100 -t 2 "http://localhost"
# Where:
# -n   Number of requests.
# -c   Concurrency.
# -t   Timelimit in seconds.
````

#### less

````
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
-F ':' # column separator

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

scp -rp access@host:~/dir/ ~/dir/
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

#### watch

execute a program periodically, showing output fullscreen.

````
-n, --interval <secs> # seconds to wait between updates
````

watch -n 2 'echo 200'
watch -n 2 'date'

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

# shows spent time (⏱)
time curl -si https://realtimelog.herokuapp.com/test
````
