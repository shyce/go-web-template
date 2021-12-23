FROM node:lts-alpine3.14

ARG GID
ARG UID

RUN apk update
RUN apk add jq moreutils shadow git

RUN deluser --remove-home node \
  && groupadd -g ${GID} node \
  && useradd -s /bin/sh -m -g ${GID} -u ${UID} node

RUN mkdir /build
COPY entrypoint.sh /build/entrypoint.sh
RUN chmod +x /build/entrypoint.sh

USER node

ENTRYPOINT ["/build/entrypoint.sh"]
CMD ["npm", "start"]