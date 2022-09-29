
#
# Build Content
#
FROM node:18-alpine3.16 as build
WORKDIR /app
RUN apk add --no-cache python3-dev g++ gcc make musl-dev
RUN adduser -D static

# Better caching of package.json
ADD package.json /tmp/package.json
RUN cd /tmp && npm install 
RUN mkdir -p /opt/app && cp -a /tmp/node_modules /app/

# Bring in everything now
COPY . .
RUN npm run build

#
#  Build a scratch image with nginx and our content
#
FROM caddy:2-alpine

RUN mkdir -p /home/static

COPY Caddyfile /etc/caddy/Caddyfile

# Copy the static website
COPY --from=build /app/dist /home/static