####shell
<pre>
~/.bash_history

cd -       # go to previous dir
pushd path # remember path (save it at stack)
popd       # got to pushed path (and delete it from stack)
history
</pre>

####MC
<pre>
ctrl+u # change pannels palcements
alt+i  # open dir in second pannel
alt+s  # find match, again alt+s - find next match
</pre>

####grep
<pre>
-z, --null-data             строки разделяются байтом с нулевым значением, а не символом конца строки
-v, --revert-match          выбирать не подходящие строки
-m, --max-count=ЧИСЛО       остановиться после указанного ЧИСЛА совпадений
-h, --no-filename           не начинать вывод с имени файла
-L, --files-without-match   печатать только имена ФАЙЛОВ без совпадений
</pre>

`grep pattern -Pz` # несколько строк
`grep "$DC"`

####curl
<pre>
-b, --cookie STRING/FILE  String or file to read cookies from (H)
-d, --data DATA     HTTP POST data (H)
-F, --form CONTENT  Specify HTTP multipart POST data (H)
-o, --output FILE   Write output to <file> instead of stdout
-x, --proxy [PROTOCOL://]HOST[:PORT] Use proxy on given port
-e, --referer       Referer URL (H)
-u, --user USER[:PASSWORD]  Server user and password
</pre>