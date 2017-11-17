# kubebuilder
<img src=https://i.imgur.com/xxDRsik.jpg width="250" />

[![GitHub license](https://img.shields.io/github/license/AlexsJones/kubebuilder.svg)](https://github.com/AlexsJones/kubebuilder/blob/master/LICENSE)

I initially built kubebuilder for the developers at [BeameryHQ](https://github.com/BeameryHQ)
It is a deployment for kubernetes that lets developers have minimal interaction but have their own subdomains
and namespaces for creating applications.

It requires a few pieces to get it all working, but is designed to be simple for the end user.

### Requirements

- protobuf-compiler
- sshkey for VCS (git currently)
- Requires a json service account file for GCP named auth
- Requires an ssh key that works as a secret named key
- kepler for CLI interaction `go get https://github.com/AlexsJones/kepler`


### Usage

![Diagram](https://i.imgur.com/Ukf7vF2.jpg)
