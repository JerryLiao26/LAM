# LAM
LAM is Leave-A-Message, a web-based light-weight message center

## Build
### Prepare for it
To build LAM server, you first need ```go``` package installed on your OS. Then, make sure you have mysql driver for golang from [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql). And of course you should have ```mysql``` or ```mariadb``` ready. We currently only support **LOCAL** database.

### Build and use
Navigate to server folder and run command:
```
go build *
```
And you'll get executable file ```lam``` or ```lam.exe```(under Windows OS) or else. If you're using *nix systems, we suggest you copy the executable file to paths like ```/usr/bin/``` so you can run the command globally. After the copy, you can simply start a LAM server by:
```
lam start
```
For more help about the commands, you can run:
```
lam help
```

## Configure
working on this part...

## APIs
working on this part...
