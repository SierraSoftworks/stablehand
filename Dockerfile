FROM scratch
ADD ca-certificates.crt /etc/ssl/certs
ADD stablehand /

ENTRYPOINT /stablehand

ENV RANCHER_SERVER
ENV RANCHER_ACCESS_KEY
ENV RANCHER_SECRET_KEY

CMD ["stablehand"]