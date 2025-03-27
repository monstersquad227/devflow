FROM alpine

COPY ./entrypoint.sh /

USER root

#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
#	apk add rpm && \
#	chmod +x /entrypoint.sh && \
#	wget -P /opt http://192.168.1.87/java/jdk-17.0.14_linux-x64_bin.rpm && \
#	rpm -ivh /opt/jdk-17.0.14_linux-x64_bin.rpm && \
#	rm -rf /opt/*

CMD ["/entrypoint.sh"]