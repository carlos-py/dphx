# sshKraken

[![Build Status](https://travis-ci.org/MOZGIII/dphx.svg?branch=master)](https://travis-ci.org/MOZGIII/dphx)

A SOCKS 5 server that provides a local tunnel to multiple SSH Server connections. Traffic is routed based on URL string matching specified in each SSH server configuration.

Designed to be managed as an unprivileged system service and run in the background. Will establish all connections at launch and will attempt to re-establish connection if necessary (ie. SSH connection is lost).

Remote DNS resolution is supported.

## Installation

To install via **Go** toolchain, use the following command:

```
go get github.com/carlos-py/sshKraken/cmd/sshKraken
```

## Configuration

SSH Server configuration is loaded from a JSON file. Each JSON object represents 1 server configuration.

- host: Server address and SSH port
- key: Private SSH key for connecting to server. **Must use full root path!**
- username: SSH username
- urlMatch: Socks proxy traffic containing this string will be routed to this server. The default route is the first server configured.

A sample config file below:

```bash
[
   {
      "host":"54.183.113.4:22",
      "key":"/home/toor/Downloads/thinkpad2018.pem",
      "username":"ubuntu",
      "urlMatch":"twitter.com"
   },
   {
      "host":"54.193.33.75:22",
      "key":"/home/toor/Downloads/thinkpad2018.pem",
      "username":"ubuntu",
      "urlMatch":"canihazip.com"
   }
]
```

