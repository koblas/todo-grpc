
#
# Build Content
#
FROM node:16-alpine3.14 as build
WORKDIR /app
RUN apk add --no-cache python3-dev g++ gcc make musl-dev
RUN adduser -D static

# Better caching of package.json
ADD package.json /tmp/package.json
RUN cd /tmp && npm install --legacy-peer-deps
RUN mkdir -p /opt/app && cp -a /tmp/node_modules /app/

# Bring in everything now
COPY . .
RUN npm run build

#
#  Build a scratch image with nginx and our content
#
FROM alpine:latest
# Install all dependencies required for compiling thttpd
RUN apk update && apk add --no-cache nginx

# Create a non-root user to own the files and run our server
RUN adduser -D static

COPY nginx.conf /etc/nginx/http.d/default.conf

# Use our non-root user
#USER static
WORKDIR /home/static

# Copy the static website
# Use the .dockerignore file to control what ends up inside the image!
COPY --from=build /app/dist .

# Change default stop signal from SIGTERM to SIGQUIT
STOPSIGNAL SIGQUIT

EXPOSE 80

# Define entrypoint and default parameters
CMD ["/usr/sbin/nginx","-g", "daemon off;"]
