FROM gcr.io/distroless/static:nonroot

LABEL maintainer="GO"
LABEL description="Golang Base Image"
LABEL author="monstersquad227"

WORKDIR /opt

COPY ./config ./config
COPY ./main ./

ENTRYPOINT ["./main"]