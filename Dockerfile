FROM node As builder

MAINTAINER "lius"

WORKDIR /app

COPY . .

RUN npm install -g pnpm \
    && pnpm install \
    && pnpm run build \
    && rm -rf node_modeuls

FROM nginx

COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
