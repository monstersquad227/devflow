FROM nginx:latest

LABEL maintainer="NGINX"
LABEL description="The startup user is root"
LABEL author="monstersquad227"
LABEL version="1.23.3"


USER root

RUN sed -i "s#nginx;#root;#g" /etc/nginx/nginx.conf && \
        rm -rf /etc/nginx/conf.d/*

COPY ./default.conf /etc/nginx/conf.d/