# gotftpd

Go TFTP server

## TFTP Background

This is a very simple implementation of a TFTP server.
TFTP is a very simple protocol to transfer files.
It is from this that the name comes, Trivial File Transfer Protocol or TFTP.

TFTP is defined in [RFC 1350](https://tools.ietf.org/html/rfc1350)

## Usage

```
  -addr string
    	Address to listen (default "0.0.0.0:69")
  -path string
    	Local file path (default ".")
```

`-addr` is the address/port that the system will listen on, 0.0.0.0 means the system will listen to any incoming connections.

`-path` is the directory where the files that need to be served are currently stored.

### Example

We want to server a firmware image for a router to update from. The file is called `newfirmware.bin` and we are storing it in our `firmware` folder. Our computer and user has the permissions to run on privileged ports (<1024).

Change directory to where the file is stored.

```
cd firmware
```

List the files to make sure the file is stored.

```
ls -l

total 1640
-rw-r--r-- 1 peterp peterp 1678336 Sep 26 09:38 newfirmware.bin
```

Start running the gotftpd server.

```
gotfptd
```

A client (router) connects and downloads the file.

```
1678336 bytes sent
```

When a client downloads the file we can that the transfer happened.

## Common messages

You may see messages like.

```
read udp [::]:47539: i/o timeout
read udp [::]:32901: i/o timeout
```

This is common and is just the connection from a previous active session closing.

