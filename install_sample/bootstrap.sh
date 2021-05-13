#!/bin/bash

echo "Executing bootstrap script..."

HOME=/app speedtest -f json --accept-license --accept-gdpr

ls -lha /etc/systemd/system/speedtest.service

cat /etc/systemd/system/speedtest.service



#cat <<EOF > /etc/systemd/system/speedtest.service
##########################################
## /etc/systemd/system/speedtest.service #
##########################################
#
#[Unit]
#Description=Logs system statistics to the systemd journal
#Wants=SpeedtestMonitor.timer
#
#[Service]
#Type=oneshot
#ExecStart=/app/speedtest-wrapper-go -p --config ./config.yml
#
#[Install]
#WantedBy=multi-user.target
#EOF

