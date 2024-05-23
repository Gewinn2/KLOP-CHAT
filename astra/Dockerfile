FROM mobiledevops/flutter-sdk-image AS build

# Set the working directory in the Docker image
WORKDIR /app

# Copy the Flutter app files to the Docker image
COPY . .

# Get Flutter dependencies
RUN flutter pub get

# Build the Flutter web app
RUN flutter build web

# Use the nginx image to serve the Flutter web app
FROM nginx:alpine

# Copy the Flutter web app build output to the nginx server
COPY --from=build /app/build/web /usr/share/nginx/html

# Expose port 80 to the Docker network
EXPOSE 80

# Start nginx
CMD ["nginx", "-g", "daemon off;"]