FROM centurylink/ca-certs
ADD stablehand /

ENTRYPOINT /stablehand

ENV CATTLE_URL="http://localhost/"
ENV CATTLE_ACCESS_KEY=""
ENV CATTLE_SECRET_KEY=""

ARG VERSION="development"
LABEL VERSION=$VERSION

CMD ["stablehand"]