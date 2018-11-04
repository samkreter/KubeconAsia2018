FROM ubuntu:16.04

RUN apt update
RUN apt install -y curl

# Install pachctl 
RUN curl -o /tmp/pachctl.tar.gz -L https://github.com/pachyderm/pachyderm/releases/download/v1.7.10/pachctl_1.7.10_linux_amd64.tar.gz && tar -xvf /tmp/pachctl.tar.gz -C /tmp && cp /tmp/pachctl_1.7.10_linux_amd64/pachctl /usr/local/bin

RUN alias p='pachctl'

COPY data /data
COPY setup.sh /
COPY update.sh /

RUN chmod a+x /setup.sh
RUN chmod a+x /update.sh