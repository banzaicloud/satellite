![license](http://img.shields.io/badge/license-Apache%20v2-orange.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/banzaicloud/whereami)](https://goreportcard.com/report/github.com/banzaicloud/whereami)
[![Docker Automated build](https://img.shields.io/docker/automated/banzaicloud/whereami.svg)](https://hub.docker.com/r/banzaicloud/whereami/)

*NOAA is the latest generation of polar-orbiting, non-geosynchronous, environmental satellites. NOAA-20 was launched on November 18, 2017 to give meteorologists information on clouds (among many others) and surfers on waves.*

*NOAA is a Golang library and RESTful API to determine the host cloud provider with a simple HTTP call. Behind the scenes it uses the file system and provider metadata to properly identify the cloud provider.*

---

# NOAA

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

## Contributing

If you find this project useful here's how you can help:

- Send a pull request with your new features and bug fixes
- Help new users with issues they may encounter
- Support the development of this project and star this repo!
