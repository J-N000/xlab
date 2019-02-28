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
```    
Depending on your GOBIN / os setup you may need to make an alias to point to the output binary  
Options are -t "terminal emulator", -n "container name", -v "image version", -c "commit"

### USAGE

Please use & when mounting to run in detach!  
  
All default params (t="urxvt", n="xlab", v="latest", c="")      
```
xlab &
```   
starts the default xlab container  
  
-t   
```
xlab -t xterm &
```  
opens xlab in an xterm window  
  
-n
```
xlab -n test &
```  
starts an xlab container named test  
    
-c
```
xlab -c test -v test
```  
commits the xlab container named test to tag n0ja/xlab:test  
  
-v
```
xlab -v test
```  
starts an xlab container based on the image tagged n0ja/xlab:test  

---

### EXAMPLE
```
#host 
cd ~/workingDir
touch example.txt
xlab -n example &
```  
Opens a tty in an xlab container named example  
  
```
#example
pwd -> /root/mount
ls -> example.txt
apt-get update
```  
Update w/ apt-get
  
```
#host
xlab -c example
```  
Commits the xlab container named example to n0ja/xlab:latest  
  
### DISCLAIMER

This is an early build forked from EpicTreasure and the Dockerfile will be modified heavily.  
As such, I garuntee no compatibility for current toolsets / environments until a stable release.
The removal of this disclaimer will signify stability so please refer back here when updating as I will not be versioning the image on docker hub until I have solidified my plans.
