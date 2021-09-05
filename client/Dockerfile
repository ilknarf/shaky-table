# multi-stage build for client application
FROM node:16-alpine
WORKDIR /client

COPY ./package-lock.json .
COPY ./package.json .
RUN npm i

COPY . .
# SvelteKit builds a self-contained node application
RUN npm run build

FROM node:16-alpine
WORKDIR /app
COPY --from=0 /client .
EXPOSE 3000
CMD ["node", "./build"]
