FROM rust:1 AS build

WORKDIR /app
COPY . .
RUN cargo build --release

FROM scratch
COPY --from=build /app/target/release/* .

CMD ["./Rust"]
