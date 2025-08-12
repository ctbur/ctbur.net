FROM gcr.io/distroless/static-debian12

WORKDIR /

COPY bin/server ./server

COPY static ./static
COPY templates ./templates
COPY content ./content

EXPOSE 8080

CMD ["./server"]
