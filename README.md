# yundict-cli

Yundict CLI is a command line tool for [Yundict](https://yundict.com), easily export your translation to your local machine.

## Documentation

For more information, please refer to [Yundict CLI Documentation](https://yundict.com/docs/cli/).

## Installation

### macOS

```bash
brew tap yundict/yundict-cli
brew install yundict
```

### Linux (TODO)

Install with `curl`:

```bash
curl -sSL https://raw.githubusercontent.com/yundict/yundict-cli/main/install.sh | sh
```

## Usage

#### export

Export translations from Yundict to your local machine.

```bash
yundict export \
  --token <token> \
  --team <teamName> \
  --project <projectName> \
  --type key-value-json \
  --languages en \
  --out ./en.json
```