# 操作记录

###  docker build . -t httpserver:0.0.1
###  docker run -d httpserver:0.0.1
### nsenter -t $PID -n ip a
```
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
22: eth0@if23: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether 02:42:c0:a8:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 192.168.0.2/20 brd 192.168.15.255 scope global eth0
       valid_lft forever preferred_lft forever
```

### docker  tag httpserver:0.0.1  wangchaoyang/cloudnative:httpserver-0.0.1
### docker  push wangchaoyang/cloudnative:httpserver-0.0.1
