# VistA Microservices using Octo and Go

This repository contains the files used to bring up a VistA container in a Kubernetes pod, then create and bring up a microservice implemented in go.

For the purposes of demonstration, we create a web page which renders a single table, which has the results of:

```
SELECT h1.NAME, p1.NAME, p1.SEX, p1.AGE, v1.CHECK_OUT_DATE_TIME, v1.VISIT_ADMIT_DATE_TIME
FROM HOSPITAL_LOCATION h1
INNER JOIN VISIT v1 ON (v1.HOSPITAL_LOCATION = h1.HOSPITAL_LOCATION_ID)
INNER JOIN PATIENT p1 ON (v1.PATIENT_NAME = p1.PATIENT_ID) 
WHERE h1.NAME = 'ICU/CCU'
```

# Usage

## Step 0

Install minikube, install kubernetes, install Docker

During the demo, a local Docker registry was used. This was an insecure registry hosted directly on my machine.
Below, we fetch the IP of an interface which we can use to tell the Kubernetes containers to use when fetching docker images; this should be updated for each demo to reflect the IP of the host.

## Step 1 - Setup a minikube environment

```
minikube delete
minikube start --insecure-registry $(ifconfig | grep -C 1 wlp0 | grep inet | awk '{print $2}'):5000
```

## Step 2 - Build Docker image

Build the docker images; again; this is using a local registry. Make sure you configure Docker to allow insecure registries

```
docker build -t localhost:5000/go-microservice docker
docker push localhost:5000/go-microservice
```

If you need an insecure registry, try making a file /etc/docker/daemon.json:

```
charles@mcom ~/p/vista-go-microservice-demo $ cat /etc/docker/daemon.json 
{
  "insecure-registries" : ["127.0.0.1:5000"]
}
```

## Step 3 - Setup VistA Kubernetes Image

Ensure that the VistA/rocto Docker image has been uploaded to the local registry; update pod.yml to reflect that by replacing 172.22.124.139 with an IP which can be resolved to the local registry.

```
kubectl apply -f claim-pv.yml
kubectl apply -f claim.yml
kubectl apply -f pod.yml
kubectl apply -f service.yml

# Wait for the pod to come up; check with
kubectl get pods

minikube service vista-deployment --url
```

The last command will output a string of the form ``http://192.168.11.1:12345``; copy the IP and port number to verify SSH works:

```
ssh -lroot -p12345 192.168.11.1
# Password is 'docker'
```

It will display two output strings; one is for rocto, one for ssh. You could try to determine which is which by looking at the output of `kubectl get services`, or by just trying to connect to each.

If you can login, things are good

## Step 4 - Setup the Micro service

Make sure docker-pod.yml has been updated to have your current IP, rather than `172.22.124.139`, under spec.template.spec.containers.image.

```
kubectl apply -f go-service.yml
kubectl apply -f docker-pod.yml

kubectl expose deployment go-microservice-deployment --type=NodePort

minikube service go-microservice-deployment --url
```

And you should be able to access the web page displayed by the last command.
