FROM rust:1

WORKDIR /app
COPY Cargo.toml .

RUN cargo build --release
COPY target/release/* .

CMD ["./Rust"]
