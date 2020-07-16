FROM alpine
ADD post_svc-service /post_svc-service
ENTRYPOINT [ "/post_svc-service" ]
