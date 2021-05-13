# install samples #

```
#########################################
# /etc/systemd/system/speedtest.service #
#########################################

[Unit]
Description=Logs system statistics to the systemd journal
Wants=speedtest.timer

[Service]
Type=oneshot
ExecStart=/app/speedtest-wrapper-go -p --config ./config.yml

[Install]
WantedBy=multi-user.target
```

```
#########################################
# /etc/systemd/system/speedtest.timer   #
#########################################
[Unit]
Description=Logs some system statistics to the systemd journal
Requires=myMonitor.service

[Timer]
Unit=speedtest.service
OnCalendar=*-*-* *:*:00

[Install]
WantedBy=timers.target
```

## references ##

- [Use systemd timers instead of cronjobs](https://opensource.com/article/20/7/systemd-timers)
- [How To Use Vagrant With Libvirt KVM Provider](https://ostechnix.com/how-to-use-vagrant-with-libvirt-kvm-provider/)
