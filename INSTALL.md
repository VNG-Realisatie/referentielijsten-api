# Installation

Running `VRL` locally can be done by using the following methods:

## docker

Simply run the following command:

```shell
docker run -it -e PORT=:8000 -p 8000:8000 ghcr.io/vng-realisatie/referentielijsten-api:0.1.0
```

The env variable is optional, which means below command is basically the same but it also means you cannot change the
port since 8000 is the default.

```shell
docker run -it -p 8000:8000 ghcr.io/vng-realisatie/referentielijsten-api:0.1.0
```

## kubernetes

Run the following commands:

```shell
cat <<EOF >./deploy.yaml                                          
apiVersion: v1
kind: Service
metadata:
  name:  vrl
  namespace: vrl
spec:
  selector:
    app.kubernetes.io/name: vrl
    app.kubernetes.io/instance: vrl
  ports:
    - name: "http"
      port: 8000
      targetPort: 8000
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name:  vrl
  namespace: vrl
spec:
  replicas: 1
  selector:
    matchLabels:
      app.kubernetes.io/name: vrl
      app.kubernetes.io/instance: vrl
  template:
    metadata:
      labels:
        app.kubernetes.io/name: vrl
        app.kubernetes.io/instance: vrl
    spec:
      containers:
        - name: vrl
          image: ghcr.io/vng-realisatie/referentielijsten-api:0.1.0
          imagePullPolicy: Always
          ports:
            - containerPort: 8000
          env:
            - name: PORT
              value: :8000
          resources:
            requests:
              memory: "50Mi"
              cpu: "50m"
            limits:
              memory: "100Mi"
              cpu: "100m"
          readinessProbe:
            httpGet:
              path: /api/v1/health
              port: 8000
              httpHeaders:
                - name: Host
                  value: localhost
            initialDelaySeconds: 10
            periodSeconds: 10
          livenessProbe:
            httpGet:
              path: /api/v1/health
              port: 8000
              httpHeaders:
                - name: Host
                  value: localhost
            initialDelaySeconds: 15
            periodSeconds: 30
EOF
```

This will save the `yaml` to your current working dir it can then be applied:

```shell
kubectl apply -f deploy.yaml
```

Using this method you will deploy a simple service with a deployment. Change values accordingly.