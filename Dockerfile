FROM  --platform=linux/arm64 hephy/go-dev
# This Dockerfile is used to bundle the source and all dependencies into an image for testing.

RUN echo "deb http://packages.cloud.google.com/apt cloud-sdk-jessie main" \
  | tee /etc/apt/sources.list.d/google-cloud-sdk.list \
  && curl -k https://packages.cloud.google.com/apt/doc/apt-key.gpg \
  | apt-key add - \
  && apt-get update \
  && apt-get install -y google-cloud-sdk \
  --no-install-recommends \
  && rm -rf /var/lib/apt/lists/*

ENV CGO_ENABLED=0

ADD https://codecov.io/bash /usr/local/bin/codecov
RUN chmod +x /usr/local/bin/codecov

WORKDIR /go/src/github.com/teamhephy/workflow-cli

ENV GO111MODULE=on

COPY ./_scripts /usr/local/bin

COPY . /go/src/github.com/teamhephy/workflow-cli
COPY ./zscaler_root_ca.crt /etc/ssl/certs/zscaler_root_ca.crt

RUN go build
