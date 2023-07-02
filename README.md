# tictoken-go-bindings

Go Bindings for OpenAI Tictoken using [tiktoken-rs](https://github.com/zurawiki/tiktoken-rs). The tiktoken-rs is built using the official OpenAI tictoken implementations in Rust, so there should not have any difference compared to the official implementations.

The Golang and Rust setup are mostly from [Daulet's Golang bindings](https://github.com/daulet/tokenizers) for the Huggingface Tokenizers library. It demonstrates how to pass data between Golang and Rust.

## How to play with the code

To build, run `make build`
To run unit tests, run `make test`
To format the code, run `make format`

## Potential Issues

1. `ld: symbol(s) not found for architecture arm64`

Add the following to `~/.cargo/config.toml`
```shell
[target.aarch64-apple-darwin]
rustflags = [
  "-C", "link-arg=-undefined",
  "-C", "link-arg=dynamic_lookup",
]
```

2. Set up Goproxy in China

Run `export GOPROXY=https://goproxy.cn/`

## Reference

1. [Rust + Go Examples](https://github.com/mediremi/rust-plus-golang)
1. [Huggingface Golang Bindings by daulet](https://github.com/daulet/tokenizers)
1. [tictoken-go](https://github.com/pkoukk/tiktoken-go)
