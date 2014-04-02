xdebug
-

	{
	    "path_mapping": {
	        "/var/www/vhosts/kovpak/htdocs/px/": "/home/kovpak/web/kovpak/px/",
	        "/usr/share/yii/": "/home/kovpak/web/kovpak/yii/framework/"
	    },
	    "close_on_stop": true
	}	

~/.bashrc
// home ip address: 007.02.03.04

	export XDEBUG_CONFIG="idekey=sublime.xdebug remote_host=007.02.03.04 remote_enable=1 remote_autostart=0"


.htaccess

	php_value xdebug.remote_enable 1
