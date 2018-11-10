FROM jenkins/jenkins
USER root

ENV KUBECTL_VERSION=v1.12.2

RUN apt-get update  \
    && apt-get -y install \
       apt-transport-https \
       ca-certificates \
       curl \
       software-properties-common

RUN curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add - \
    && add-apt-repository \
       "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
       xenial \
       stable" \
    && apt-get update

RUN apt-get -y install docker-ce
RUN addgroup --gid 1001 dsock
RUN gpasswd -a jenkins dsock 

RUN curl -L https://storage.googleapis.com/kubernetes-release/release/${KUBECTL_VERSION}/bin/linux/amd64/kubectl -o /usr/local/bin/kubectl \
    && chmod +x /usr/local/bin/kubectl

COPY plugins.txt /usr/share/jenkins/ref/plugins.txt
RUN /usr/local/bin/install-plugins.sh < /usr/share/jenkins/ref/plugins.txt

USER jenkins
