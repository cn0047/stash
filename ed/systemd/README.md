Systemd
-

Place bin file in `/home/k/myapp`
run `vim /lib/systemd/system/myapp.service`.

````
[Unit]
Description=myapp

[Service]
Type=simple
Restart=always
RestartSec=5s
ExecStart=/home/k/myapp

[Install]
WantedBy=multi-user.target
````

and `service myapp start; service myapp status`
enable it on boot `service myapp enable`.
