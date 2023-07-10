<p align="center"><b>Chrono: The ultimate Kubernetes operator for scheduling jobs across nodes.</b></p>

<h4 align="center">
    <a href="https://github.com/jebinjeb/chrono/discussions">Discussions</a> 
</h4>

<h4 align="center">

[![Docker Image CI](https://github.com/jebinjeb/chrono/actions/workflows/docker-image.yaml/badge.svg)](https://github.com/jebinjeb/chrono/actions/workflows/docker-image.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/jebinjeb/chrono)](https://goreportcard.com/report/github.com/jebinjeb/chrono)


[![Price](https://img.shields.io/badge/price-FREE-0098f7.svg)](https://github.com/jebinjeb/chrono/blob/main/LICENSE)
[![Discussions](https://badgen.net/badge/icon/discussions?label=open)](https://github.com/jebinjeb/chrono/discussions)
[![Code of Conduct](https://badgen.net/badge/icon/code-of-conduct?label=open)](./code-of-conduct.md)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

</h4>

<hr>

# chrono

## Features of Chrono:
- Node-Aware Job Scheduling: Chrono allows you to schedule jobs across nodes in your Kubernetes cluster, taking into account the availability and resources of each node. It ensures optimal utilization of your cluster resources by intelligently distributing jobs.
- Node Selector Support: With Chrono, you can schedule jobs based on node selectors, allowing you to target specific nodes or groups of nodes in your cluster. This gives you fine-grained control over job placement and resource allocation.
- Enhanced Job Monitoring: Chrono provides advanced job monitoring capabilities, giving you real-time insights into the status and progress of your scheduled jobs. You can easily track job execution, view logs, and receive notifications for job completions or failures.
- Scalability and Reliability: Chrono is designed to handle large-scale job scheduling in Kubernetes clusters. It leverages the scalability and reliability features of Kubernetes, ensuring efficient and robust execution of scheduled jobs.

## Problems Addressed by Chrono:
- Limited Job Scheduling Options: Kubernetes provides basic job scheduling capabilities through Jobs and CronJobs, but they have limitations when it comes to node-aware scheduling and node selection. Chrono addresses these limitations by offering advanced scheduling features.


###  How to install and run chrono:

#### Prerequisites
* A Kubernetes cluster 
* Helm binary

#### Prepare Namespace
```bash
kubectl create namespace chrono
```

#### Client Installation
```bash
helm repo add chrono https://jebinjeb.github.io/chrono/
helm repo update

helm upgrade -i chrono chrono/chrono -n chrono 
```
