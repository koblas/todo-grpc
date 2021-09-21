FROM alpine:3.14 AS thttpd

ARG THTTPD_VERSION=2.29

# Install all dependencies required for compiling thttpd
RUN apk add gcc musl-dev make

# Download thttpd sources
RUN wget http://www.acme.com/software/thttpd/thttpd-${THTTPD_VERSION}.tar.gz \
  && tar xzf thttpd-${THTTPD_VERSION}.tar.gz \
  && mv /thttpd-${THTTPD_VERSION} /thttpd

# Compile thttpd to a static binary which we can copy around
RUN cd /thttpd \
  && ./configure \
  && make CCOPT='-O2 -s -static' thttpd

# Create a non-root user to own the files and run our server
RUN adduser -D static

#
# Build Content
#
FROM node:16-alpine3.14 as build
WORKDIR /app
RUN adduser -D static
RUN apk add --no-cache python3-dev g++ gcc make musl-dev

# Better caching of package.json
ADD package.json /tmp/package.json
RUN cd /tmp && npm install
RUN mkdir -p /opt/app && cp -a /tmp/node_modules /app/

# Bring in everything now
COPY . .
RUN npm run build

FROM scratch
# Copy over the user
COPY --from=thttpd /etc/passwd /etc/passwd

# Copy the thttpd static binary
COPY --from=thttpd /thttpd/thttpd /

# Use our non-root user
USER static
WORKDIR /home/static

# Copy the static website
# Use the .dockerignore file to control what ends up inside the image!
COPY --from=build /app/dist .

EXPOSE 80

# Run thttpd
CMD ["/thttpd", "-D", "-h", "0.0.0.0", "-p", "80", "-d", "/home/static", "-u", "static", "-l", "-", "-M", "60"]
