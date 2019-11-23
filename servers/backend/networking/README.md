sudo sysctl net.ipv4.ip_local_port_range="1500 61000"
sudo sysctl net.ipv4.tcp_fin_timeout=10
look into txqueuelen for ethernet card speeds

servers: 
sysctl -w net.core.somaxconn=10000
sysctl net.core.netdev_max_backlog=30000
sysctl net.ipv4.tcp_max_syn_backlog=30000
