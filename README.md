# kubebuilder

![Imgur](https://i.imgur.com/xxDRsik.jpg)


A kubernetes deployment that allows developers to push their repositories via pubsub and have it deployed.
The purpose here is to let multiple developers have their own namespaces in a k8s cluster to work on, shared by a common ingress.

---

## Development

### Requirements

- protobuf-compiler
- sshkey for VCS (git currently)
- Requires a json service account file for GCP 


### Usage

![Diagram](https://i.imgur.com/Ddn3oyU.png)
