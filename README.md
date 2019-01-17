# xlab

A reusable and dynamic container environment for binary analysis / general CTF shenanigans.
I intend to build a robust cli around the container to help with dependency management / general ease of use.  

### BUILD

```
docker pull n0ja/xlab:latest
```

### COMPILE GO  

If you want to run the go cli

```
go get github.com/n0ja/xlab
cd {GODIR}/src/github.com/n0ja/xlab && go install
xlab&
```    
Depending on your GOBIN / os setup you may need to make an alias to point to the output binary  
Options are -t "terminal emulator", -n "container name", -v "image version"  

### DISCLAIMER

This is an early build forked from EpicTreasure and the Dockerfile will be modified heavily.  
As such, I garuntee no compatibility for current toolsets / environments until a stable release.
The removal of this disclaimer will signify stability so please refer back here when updating as I will not be versioning the image on docker hub until I have solidified my plans.
