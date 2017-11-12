# Dolly.go

This is a simple tool to deploy your application by a git repository, written in golang. It works like a webserver that you have a endpoint that will be your hook.

## Getting Started

First you need to have golang installed in your pc to compile the bin. You also will need ssh keys from your github/bitbucket repository configured, and a server with a ssh connection.


### Installing

You can install it by the two ways:

from the git:
```
# git clone https://github.com/jorgeluisrezende/Dolly.git
```

or by the go package mannager: 

```
# go get -v https://github.com/jorgeluisrezende/Dolly
```
After you download the source code, you need to navigate to project folder, like:

```
# cd "YOURGOPATH"/src/github.com/jorgeluisrezende/Dolly/
```
and build the bin:

```
# go build *.go
```

So, after tha you will get a bin that you will run in you repo folder at the your server.

##

## Authors

* **Jorge Luis** - *Initial work* - [jorgeluisrezende](https://github.com/jorgeluisrezende)

## License

This project is licensed under the MIT License 
