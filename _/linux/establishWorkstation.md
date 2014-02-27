Установка LAMP (аббревиатура от Linux Apache MySQL PHP) на Ubuntu
Внимание: способ писался под Ubuntu 11.10, для других версий возможны отличия.
Внимание: для корректной работы необходит PHP 5.3, рекомендуется использовать oldstable репозитории для установки.

Мы можем быстро поставить LAMP с помощью tasksel:

sudo apt-get install tasksel
sudo tasksel install lamp-server
Но tasksel поставит много лишнего, при этом нужные нам вещи все равно придется доставлять вручную, потому мы будем использовать попакетную установку, плюс это позволит нам понимать что мы делаем. Итак, начнем.

Установка apache2:

sudo apt-get install apache2
Апач нужно сразу перезапустить:

sudo service apache2 restart
Проверяем работоспособность апача: по адресу http://localhost в браузере мы должны увидеть сообщение It works!

Установка PHP5

Ставим PHP, активируем его и перезапускаем апач:

sudo apt-get install libapache2-mod-php5
sudo a2enmod php5
sudo service apache2 restart
Установка MySQL:

sudo apt-get install mysql-server libapache2-mod-auth-mysql php5-mysql
Установщик MySQL попросит нас придумать пароль для пользователя root в MySQL.

Мы уже имеем работоспособный сервер и вы теперь продвинутый юзверь, теперь приступим к настройке нашего сервера.

Edit
Настройка сервера
Переходим в домашнюю папку и создаем там любой каталог, например web (в домашней папке работать удобно и не заденешь лишнего).

Копируем файл настройки по умолчанию в новый файл:

sudo cp /etc/apache2/sites-available/default /etc/apache2/sites-available/mysite
Редактируем его:

sudo gedit /etc/apache2/sites-available/mysite
Изменяем корневой узел документов
с /var/www/ на

/home/<user>/web/
Также изменяем <Directory /var/www/> на

<Directory /home/<user>/web/>
И сохраняем файл.

Актвируем mysite, деактивируем default и перезапускаем апач:

sudo a2dissite default && sudo a2ensite mysite
sudo service apache2 restart
Теперь на http://localhost мы должны видеть содержимое /home/<user>/web/.

Edit
Создание доменов для работы над несколькими проектами
Внимание: для нашего проекта этот шаг является не обязательным, так как настроены авто-поддомены на <вашеимя>.alcuda.priv и работать с проектом вы будете на поддоменах. Настраивайте автоподдомены, если вам нужны также другие домены, иначе этот раздел можно пропустить.

Создаем vhosts.conf:

sudo gedit /etc/apache2/sites-available/vhosts.conf
Со следующим содержимым:

<VirtualHost *:80>
ServerName test1.my
ServerAlias www.test1.my
DocumentRoot /home/<user>/web/test1.my
</VirtualHost>
<VirtualHost *:80>
ServerName test2.my
ServerAlias www.test2.my
DocumentRoot /home/<user>/web/test2.my
</VirtualHost>
Сохраняем документ, открываем в свою папку web и там создаем папки test1.my и test2.my
Далее открываем в текстовом редакторе файл /etc/hosts:

sudo gedit /etc/hosts
И добавляем в конец файла:

127.0.0.1 test1.my
127.0.0.1 test2.my
Сохраняем файл и активируем наш vhosts.conf, затем перезапускаем апач:

cd /etc/apache2/sites-enabled/
sudo ln -s . ./sites-available/vhosts.conf (уберите пробел между точек, какой-то баг в редмайне)
sudo service apache2 restart
Теперь на http://test1.my и http://test2.my должно отображаться содержимое соответствующих папок.

Edit
Настройка и установка необходимых для проекта компонентов
Включаем отображение ошибок в PHP:

sudo gedit /etc/php5/apache2/php.ini

Ставим display_errors и html_errors в On.
Активируем Mod Rewrite:

sudo a2enmod rewrite
Также меняем в нашем файле mysite AllowOverride None на:

AllowOverride All
Установка Imagemagick:

sudo apt-get install imagemagick php5-imagick
Установка Memcached:

sudo apt-get install memcached php5-memcache
Установка Smarty:

sudo apt-get install smarty
Edit
Установка и настройка SSL
Установка сертификата:

sudo apt-get install ssl-cert
sudo mkdir /etc/apache2/ssl
sudo make-ssl-cert /usr/share/ssl-cert/ssleay.cnf /etc/apache2/ssl/apache.pem
Активируем модуль ssl:

sudo a2enmod ssl
sudo service apache2 force-reload
Копируем ssl-виртуалхост по умолчанию в новый файл и настраиваем его:

sudo cp /etc/apache2/sites-available/default /etc/apache2/sites-available/ssl
sudo gedit -w /etc/apache2/sites-available/ssl
Записываем в файл такое содержимое (не забывайте менять на свои пути):

ameVirtualHost *:443
<virtualhost *:443>
ServerAdmin webmaster@localhost

SSLEngine On
SSLCertificateFile /etc/apache2/ssl/apache.pem

DocumentRoot /home/<user>/web/
<directory />
Options FollowSymLinks
AllowOverride All
</directory>

<directory /home/<user>/web/>
Options Indexes FollowSymLinks MultiViews
AllowOverride All
Order allow,deny
allow from all
# This directive allows us to have apache2's default start page
# in /apache2-default/, but still have / go to the right place
# Commented out for Ubuntu
#RedirectMatch ^/$ /apache2-default/
</directory>

ScriptAlias /cgi-bin/ /usr/lib/cgi-bin/
<directory "/usr/lib/cgi-bin">
AllowOverride None
Options ExecCGI -MultiViews +SymLinksIfOwnerMatch
Order allow,deny
Allow from all
</directory>

ErrorLog /var/log/apache2/error.log

# Possible values include: debug, info, notice, warn, error, crit,
# alert, emerg.
LogLevel warn

CustomLog /var/log/apache2/access.log combined
ServerSignature On

Alias /doc/ "/usr/share/doc/"
<directory "/usr/share/doc/">
Options Indexes MultiViews FollowSymLinks
AllowOverride None
Order deny,allow
Deny from all
Allow from 127.0.0.0/255.0.0.0 ::1/128
</directory>

</virtualhost>
Теперь активируем наш виртуалхост и перезапускаем апач:

sudo a2ensite ssl
sudo service apache2 restart
Готово!
Затем вам нужно установить и настроить git (подробности в соотвествующей статье).

Edit
Примечания
Так как проекту нежны некоторые каталоги вне докрута, то, удобнее держать проект в папке, вложенной в папку web/.
Допустим вы сделали git clone из папки web без указания в какую папку делать clone, тогда будет создана папка multi_web/ и проект будет слит в нее.
Теперь поменяйте в виртуалхостах mysite и ssl /home/<user>/web на /home/<user>/web/multi_web/ и перезапустите апач.
При этом необходимые проекту каталоги такие как private и local удобно будет держать в /home/<user>/web/ (рядом с папкой multi_web).

Если вы настраивали дополнительные домены в vhosts.conf (как вы помните, это необязательный шаг), то в vhosts.conf пути менять не нужно.
В итоге при такой файловой структуре:

web/
    multi_web/
    local/
    private/
    test1.my/
    test2.my/

Домены test1.my и test2.my будут работать как и положено, а по адресу http://localhost или <вашеимя>.alcuda.priv будет открываться папка multi_web/.
P.S. Данная статья находится на этапе тестирования, могут быть какие-то упущения, статья будет дополняться и исправляться.