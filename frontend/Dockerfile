FROM node:16-alpine3.16 AS build

WORKDIR /app

COPY package.json .
COPY package-lock.json .
RUN npm install

COPY . .
RUN npm run build

FROM caddy:2.6.2-alpine AS run

EXPOSE 443

COPY --from=build /app/dist/ /usr/share/caddy




