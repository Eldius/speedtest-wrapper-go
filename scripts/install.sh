#!/bin/bash

if ! command -v jq& /dev/null
then
    echo "jq could not be found"
    exit
fi


release_data=$( curl https://api.github.com/repos/eldius/speedtest-wrapper-go/releases | jq '. | map({name: .name, url: .assets[0].browser_download_url, created: .created_at, asset_name: .assets[0].name}) | sort_by(.created) | last' )



curl https://api.github.com/repos/eldius/speedtest-wrapper-go/releases | jq '. | map({name: .name, url: .assets[0].browser_download_url, created: .created_at, asset_name: .assets[0].name}) | .[] | select(.asset_name | endswith("raspberry")) | sort_by(.created) | last'
