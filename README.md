# mqtt-chat

MQTT chat example with go and Google Container Engine


## Composition

```
                     +------------------+
                     |  +------------+  |
                     |  | mqtt-chat1 |  |
+----------+         |  |   (POD)    |  |
| client A +-------> |  +------------+  |
+----------+         |                  |  MQTT
             XHR /   |  +------------+  |        +-------------+
             SSE     |  | mqtt-chat2 |  +------> | mqtt-master |
                     |  |   (POD)    |  | <------+    (POD)    |
                     |  +------------+  |        +-------------+
                     |                  |
+----------+         |  +------------+  |
| Client B +-------> |  | mqtt-chat3 |  |
+----------+         |  |   (POD)    |  |
                     |  +------------+  |
                     |                  |
                     |    mqtt-chat     |
                     |       (RC)       |
                     +------------------+
```

* mqtt-chat
  - go chat server using XHR or Server Sent Events with the client, and MQTT with mqtt-master
  - replicated 3 pods by Kubernates replication controller
* mqtt-master
  - single pod of Mosquitto server

## Deployment

Setup your Container Engine Cluster. Then type following kubectl to deploy the app.

```bash
$ kubectl create -f mosquitto-controller.yaml
$ kubectl create -f mosquitto-service.yaml
$ kubectl create -f chat-controller.yaml
$ kubectl create -f chat-service.yaml
```

## License

See [LICENSE](./LICENSE) file for detail.

