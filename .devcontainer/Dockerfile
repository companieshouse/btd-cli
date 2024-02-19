FROM rockylinux:9-minimal

ENV TERM=xterm-256color

RUN microdnf upgrade -y \
  && microdnf install -y epel-release \
  && microdnf install -y git go tar vim \
  && microdnf clean all -y

