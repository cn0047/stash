xdebug
-

	{
	    "path_mapping": {
	        "/var/www/vhosts/kovpak/htdocs/px": "/home/kovpak/web/kovpak/px/"
	    },
	    "close_on_stop": true
	}

~/.bashrc

	export XDEBUG_CONFIG="idekey=sublime.xdebug remote_host=192.168.12.184 remote_enable=1 remote_autostart=0"

home 176.36.24.43

	export XDEBUG_CONFIG="idekey=sublime.xdebug remote_host=176.36.24.43 remote_enable=1 remote_autostart=0"

.htaccess

	php_value xdebug.remote_enable 1