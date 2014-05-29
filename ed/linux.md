####shell
````
colordiff -u file1 file2

scp -rp access@host:~/dir/ ~/dir/

ssh-add ~/.ssh/id_rsa

~/.bash_history
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

####crun
````
PATH=/bin:/usr/bin:/usr/local/bin
CRUN_REMOTE_HOST=Host
CRUN_EMAIL=mail@com.com
CRUN_WORK_DIR=/var/www/vhosts/host/htdocs
````

####curl
````
-b, --cookie STRING/FILE  String or file to read cookies from (H)
-d, --data DATA     HTTP POST data (H)
-F, --form CONTENT  Specify HTTP multipart POST data (H)
-o, --output FILE   Write output to <file> instead of stdout
-x, --proxy [PROTOCOL://]HOST[:PORT] Use proxy on given port
-e, --referer       Referer URL (H)
-u, --user USER[:PASSWORD]  Server user and password

# show ip
curl -Iq http://2ip.ru | grep IP | awk '{print $2}'
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
