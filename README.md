# Introduction
This is a monorepo containing gin webserver and future e-commerce web site.

## What is the purpose
This project is made for my wife and it serves as a portfolio to reveal her art works as well as e-commerce platform for selling her digital art, print works.

You can view the web site by following [bogrik.com](https://bogrik.com/) link. I hope by the time you read it, it is up and working. 

## How to start it locally
You must have Golang installed. In the project root type `go mod tidy` to install all the dependies, then, `go run .` should run the gin server. Open `localhost:8085` to see the page. 

## How it is deployed
* Build the app `go build`
* Setup a systemd service
```
[Unit]
Description=Static Website for Bogrik
After=network.target

[Service]
User={username}
Group={usergroup}
ExecStart={path to the go binary}
Restart=on-failure
RestartSec=5s

[Install]
WantedBy=multi-user.target
```
* Setup a reverse proxy (I am using Caddy) 
```
{your-domain-name} {
	reverse_proxy localhost:8085
}
```
* `systemctl enable {service-name}` `systemctl start {service-name}` and `caddy reload -f {path-to-Caddyfile}`
