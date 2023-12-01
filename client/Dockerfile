# Stage 1: Build
FROM node:16 as build

WORKDIR /app
COPY client/package*.json client/yarn*.lock ./
RUN npm install
COPY client/ .
RUN npm run build

# Stage 2: Serve
FROM nginx:alpine
COPY --from=build /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
