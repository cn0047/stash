Linux
-

screen
which

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

####shell
````
df    # Show information about the file system.
df -h
du    # Summarize disk usage of each FILE.

service supervisor restart
supervisorctl restart
supervisorctl status

service memcached restart


uptime # CPU load average

pkill /var/www/x.php # Kill runned script x.php.
pkill -f php # Kill runned script with name like php.

uname -r # Show the kernel version.

host github.com # Shows ip by host.
dig domain
whois ip

colordiff -u file1 file2

scp -rp access@host:~/dir/ ~/dir/

ssh-add ~/.ssh/id_rsa

~/.bash_history

cat
tail
````
````
cd -       # go to previous dir
pushd path # remember path (save it at stack)
popd       # got to pushed path (and delete it from stack)
history
````
````
git st -s
 M cron/.htaccess
?? ppc/Google/AdWords/AdWordsApiLibrary/src/Google/Api/Ads/AdWords/v201402/
# Output second colunm. ' ' at `awk -F ' '` - separator
git st -s | awk -F ' ' '{print $2}'
cron/.htaccess
ppc/Google/AdWords/AdWordsApiLibrary/src/Google/Api/Ads/AdWords/v201402/
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
PATH=/bin:/usr/bin:/usr/local/bin
CRUN_REMOTE_HOST=Host
CRUN_EMAIL=mail@com.com
CRUN_WORK_DIR=/var/www/vhosts/host/htdocs

rm /tmp/crun_user_at_server_or_host_3a30db060f74d9390a2eb6f8a92eab8d # crun lock file. Should be removed when execution fails...

grep cron /var/log/syslog
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
