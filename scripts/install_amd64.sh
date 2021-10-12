#!/bin/bash

if ! command -v jq& /dev/null
then
    echo "jq could not be found"
    exit
fi

wget "$( curl https://api.github.com/repos/eldius/speedtest-wrapper-go/releases | jq -r '. | sort_by(.created_at) | last | .assets[] | select(.name | endswith(".amd64")) | .browser_download_url' )"

mv speedtest-wrapper* speedtest-wrapper
sudo mv speedtest-wrapper /usr/bin
