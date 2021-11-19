# SMN-Simulator

This code is for satellite mobile network (SMN) simulation and the code is implemented by Free5GC.


## Environment
The simulations have been conducted under the following environment
```
OS: Ubuntu 20.04
gcc: 9.3.0
Go: 1.14.4 linux/amd64
kernel version 5.11.0-40-generic
```


## Setting

1. Install the Free5GC.
```
https://github.com/free5gc/free5gc
```

2. The following config files should be located in directory "free5GC/config".
```
samfcfg.yaml, ssmfcfg.yaml, and smncfg.yaml
```

3. The following network functions (NFs)/nodes folders should be located in directory "free5GC/NFs".
```
samf, ssmf, ue, a-sat, g-sat, tgw, target a-sat, and target samf
```

4. The following NFs/nodes folders should be merged with the existing folders in directory "free5GC/NFs".
```
amf, smf, udm, and upf
```

 





