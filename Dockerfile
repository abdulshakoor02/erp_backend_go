FROM golang:1.21.3

ENV SUPPORTING_FILES /app
ARG DEV

# install bash for wait-for-it script
RUN apt update && apt install -y bash nano postgresql-client git

RUN mkdir -p $SUPPORTING_FILES

WORKDIR $SUPPORTING_FILES

COPY . $SUPPORTING_FILES

