FROM openjdk:17

LABEL maintainer="OPENJDK"
LABEL description="The startup script is /entrypoint.sh"
LABEL author="monstersquad227"
LABEL version="17.0.2"

COPY ./entrypoint.sh /

USER root

ENTRYPOINT ["/bin/bash", "/entrypoint.sh"]