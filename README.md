# Hellogo
This program is playground in golang with LEGO EV3. This project is based on great project [ev3dev.org][5]. This porgram use golang library [github.com/ev3go/ev3dev][2]. Program `print` could print simple drawing with [jkbrickworks.com/telegraph-machine-and-printer/][1]. 

## How to build programs and deploy them to EV3
 
Programs are directory names from `./cmd/`. Following examples use program `find`.

Directory `scripts` contains all scripts necessary to make executable files and scripts for uploading executable files to EV3 brick. Programs could be compiled and deployed to EV3 brick with folloving command:

```
./scripts/makeAndDeploy.sh find
```

or it could be done step by step:

```
./scripts/make.sh find
./scripts/deploy.sh find
```

Note that in ./scripts/conf.sh if hard coded IP of EV3 brick. Check that's correct value.

## How to execute programs at EV3

First find IP address of EV3 brick. Look at brickman user interface. At the top there is assigned IP addess. For all work use user `robot` with password `maker`. Root have same password. For login just type:

```
ssh robot@192.168.1.46
```

Than execute program. After loging you are at robot you are at home directory `/home/robot`. There should be previously prepared programs. Execute them in a standard way `./find`.

## Usefull commands

### Login without password

EV3 could be accessed without password. Just import local ssh key to brick ss keystore. Public part of local key could be uploaded with following command:

```
ssh-copy-id -i ~/.ssh/id_ed25519.pub robot@192.168.1.46
```

### Upload files manually to EV3 brick

For upload program to `/home/robot` directory:

```
scp hellogo robot@192.168.1.46:/home/robot
```

### Switch off Brickman

When ev3dev starts it launch brickman user interface. It's nice but it controll screen and catch button events. To stop brickan user interface use:
```
sudo systemctl stop brickman
```

### Switch off Brickman permanently

To stop starting brickman after loading execute following:
```
sudo systemctl disable brickman
``` Command to let brickman start after booting again is:
```
sudo systemctl enable brickman
```.

### Switch off EV3

To switch off EV3 brick execute in EV3 shell:
```
sudo shutdown -h now
```.
Immediatelly after sending this command EV3 starts glowing red to signalize that is switching off.


## Unsorted mess

Install library for native working with USD peripherials, golang:

```
sudo apt update
sudo apt full-upgrade -y
sudo apt install -y libusb-1.0-0-dev rsync joe mc htop links dnsmasq
```

```
figlet -t -k EV3
```


### Golang
Compile specific file locally:

```
go build -o ./bin/usb ./cmd/usb/usb.go
```

Clear all locally downloaded modules:

```
go clean -modcache
```

Allows to compile hellogo with locally adjusted module:

```
go mod edit -replace github.com/gotmc/libusb/v2=/Users/honza/Documents/projects/github/libusb/
```

Get all available versions and get specific version localy:
```
go list -m -versions github.com/gotmc/libusb
go get github.com/gotmc/libusb@v1.0.21
```

### rsync

Because it's not possible to compile C code for controlling USB at local machine. It has to be uploaded to EV3 bricke and compiled there:

```
rsync -r ./hellogo/ robot@192.168.2.2:/home/robot/hellogo
rsync -r ./libusb/ robot@192.168.2.2:/home/robot/libusb

```

Because of compatibility problems with usblib version, you'll need to see package version:

```
dpkg -l | grep libusb
```

## EV3 Connecting
Connected EV3 brick is listed as:

```
root@ev3dev:/home/robot# lsusb
Bus 001 Device 003: ID 0694:0005 Lego Group Mindstorms EV3
Bus 001 Device 001: ID 1d6b:0001 Linux Foundation 1.1 root hub
```
When second EV3 brick is connect via USB to first one in a few minutes will be accessible.

## Usefull links

* Which Wi-Fi dongles are compatible with EV3 brick is described at [www.ev3dev.org/docs/networking/][3].
* Usefull guide [how to connect EV3 brick via cable to computer][4	]. After that EV3 brick could be accessed by ssh.

[1]: https://jkbrickworks.com/telegraph-machine-and-printer/ "printer building instruction"
[2]: https://github.com/ev3go/ev3dev "https://github.com/ev3go/ev3dev"
[3]: https://www.ev3dev.org/docs/networking/ "https://www.ev3dev.org/docs/networking/"
[4]: https://www.ev3dev.org/docs/tutorials/connecting-to-the-internet-via-usb/ "https://www.ev3dev.org/docs/tutorials/connecting-to-the-internet-via-usb/"
[5]: https://www.ev3dev.org "https://www.ev3dev.org"