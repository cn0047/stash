xdebug
-

	{
	    "path_mapping": {
	        "/var/www/vhosts/surname/htdocs/px": "/home/surname/web/kovpak/px/"
	    },
	    "close_on_stop": true
	}

~/.bashrc
// home ip address: 007.02.03.04

	export XDEBUG_CONFIG="idekey=sublime.xdebug remote_host=007.02.03.04 remote_enable=1 remote_autostart=0"


.htaccess

	php_value xdebug.remote_enable 1
