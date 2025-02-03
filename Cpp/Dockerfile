# Stage 1: Build stage
FROM ubuntu:latest AS build

# Install build-essential for compiling C++ code
RUN apt-get update && apt-get install -y build-essential

# Set the working directory
WORKDIR /app

# Copy the source code into the container
COPY main.cpp .

# Compile the C++ code statically to ensure it doesn't depend on runtime libraries
RUN g++ -O2 -std=c++20 -pedantic -Wall -Wextra -Werror -Wshadow -Wconversion -Wunreachable-code -o main.out main.cpp -static

# Stage 2: Runtime stage
FROM scratch

# Copy the static binary from the build stage
COPY --from=build /app/main.out /main.out

# Command to run the binary
CMD ["./main.out"]