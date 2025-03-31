# Use Node.js for building frontend
FROM node:20 AS builder

WORKDIR /app

# Copy frontend files
COPY package.json package-lock.json ./
RUN npm install

COPY . .

# Build the frontend
RUN npm run build

# Use Nginx to serve static files
FROM nginx:latest

COPY --from=builder /app/dist /usr/share/nginx/html

# Expose the port for frontend
EXPOSE 80

