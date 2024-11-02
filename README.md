# Yundict CLI

Yundict CLI is a powerful command-line interface designed to enhance your experience with [Yundict](https://yundict.com/). This tool provides a comprehensive set of commands to streamline the management of your translations, making it easier than ever to integrate with your current workflows.

Whether you're handling daily translation tasks or automating processes within CI/CD pipelines, Yundict CLI ensures smooth and efficient synchronization of your project's multilingual content. You can easily export translations to your local environment or import updated strings back into Yundict, simplifying the process of managing and updating text assets across multiple languages.

## Documentation

For more information, please refer to [Yundict CLI Documentation](https://yundict.com/docs/cli/).

## Installation

Yundict CLI is a single executable file. It can be installed via Homebrew on macOS or downloaded from the [release page](https://github.com/yundict/yundict-cli/releases).

### macOS

Install via Homebrew:

```bash
brew tap yundict/cli
brew install yundict
```

### Linux

Download the latest release from the [release page]((https://github.com/yundict/yundict-cli/releases)).

## Usage

### Export translations

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

- `--token`: Your Yundict API token, get it from [Yundict API Token](https://app.yundict.com/settings/api-tokens).
- `--team`: Your team name in Yundict.
- `--project`: Your project name in Yundict.
- `--type`: The type of the export file, support `key-value-json`, `csv`, `android-xml`, `apple-strings`...
- `--languages`: The languages you want to export, separated by commas. e.g. `en,ja,zh-CN`.
- `--out`: The output file path.