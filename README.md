# goserver
A dead simple, stupid, http service implemented in a complicated way just for the sake of following Go design patterns and scalability.
Useful for learning and testing basic kubernetes networking.
Made on an insomniac night.

### HOW TO GET IT
* Download binary from releases.
* Get container image from dockerhub: https://hub.docker.com/repository/docker/delusionaloptimist/goserver

### DEPLOY
1. For those who don't like containers
	```bash
	go build -o goserver main.go
	./goserver --port 8080
	```

2. For those who like docker
	```bash
	docker run --rm --name=goserver -p 8080:8080 delusionaloptimist/goserver:latest
	```

3. For those who are into weird shit like kubernetes
Deploy a cluster. See [Minikube](https://minikube.sigs.k8s.io/docs/)
	```bash
	kubectl apply -f manifests/k8s_deploy.yaml
	kubectl apply -f manifests/k8s_service.yaml
	```

### HTTP METHODS
```
ENDPOINT '\'
POST: Echo and log the request body
GET: Return and log the request number
```

### INTERACT
The default port which this will listen on is 8080.
You can change it with the `--port` flag and/or change one of the manifest accordingly.
```bash
curl localhost:<port>
```
```bash
curl localhost:<port> -d 'something'
```
