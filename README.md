# cobruh

A small Cobra + Viper CLI for learning how command-line tools handle configuration.
It exposes a `serve` command whose port can be set from a flag, an environment variable, a config file, or a built-in default.

## Config precedence

Flags (`--port`) > Environment variables (`COBRUH_PORT`) > Config file (`config.yaml`) > Default (`8080`)
