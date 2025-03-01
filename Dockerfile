FROM nginx:alpine3.21

RUN mkdir -p /app

COPY dist /app
COPY nginx.conf /etc/nginx/conf.d/default.conf

CMD ["nginx", "-g", "daemon off;"]

EXPOSE 80
