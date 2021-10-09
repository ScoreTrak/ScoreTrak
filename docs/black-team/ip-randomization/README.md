# Rotating source IPs of the scoring engine

It might be useful sometimes to rotate the IP addresses of the scoring engine, in order to prevent competitors from easily whitelisting the scoring infrastructure.

Following is one way to introduce source IP randomization.

1) Create a simple linux machine, that will be responsible for randomizing the addresses using iptables.
2) Assign multiple IP addresses to the machine(they don't have to be on different interfaces). This will depend on your specific network manager, so for now let's assume it is netplan:

```
    ens192:
     addresses: [10.0.4.25/20, 10.0.4.30/20, 10.0.4.10/20]
     gateway4: 10.0.0.1
    ens160:
     addresses: [10.2.4.5/20]
``` 

Let the .5 be "gateway" for the workers, and the rest of the IPs will serve as random source

3) Enable IP forwarding and load iptables modules: https://www.karlrupp.net/en/computer/nat_tutorial
```
echo 1 > /proc/sys/net/ipv4/ip_forward
modprobe ip_tables
modprobe ip_conntrack
modprobe ip_conntrack_irc
modprobe ip_conntrack_ftp
```
4) Allow forwarding:
```
iptables -I FORWARD -i ens160 -o ens192 -j ACCEPT
```
5) Allow established connections:
```
iptables -I FORWARD -i ens192 -o ens160 -m state --state ESTABLISHED,RELATED -j ACCEPT
```
6) Configure SNAT:
```
iptables -t nat -I POSTROUTING -d 0.0.0.0/0 -m state --state NEW -m statistic --mode nth --every 3 --packet 0 -j SNAT --to-source 10.0.4.10
iptables -t nat -I POSTROUTING -d 0.0.0.0/0 -m state --state NEW -m statistic --mode nth --every 2 --packet 0 -j SNAT --to-source 10.0.4.25
iptables -t nat -I POSTROUTING -d 0.0.0.0/0 -m state --state NEW -m statistic --mode nth --every 1 --packet 0 -j SNAT --to-source 10.0.4.30
```

NOTE that `--every` is decreasing with every rule. Following explains the behavior: https://serverfault.com/questions/490854/rotating-outgoing-ips-using-iptables/491517#491517

The above configuration will use one of the IPs as source NAT in a round-robin.
7) Ensure competition workers are routed to the specified IP address


# Caveats

Certain protocols like FTP might not play well with IP randomization.
In case IP randomization would not work for certain protocols, it might be best to send the packets directly to the team's gateway from the worker node(instead of sending them to IP randomizing gateway)