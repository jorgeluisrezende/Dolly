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

##Deploy

Here we need a hash to be your endpoint, so you just need run the bin built in prev step:

```
"YOURGOPATH"/src/github.com/jorgeluisrezende/Dolly:# ./dolly -gh "whateveryouwanthere"
```
It will return to you a md5 hash, you don't need expecificaly a hash but it is a simple way to keep your and point "safe", you just need copy to use in the next step:

```
"YOURGOPATH"/src/github.com/jorgeluisrezende/Dolly:# ./dolly -s -p "a port that you want" -hash "you hash got at the prev step" &
```
With that you will get up your hook server in the background, if you dont want to put a port the server will run at the 8001 port. Your endpoint will be something like this:

```
localhost:8001/"yourhash" or localhost:"yourport"/"yourhash"
```
Now you just need open it in you nginx, apache or whatever server you want and do the next step that is configurate your hook on bitbucket or github.

## Authors

* **Jorge Luis** - *Initial work* - [jorgeluisrezende](https://github.com/jorgeluisrezende)

## License

This project is licensed under the MIT License 
