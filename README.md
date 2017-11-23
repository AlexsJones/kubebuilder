# kubebuilder
<img src=https://i.imgur.com/xxDRsik.jpg width="250" />
=======
[![GitHub license](https://img.shields.io/github/license/AlexsJones/kubebuilder.svg)](https://github.com/AlexsJones/kubebuilder/blob/master/LICENSE)

I initially built kubebuilder for the developers at [BeameryHQ](https://github.com/BeameryHQ)
It is a deployment for kubernetes that lets developers have minimal interaction but have their own subdomains
and namespaces for creating applications.

It requires a few pieces to get it all working, but is designed to be simple for the end user.

### Development Requirements

- protobuf-compiler
- kepler for CLI interaction `go get https://github.com/AlexsJones/kepler`

### Deployment Requirements

Place google cloud service account in `k8s/required/files/auth.json`
By Default will use current ssh key at `~/.ssh/id_rsa`

Run `./scripts/deployment_kubernetes.sh`

### Usage

![Diagram](https://i.imgur.com/Ukf7vF2.jpg)
