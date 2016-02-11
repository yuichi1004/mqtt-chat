# mqtt-chat

MQTT chat example with go and Google Container Engine

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

## Deployment

Setup your Container Engine Cluster. Then type following kubectl to deploy the app.

```bash
$ kubectl create -f mosquitto-controller.yaml
$ kubectl create -f mosquitto-service.yaml
$ kubectl create -f chat-controller.yaml
$ kubectl create -f chat-service.yaml
```

