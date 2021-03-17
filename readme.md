## GCP ENUM

This project was inspired by the [GCPBucketBrute](https://github.com/RhinoSecurityLabs/GCPBucketBrute).

GCP enum is the tool that helps to find the existence of the GCP buckets. This might be helpful during the pentesting or recon.

> This tools only checks for the existence of the bucket ,it doesn't check like permission on the existing bucket If you want to check the permission I highly recommend using GCPBucketBrute

### Installation

- Install the Golang then :

```
go get github.com/JOSHUAJEBARAJ/gcp-enum
```

### Usage

```
$ gcp-enum
  -c int
        Default concurrency value is 5 you can change the value using the c flag (default 5)
  -file string
        File name containing the word list
  -k string
        keyword that you want to enumerate

```

For example to scan for the word netflix

```
$ gcp-enum -k netflix -file wordlist
```


By default, this tool will create the 3 concurrent processes To increase the concurrency use the **-c flag**

```
$ gcp-enum -k netflix -file wordlist -c 10
```

The above command will create 10 concurrent process
### Why GCP enum

- Its somewhat faster than **GCPBucketBrute** 
- Because I want to learn GO