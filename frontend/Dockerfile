FROM node:current-alpine as build

# Declare environment
ENV NODE_ENV=production

# Set destination for COPY
WORKDIR /app

# Install dependencies
COPY package.json ./
RUN npm install

# Copy the source code
COPY . . 

# Build
RUN npm run build

FROM nginx:mainline-alpine

# Copy built files to nginx
COPY --from=build /app/build /usr/share/nginx/html

EXPOSE 80
