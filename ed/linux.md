Linux
-

####image optimization

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
````
# install jpegoptim on centos 6
sudo yum install libjpeg*
cd /tmp
wget http://www.kokkonen.net/tjko/src/jpegoptim-1.4.1.tar.gz
tar -xzf jpeg*
cd ./jpeg*
./configure
make
make install
sudo ln -s /usr/local/bin/jpegoptim /usr/bin/

# install pngquant on centos 6
sudo yum install libpng*
wget http://pngquant.org/pngquant-2.4.0-src.tar.bz2
bunzip2 pngquant-2.4.0-src.tar.bz2
tar -xvf pngquant-2.4.0-src.tar
./configure
make
make install
sudo ln -s /usr/local/bin/pngquant /usr/bin/
````

####rsync

````
rsync -az --progress --exclude=.git/* --rsh='ssh -p26' kovpak@trunk-web-php.pmmedia.priv:/usr/share/yii/ /home/volodymyr/web/kovpak/yii/framework/
rsync -az --progress --rsh='ssh -i /home/kovpak/web/storage/ziipr.pem' ./ ec2-user@ec2-52-210-246-232.eu-west-1.compute.amazonaws.com:/var/www/
````

####sed

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

````
sed -i "s/admin_user/user/" /var/www/html/config.php

####awk

AWK was created in the 1970s.

````
echo 'one and two' | awk '{print $1}' # will print one
awk 'BEGIN {print "Hello, world!"}'
ps aux|awk 'length($0) > 150' # Print lines longer than 150 characters
printf "one\n* two\n" | awk '{print ($1=="*" ? $2 : $1)}' # Print one \n two
printf "1\n 2\n 3\n" | awk 'ORS=NR?",":"\n"' # Replace new line with comma
````

####shell
````
df    # Show information about the file system.
df -h
du    # Summarize disk usage of each FILE.

service memcached restart


uptime # CPU load average

pkill /var/www/x.php # Kill runned script x.php.
pkill -f php # Kill runned script with name like php.

kill -9 `ps -aux|grep {{PROCESS_NAME}}|grep -v grep|awk '{print $2}'`

uname -r # Show the kernel version.

host github.com # Shows ip by host.
dig domain
whois ip

colordiff -u file1 file2

scp -rp access@host:~/dir/ ~/dir/

~/.bash_history

ln -s {file} {symbolic-name}

cat
tail

screen
which

DISPLAY=:7
echo $DISPLAY

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

####ssh
````
ssh-add ~/.ssh/id_rsa
vim /etc/ssh/sshd_config

sshfs -o nonempty -p22 root@host:/home/host/www /home/user/web/www
fusermount -u /home/user/web/www
````

####ps
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

####chmod
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
file      - -rw-rw-r-- (664)
````

####mail
````
mail -s 'subject' mail@com.com < fileName
uuencode card.jpg card.jpg | mail mail@com.com
````

####grep
````
-z, --null-data             строки разделяются байтом с нулевым значением, а не символом конца строки
-v, --revert-match          выбирать не подходящие строки
-m, --max-count=ЧИСЛО       остановиться после указанного ЧИСЛА совпадений
-h, --no-filename           не начинать вывод с имени файла
-L, --files-without-match   печатать только имена ФАЙЛОВ без совпадений
````

`grep pattern -Pz` # несколько строк

`grep "$DC"`

`grep match -wrni --color=always --include=*.php . | grep notMatch -v --color=always`

####crun
````
cat /etc/cron.d/
````
````
SHELL=/bin/bash
PATH=/bin:/usr/bin:/usr/local/bin
CRUN_REMOTE_HOST=Host
CRUN_EMAIL=mail@com.com
CRUN_WORK_DIR=/var/www/vhosts/host/htdocs

0 * * * * /bin/echo `date` >> /tmp/d.tmp

rm /tmp/crun_user_at_server_or_host_3a30db060f74d9390a2eb6f8a92eab8d # crun lock file. Should be removed when execution fails...

grep cron /var/log/syslog
````
````
# restart on centos
sudo service crond restart
````

####curl
````
-b, --cookie STRING/FILE             String or file to read cookies from (H)
-d, --data DATA                      HTTP POST data (H)
-F, --form CONTENT                   Specify HTTP multipart POST data (H)
-o, --output FILE                    Write output to <file> instead of stdout
-x, --proxy [PROTOCOL://]HOST[:PORT] Use proxy on given port
-e, --referer                        Referer URL (H)
-u, --user USER[:PASSWORD]           Server user and password

# (REST) JSON at POST.
curl -X POST -H 'application/json' -d '{"key":"val"}' http://url.com
curl -X POST -H 'Content-Type: text/plain' -d @/tmp/foo.txt http://url.com
curl http://url.com -d x=1 -d y=2

curl http://login:pass@base-http-auth.com/
curl -u login:pass http://base-http-auth.com/

# upload file
curl http://localhost:8000 -F "file=@/home/kovpak/Downloads/download.jpg"
````

####SNIPPETS
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
````
