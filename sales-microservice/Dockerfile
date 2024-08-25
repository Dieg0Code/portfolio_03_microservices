FROM node:19-alpine AS builder

LABEL mantainer="github.com/dieg0code"

WORKDIR /app

COPY package*.json ./
COPY tsconfig.json ./

RUN npm install

COPY src/ ./src/
COPY index.ts ./

RUN npm run build

FROM node:19-alpine

LABEL mantainer="github.com/dieg0code"

WORKDIR /app

COPY --from=builder /app/dist/ ./dist/
COPY --from=builder /app/node_modules/ ./node_modules/
COPY --from=builder /app/package*.json ./

EXPOSE 8082

CMD ["node", "dist/index.js"]