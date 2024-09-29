FROM alpine
RUN apk update && apk upgrade
WORKDIR /app
VOLUME ["/app/db"]
# COPY db/note.db db/
RUN mkdir -p /app/src/as
COPY src/as/ind.css /app/src/as/
COPY todo .
EXPOSE 8080
CMD [ "/app/todo" ]
