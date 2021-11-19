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

2. The following config files should be located in the directory "free5GC/config".
```
samfcfg.yaml, ssmfcfg.yaml, and smncfg.yaml
```

3. The following network functions (NFs)/nodes folders should be located in the directory "free5GC/NFs".
```
samf, ssmf, ue, a-sat, g-sat, tgw, target a-sat, and target samf
```

4. The following NFs/nodes folders should be merged with the existing folders in the directory "free5GC/NFs".
```
amf, smf, udm, and upf
```

5. The following event trigger file should be located in the directory "free5GC".
```
start.go
```

 
## Running Simulation

1. Run NRF for registering NFs.
```
cd free5gc/NFs/nrf
go run nrf.go
```

2. Run required NFs/nodes for procedures.
```
cd free5gc/NFs/***
go run ***.go
```
(*** is the required NFs/nodes.)


- UE-triggered session establishment procedure requires at least the following NFs/nodes.
```
ue, a-sat, samf, ssmf, and g-sat
```


- Network-triggered session establishment procedure requires at least the following NFs/nodes.
```
ue, a-sat, samf, ssmf, g-sat, tgw, smf, udm, amd upf
```


- Intra-SC handover procedure requires at least the following NFs/nodes.
```
ue, a-sat, target a-sat, and samf
```


- UE-triggered session establishment procedure requires at least the following NFs/nodes.
```
ue, a-sat, target a-sat, target samf, g-sat, tgw, amf, and udm
```

3. Add the propagation delay in "smn_service/service.go" files in "a-sat", "target a-sat", and "tgw" folders.
You can use the propagation delay value from references. Or, you can measure the propagation delay from satellite network simulator 3 (SNS3).
```
https://github.com/sns3/sns3-satellite
```
The propagation delay can be added by using "time.Sleep()" in the following comments of "service.go" files.
```
	/*
		Add Propatation Delay
	*/
```

4. Add the customized code in "samf_service/service.go", and "ssmf_service/service.go" files in samf, and ssmf folders, respectively.
Also, add the customized code in "smn_service/serivce.go" files in "ue", "a-sat", "g-sat", "tgw", "target a-sat", "target samf", "udm", "upf", "smf", and "amf" folders.
You can add the customized code for specific control plane services. Or, you can measure the delay of the corresponding service from the existing Free5GC test code and add the delay using "time.Sleep()".
The customized code or delay should be added in the following comments of "service.go" files.
```
	/*
		Add *** Code
	*/
```

5. Run the event trigger file.
```
cd free5gc
go run start.go
```
Choose the control procedure for simulation and enter the number.

