FROM node:14.8.0-alpine3.10

RUN apk update
RUN apk add jq moreutils
RUN mkdir /build
COPY entrypoint.sh /build/entrypoint.sh
RUN chmod +x /build/entrypoint.sh

ENTRYPOINT ["/build/entrypoint.sh"]
CMD ["npm", "start"]