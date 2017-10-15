FROM docker:dind
LABEL MAINTAINER Alex Jones <tibbar@freedommail.ch>
ENV GOPATH=/go
WORKDIR /go/src/github.com/AlexsJones/kubebuilder
COPY . .

RUN set -x && \
  apk --update --no-cache add ca-certificates git wget python go musl-dev && \
  cd /tmp && \
  wget https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-172.0.1-linux-x86_64.tar.gz && \
  tar -xvf google-cloud-sdk-172.0.1-linux-x86_64.tar.gz && \
  mv google-cloud-sdk / && \
  CLOUDSDK_CORE_DISABLE_PROMPTS=1 /google-cloud-sdk/install.sh && \
  rm -rfv /tmp/* && \
  /google-cloud-sdk/bin/gcloud components update

ENV PATH ${PATH}:/google-cloud-sdk/bin:${GOPATH}/bin

RUN go get github.com/tools/godep
RUN cd /go/src/github.com/AlexsJones/kubebuilder && godep restore && go build


ENTRYPOINT ["kubebuilder"]
