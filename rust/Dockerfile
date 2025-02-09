FROM rust:1 AS build

WORKDIR /app
COPY . .
RUN cargo build --release

FROM debian:stable-slim
COPY --from=build /app/target/release/ /app/

CMD ["./app/Rust"]
