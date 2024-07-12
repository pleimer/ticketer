FROM alpine:latest

RUN apk add curl

WORKDIR /usr/temporal

RUN curl -sSf https://temporal.download/cli.sh | sh

ENV PATH="${PATH}:/root/.temporalio/bin"
RUN echo $PATH

EXPOSE 7233
EXPOSE 8233

CMD ["temporal", "server", "start-dev", "--ip", "0.0.0.0", "--ui-ip", "0.0.0.0", "--db-filename", "/usr/temporal/temporal.db"]