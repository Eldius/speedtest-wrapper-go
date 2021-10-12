# speedtest-wrapper-go #

It's a tool to execute the speedtest network test and save it in
a database and/or post it to a MQTT broker.

## features/commands ##

- `speedtest-wrapper-go`
  - `test`: actually it just execute speedtes test and display the result


## snippets ##

### mongodb raspberry ###

```bash
echo 'deb http://ftp.br.debian.org/debian stretch main' /etc/apt/sources.list.d/repo_mongodb_org_debian.list
sudo apt-get update && sudo apt-get install mongodb-server
```

### install ###

```bash
# raspberry
bash <(curl -s -L https://raw.githubusercontent.com/Eldius/speedtest-wrapper-go/main/scripts/install_raspiberry.sh) --argument1=true
# amd64
bash <(curl -s -L https://raw.githubusercontent.com/Eldius/speedtest-wrapper-go/main/scripts/install_amd64.sh) --argument1=true
```
