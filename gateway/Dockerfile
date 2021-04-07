# add the necessary instructions
# to create a Docker container image
# for your Go API gateway server
FROM alpine
RUN apk add --no-cache ca-certificates
COPY  gateway /gateway
EXPOSE 443
ENTRYPOINT ["/gateway"]