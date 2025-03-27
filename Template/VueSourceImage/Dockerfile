FROM nginx:1.20

USER root

WORKDIR /opt

RUN rm -rf /etc/nginx/nginx.conf && \
    rm -rf /etc/nginx/conf.d/

COPY ./nginx.conf /etc/nginx/
