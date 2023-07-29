FROM node:20-alpine
RUN npm i docsify-cli -g
COPY docs /docsify/docs
WORKDIR /docsify
CMD ["/usr/local/bin/docsify", "serve", "docs"]