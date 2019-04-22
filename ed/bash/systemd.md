Systemd
-

Systemd is a system management daemon.

````sh
systemctl status
systemctl list-units

systemctl enable lf@1323.service
systemctl daemon-reload
systemctl restart lf@1323.service
systemctl status lf@1323.service
systemctl list-units | grep lf

ls -la /etc/systemd/system/

systemd-analyze
````

#### Example

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
