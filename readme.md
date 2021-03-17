## GCP ENUM

This project was inspired by the [GCPBucketBrute](https://github.com/RhinoSecurityLabs/GCPBucketBrute).

GCP enum is the tool that helps to find the existence of the GCP buckets. This might be helpful during the pentesting or recon.

> This tools only checks for the existence of the bucket ,it doesn't check like permission on the existing bucket If you want to check the permission I highly recommend using GCPBucketBrute

### Installation

If you want to  build manually Clone the repo and install it using **go install**

or You can download the premade-binary and put it in your PATH

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


### Why GCP enum

- Its  faster than **GCPBucketBrute** 
- It has pre-made binary ,So no need to install the tools and dependencies