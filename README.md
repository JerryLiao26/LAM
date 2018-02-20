# LAM
LAM is Leave-A-Message, a web-based light-weight message center

## Built Release
[v0.1](https://github.com/JerryLiao26/LAM/releases/tag/v0.1)

## Build
### Prepare for it
To build LAM server, you first need ```go``` package installed on your OS. Then, make sure you have MySQL driver for golang from [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql). And of course you should have ```MySQL``` or ```MariaDB``` ready. We currently only support *LOCAL* database.

### Build and use
Navigate to server folder and run command:
```
go build *
```
And you'll get executable file ```lam``` or ```lam.exe```(under Windows OS) or else. If you're using *nix systems, we suggest you copy the executable file to paths like ```/usr/bin/``` so you can run the command globally. After the copy, you can simply start a LAM server by:
```
lam start 0.0.0.0:12580
```
For more help about the commands, you can run:
```
lam help
```

## Configure
### Server
At your first start, you shall go through a configure prompt, then everything should be settled. If you want to re-configure database settings, use:
```
lam set LAM:lamLAM
```
to set database username to LAM and password to lamLAM.

### Database
To use LAM server, you need a database named ```LAM```, and tables as follows:
- message
  - id(INT) [PRIMARY, AUTO_INCREMENT]
  - tag(VARCHAR)
  - content(TEXT)
  - timestamp(TIMESTAMP)
  - ifRead(TINYINT)

- token
  - tag(VARCHAR) [PRIMARY]
  - token(VARCHAR)
  - timestamp(TIMESTAMP)

There's an sql file under ```sql/``` for you to import

## APIs
### /hello
- Method: Any
- Require: Nothing
- Respond:
```json
{
  "code": 200,
  "text": "LAM Server here",
  "method": "Method name"
}
```

### /send
- Method: POST
- Require:
```json
{
  "text": "Your message here",
  "token": "Your token here"
}
```
- Respond:
```json
{
  "code": 200,
  "text": "Message from Your tag saved",
  "method": "POST"
}
```

## Related Links
- [LAM-client-Web](https://github.com/JerryLiao26/LAM-client-Web)
