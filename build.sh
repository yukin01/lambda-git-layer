#!/bin/bash
set -eu

rm -r layer || true

docker run --rm -it -v "$PWD"/layer:/tmp/layer --entrypoint "" public.ecr.aws/lambda/go /bin/bash -c "
  yum install -y zip git && \
  mkdir -p /opt/bin && \
  cp /usr/bin/git /opt/bin && \
  cd /opt && \
  zip -yr /tmp/layer/git.zip .
"
