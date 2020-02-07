FROM node as builder
WORKDIR /app
COPY *.js ./
COPY package.json ./
COPY yarn.lock ./
RUN yarn install
# Install grpc-web
# Install google-protobuf
RUN yarn run webpack

FROM nginx:1.17
RUN mkdir -p /usr/share/nginx/html/client
COPY index.html /usr/share/nginx/html/client
COPY --from=builder /app/dist/main.js /usr/share/nginx/html/client
