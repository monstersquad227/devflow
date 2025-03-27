FROM nginx:1.20

USER root

WORKDIR /opt

RUN sed -i "s#nginx;#root;#g" /etc/nginx/nginx.conf && \
    rm -rf /etc/nginx/conf.d/default.conf || true

COPY ./nginx.conf /etc/nginx/conf.d/
