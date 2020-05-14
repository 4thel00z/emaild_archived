# emaild

## Description

emaild is a cool email daemon 😎 which can schedule email sending to a later time in the day and also handle bulk requests.
Communication happens through a unix domain socket, which means two things:
- this is a unix software
- this software is intended for local usage

## Installation

```
make
# in case you want emaild to be installed under /usr/local/bin
make install
```

## Configuration

The configuration of the daemon can be done via the accounts.json file.
It currently stores your SMTP accounts in plain json.
This is subject to change. 😇

## Todos

- Add makefile (lol) 😇
- Provide unix domain socket webserver
- Document the interface
- Provide a client
