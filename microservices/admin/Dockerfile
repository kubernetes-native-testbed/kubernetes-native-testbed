FROM node as builder
WORKDIR /app
COPY admin/src/ ./src/
COPY admin/webpack.config.js ./
COPY admin/package.json ./
COPY admin/yarn.lock ./
RUN yarn install
RUN yarn run webpack

FROM nginx:1.17
COPY --from=builder /app/public/bundle.js /usr/share/nginx/html/
COPY admin/public/ /usr/share/nginx/html/





