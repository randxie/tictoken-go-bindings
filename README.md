# tictoken-go-bindings

Go Bindings for OpenAI Tictoken

## Potential Issues

1. `ld: symbol(s) not found for architecture arm64`

Add the following to `~/.cargo/config.toml`
```shell
[target.x86_64-apple-darwin]
rustflags = [
  "-C", "link-arg=-undefined",
  "-C", "link-arg=dynamic_lookup",
]

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