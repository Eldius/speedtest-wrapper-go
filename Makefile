
.EXPORT_ALL_VARIABLES:

SPEEDTEST_MQTT_HOST="192.168.100.195"
SPEEDTEST_MQTT_PORT="1883"
SPEEDTEST_MQTT_USER=""
SPEEDTEST_MQTT_PASS=""
SPEEDTEST_MQTT_CLIENTNAME="Speed_Test"
SPEEDTEST_MQTT_TOPIC="/speedtest_wrapper/tv_room"

clean:
	-rm -rf bin

bin/speedtest-wrapper-go:
	go build -v -o bin/speedtest-wrapper-go -a -ldflags '-extldflags "-static"'  -ldflags "-X 'github.com/Eldius/speedtest-wrapper-go/config.BuildDate=$(date +"%Y-%m-%dT%H:%M:%S%:z")' -X 'github.com/Eldius/speedtest-wrapper-go/config.Version=$(git rev-parse --short HEAD)'" .

build: bin/speedtest-wrapper-go
	@echo "Build with success"

start:
	go run main.go test -p --config config/samples/config.yml

vagrantclean:
	-cd install_sample ; vagrant destroy -f
	-rm -rf install_sample/.tmp

vagrantssh:
	-cd install_sample ; vagrant ssh

vagrantsetup: vagrantclean build
	-mkdir -p install_sample/.tmp
	cp bin/speedtest-wrapper-go install_sample/.tmp/speedtest-wrapper-go
	cp config/samples/config.yml install_sample/.tmp/config.yml
	cp install_sample/scripts/*.sh install_sample/.tmp/
	cd install_sample ; vagrant up

vagrantlibvirt: vagrantclean build
	-mkdir -p install_sample/.tmp
	cp bin/speedtest-wrapper-go install_sample/.tmp/speedtest-wrapper-go
	cp config/samples/config.yml install_sample/.tmp/config.yml
	cp install_sample/scripts/*.sh install_sample/.tmp/
	cd install_sample ; vagrant up --provider=libvirt

molecule:
	cd install_sample/ansible/roles/setup_speedtest_wrapper ; pwd ; molecule test
