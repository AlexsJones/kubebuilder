FROM golang:latest
LABEL MAINTAINER Alex Jones <tibbar@freedommail.ch>
ENV GOPATH=/go
WORKDIR /go/src/github.com/AlexsJones/kubebuilder
COPY . .

RUN set -x && \
  cd /tmp && \
  wget https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-172.0.1-linux-x86_64.tar.gz && \
  tar -xvf google-cloud-sdk-172.0.1-linux-x86_64.tar.gz && \
  mv google-cloud-sdk / && \
  CLOUDSDK_CORE_DISABLE_PROMPTS=1 /google-cloud-sdk/install.sh && \
  rm -rfv /tmp/* && \
  /google-cloud-sdk/bin/gcloud components update

ENV PATH ${PATH}:/google-cloud-sdk/bin:${GOPATH}/bin

RUN go-wrapper download
RUN go-wrapper install
ENTRYPOINT ["kubebuilder","--conf cluster-config.yaml"]
