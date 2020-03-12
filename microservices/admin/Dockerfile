FROM node as builder
WORKDIR /app
COPY src/ ./src/
COPY webpack.config.js ./
COPY package.json ./
COPY yarn.lock ./
RUN yarn install
RUN yarn run webpack

FROM nginx:1.17
COPY --from=builder /app/public/bundle.js /usr/share/nginx/html/
COPY public/index.html /usr/share/nginx/html/
