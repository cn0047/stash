xdebug
-

    {
    "path_mapping": {
        "/var/www/vhosts/kovpak/htdocs/px/": "/home/kovpak/web/kovpak/px/",
        "/usr/share/yii/": "/home/kovpak/web/kovpak/yii/framework/"
    },
    // "break_on_start": true,
    "debug_layout" : {
        "cols": [0.0, 1.0, 1.0],
        "rows": [0.0, 0.7, 1.0],
        "cells": [[0, 0, 2, 1], [0, 1, 1, 2], [0, 1, 1, 2]]
    },
    "breakpoint_group": 2,
    "breakpoint_index": 2,
    "context_group": 1,
    "context_index": 2,
    "stack_group": 2,
    "stack_index": 2,
    "watch_group": 2,
    "watch_index": 2,
    "close_on_stop": true
    }

~/.bashrc
// home ip address: 007.02.03.04

    export XDEBUG_CONFIG="idekey=sublime.xdebug remote_host=007.02.03.04 remote_enable=1 remote_autostart=0"
    source ~/.bashrc


.htaccess

    php_value xdebug.remote_enable 1

    php_value xdebug.remote_enable 1
    php_value xdebug.remote_host localhost
    php_value xdebug.remote_connect_back 1
    php_value xdebug.remote_port 9000
    php_value xdebug.remote_handler dbgp
