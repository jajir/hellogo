# Hellogo
This program is playground in golang with LEGO EV3. It's based on library [github.com/ev3go/ev3dev][2]. Program `print` could print simple drawing with [jkbrickworks.com/telegraph-machine-and-printer/][1]. 

## How to build and execute programs
 
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

## Usefull commands

### Log to brick

First find IP address. Look at brickman user interface. At the top there is assigned IP addess. For all work use user `robot` with password `maker`. Root have same password. For login just type:

```
ssh robot@192.168.1.46
```
### Login without password

EV3 could be accessed without password. Just import local ssh key to brick ss keystore. Public part of local key could be uploaded with following command:

```
ssh-copy-id -i ~/.ssh/id_ed25519.pub robot@192.168.1.46
```

### Upload program to brick

For upload program to `/home/robot` directory:

```
scp hellogo robot@192.168.1.46:/home/robot
```

### Switch off Brickman

When ev3dev starts it launch brickman user interface. It's nice but it controll scrren and catch buttons. To stop brickan user interface use:
```
sudo systemctl stop brickman
```

[1]: https://jkbrickworks.com/telegraph-machine-and-printer/ "printer building instruction"
[2]: https://github.com/ev3go/ev3dev "https://github.com/ev3go/ev3dev"