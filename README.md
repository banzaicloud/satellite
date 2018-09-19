![license](http://img.shields.io/badge/license-Apache%20v2-orange.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/banzaicloud/noaa)](https://goreportcard.com/report/github.com/banzaicloud/noaa)
[![Docker Automated build](https://img.shields.io/docker/automated/banzaicloud/noaa.svg)](https://hub.docker.com/r/banzaicloud/noaa/)

*NOAA is the latest generation of polar-orbiting, non-geosynchronous, environmental satellites. NOAA-20 was launched on November 18, 2017 to give meteorologists information on clouds (among many others) and surfers on waves.*

*NOAA is a Golang library and RESTful API to determine the host cloud provider with a simple HTTP call. Behind the scenes it uses the file system and provider metadata to properly identify the cloud provider.*

---

# NOAA

<p align="center">
<img src="docs/images/noaa_logo.jpg">
</p>

NOAA is widely used across the [Pipeline](https://github.com/banzaicloud/pipeline) platform. We are cloud agnostic but at the same time cloud aware - using the [Banzai Cloud](https://banzaicloud.com) Kubernetes operators all our code relies on NOAA to determine the **cloud provider** and *inject* the cloud specific code.

With this simple service the cloud provider can be determined easily with a simple HTTP call.

## Supported Cloud Providers

- Amazon
- Alibaba Cloud
- Azure
- Google Cloud
- Oracle Cloud
- DigitalOcean

## API

- GET "/noaa" -> {name: cloudprovider}

## Helm chart

You can deloy Noaa to a cloud based Kubernetes cluster using this [Helm chart](https://github.com/banzaicloud/banzai-charts/tree/master/noaa). 

To install the chart from the Banzai Cloud chart repo:

```
$ helm repo add banzaicloud-stable http://kubernetes-charts.banzaicloud.com/branch/master
$ helm install banzaicloud-stable/noaa
```

## Examples

The Banzai Cloud PVC operator uses Noaa in order to keep Helm charts cloud `agnostic`. For further details please check the [PVC operator](https://github.com/banzaicloud/pvc-operator/blob/master/README.md).

This operator makes using [Kubernetes Persistent Volumes](https://kubernetes.io/docs/concepts/storage/persistent-volumes/) easier on cloud providers, by dynamically creating the required accounts, classes and more. It allows to use exactly the same [Helm](https://helm.sh) chart on all the supported providers, there is no need to create cloud specific Helm charts.


## Contributing

If you find this project useful here's how you can help:

- Send a pull request with your new features and bug fixes
- Help new users with issues they may encounter
- Support the development of this project and star this repo!
