# Todos Restful Api

[![Build Status](https://travis-ci.com/bilalislam/todo-app-api.svg?branch=master)](https://travis-ci.com/bilalislam/todo-app-api)

![Random GIF](./images/swagger.png)


Todo-api  is a backend project developed with golang and developed as the api of frontend.

You can reach ui github repo [link.](https://github.com/bilalislam/todo-app-ui)


# Demo-Preview

swagger

# Table of contents

- [Todos](#todos)
- [Demo-Preview](#demo-preview)
- [Table of contents](#table-of-contents)
- [Installation](#installation)
- [Usage](#usage)
- [Deployment](#deployment)
- [Known Issues](#known-issues)
- [Architecure](#architecure)
- [References](#references)

# Installation
[(Back to top)](#table-of-contents)

To use this project, first clone the repo on your device using the command below:

```git init```

```git clone https://github.com/bilalislam/todo-app-api```


# Usage
[(Back to top)](#table-of-contents)

### `build`

```sh
$ cd todo-app-api
$ go build main.go -o api
```

### `api start`

```sh
$ ./api
```

Runs the app in the development mode.<br />
Open [http://localhost:8080](http://localhost:3000) to view it in the browser.

The page will reload if you make edits.<br />
You will also see any lint errors in the console.

### `unit test`

```sh
$ go test
```

# Deployment

[(Back to top)](#table-of-contents)

### `docker compose for local build`

Compose is a tool for defining and running multi-container Docker applications. With Compose, you use a YAML file to configure your application’s services. Then, with a single command, you create and start all the services from your configuration. To learn more about all the features of Compose, [see the list of features ](https://docs.docker.com/compose/#features)

A docker-compose.yml looks like:

```docker

version: '3'

services:
  web:
    image: ninjafx/todo-app-ui:latest
    ports:
      - "5000:5000"
    depends_on:
      - api

  api:
    image: ninjafx/todo-app-api:latest
    ports:
      - "8080:8080"

```

```sh
$ docker-compose up
```

Open [http://localhost:8080](http://localhost:5000) to view it in the browser.

### `kompose`

Kompose is a conversion tool for Docker Compose to container orchestrators such as Kubernetes.

We are going to create Kubernetes Deployments, Services  for your Dockerized application. 

If you need different kind of resources, use the 'kompose convert' and 'kubectl create -f' commands instead. 


```sh
$ kompose convert  -f docker-compose.yaml
or
$ kompose --file docker-compose.yml up 
```

### `aws eks`

Amazon Elastic Kubernetes Service (Amazon EKS) is a fully managed Kubernetes service.

Getting started with eksctl: This getting started guide helps you to install all of the required resources to get started with Amazon EKS using eksctl, a simple command line utility for creating and managing Kubernetes clusters on Amazon EKS. At the end of the tutorial, you will have a running Amazon EKS cluster that you can deploy applications to. This is the fastest and simplest way to get started with Amazon EKS.

simple usage ;

```sh
$ eksctl create cluster --name todo-app --nodes 1 --node-type t2.medium --region eu-west-1
```

After created our  cluster and we are ready for deployment with kubectl .

kubectl controls the Kubernetes cluster manager.

```sh
$ kubectl apply -f api-deployment.yaml
$ kubectl apply -f api-service.yaml
```

Open [http://ec2-ip:8080](http://ec2-ip:8080) to view it in the browser.

### `travis ci`

Or travis ci  all written the above things will run in .travis.yml. After commit,It will build  dockerfile and push the image to aws container service. But this ci/cd pipeline could not completed yet.


# Known Issues

[(Back to top)](#table-of-contents)

1. it could not include pact for mocking a-tdd not yet
2. it could not include a-tdd process in ci/cd pipeline not yet

# Architecure

[(Back to top)](#table-of-contents)

This project has been created inspired by clean architecture.

Open [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/) to view it in the browser.

# References

[(Back to top)](#table-of-contents)

1. https://kompose.io/
2. https://eksctl.io/
3. https://drive.google.com/file/d/1slZgPh8yOvhGC_r4wQpvYC-OSDw1I__0/view
4. https://travis-ci.com/
5. https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html
6. https://github.com/vektra/mockery
7. https://github.com/swaggo/swag
