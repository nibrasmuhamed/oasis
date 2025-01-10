# Distributed Cluster Management POC

## Introduction
There are many commercial and open source distributed in-memory databases. 
- Memcached / Hazelcast
- https://github.com/mailgun/groupcache
- https://github.com/buraksezer/olric/ 
- Redis
- Oracle coherence
- Apache Ignite

Caches are either embedded in  application cluster or hosted externally in a cluster. 

Embedded caches provide low latency as they 
* avoid a network hop 
* low operation overhead with no extra management/monitoring of cache servers
* easier solution 

External Cache cluster help in 
* scaling horizontally easily 
* are independent of application crashes&restarts  
* don't compete with application memory reducing cache eviction
* cache coherence problems are less here.

Cluster is managed either in a decentralized elected coordinator within the cluster (gossip protocol) or with an external centralized service. This service has to do the following
* Membership management: Detect Cluster Membership changes (ie) Node addition, removal (by listening to kubernetes API service)
* Protocol definition for communication.  
* Message Notification 
  * Notifying listeners (clients) about changes in node membership
  * Notify Nodes about the partition that needs to be evicted or hosted in its memory

## POC 
This POC demonstrates a simulated centralized coordinator to manage the partition distribution in the cluster and also inform the clients on cluster membership changes. It uses consistent hashing with load balancing logic to evenly distribute the cache keys among the cluster. 

Check the architecture diagram of the proposed solution. 
https://app.diagrams.net/#G1vQE0qwHJ_yBUqhTlDCCFkRbzQQzf1dhl

### Coordinator:

Coordinator service saves the latest information of memberlist in the file "members.json" and also configures the hash replication and key partition count.

Coordinator runs on 8081 port.
Coordinator starts with 8 nodes and every 5 seconds alternatively adds or removes a node

```
cd coordinator
go run testCoordinator.go
```

### Client:

Clients run on 900* series and connects to coordinator (:8081). The last digit of client port can be mentioned in the command line
```
cd testClient
go run test.go -port 1
```

### Get Key:

Get the customer key's node location from the client, check the sample command below
```
curl http://127.0.0.1:9001/customer/1232
```

### TODO

* Handling consistently the hash collision of node names
* Key replication in secondary nodes in case primary fails, also handling sync up of these data
* Handling cold start issues of thundering db hits