FROM ubuntu:latest
LABEL authors="whither"

ENTRYPOINT ["top", "-b"]