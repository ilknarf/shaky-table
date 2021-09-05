FROM node:16-alpine
WORKDIR /client
COPY . .
RUN npm i
RUN npm run build

FROM node:16-alpine
WORKDIR /app
COPY --from=0 /client .
EXPOSE 3000
CMD ["node", "./build"]
