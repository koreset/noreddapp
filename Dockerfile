FROM ubuntu:18.04

WORKDIR /app

COPY noredd-app .
COPY public/ public/
COPY templates/ templates/
COPY assets/ assets/

ENV PORT=5000
ENV GIN_MODE=release
ENV DBPASSWD=noredduser
ENV DBUSER=noredduser
ENV DBHOST=104.248.255.136
EXPOSE 5000

ENTRYPOINT ["./noredd-app"]
