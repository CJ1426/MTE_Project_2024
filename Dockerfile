FROM alpine
RUN apk update && apk upgrade
WORKDIR /app
VOLUME ["/app/db"]
COPY todo .
EXPOSE 8080
CMD [ "/app/todo" ]
