# Use an official Node.js runtime as the base image
FROM node:14

# Set the working directory in the container
WORKDIR /node-service

# Copy the package.json and package-lock.json files to the working directory
COPY package*.json ./

# Install project dependencies
RUN npm install

# Copy all the application files to the working directory
COPY . .

# Expose the port that your application will run on
EXPOSE 8000

# Define the command to start your application
CMD ["node", "server.js"]