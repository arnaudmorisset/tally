# Tally

Tally is a command-line tool written in Go that provides word, line, and byte counting functionality for files.
It serves as an alternative to the traditional UNIX tool `wc`, with added support for exporting results in JSON.

## Features

- Count words, lines, and bytes in a file.
- Export count results in plain text or JSON.
- Simple and efficient command-line interface.

## Usage

```
tally [OPTIONS] FILE

Options:
  -f string
        Output format: text, json (default "text")

Examples:
  # Count words, lines, and bytes in a file (default output format is text)
  tally filename.txt

  # Count words, lines, and bytes in a file and export results in JSON format
  tally -f json filename.txt
```

## Development

### Prerequisites

- Go 1.16 or newer

### Building from Source

```bash
git clone https://github.com/arnaudmorisset/tally.git
cd tally
go build
```

### Running Tests

```bash
go test ./...
```

## Contributing

Contributions are welcome! Please feel free to open issues for feature requests, bug fixes, or general feedback.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Note

This project is primarily educational and intended for learning purposes.
While Tally can be used for basic file counting tasks, it is not intended to replace `wc` for real-world usage scenarios.
For production environments, `wc` or similar established tools are recommended.

In a real world scenario, you will probably use something along the lines of:

```bash
#!/bin/bash

# Run wc -wl command and capture its output
wc_output=$(wc -wlc "$1")

# Extract word count, line count, and byte count using awk
words=$(echo "$wc_output" | awk '{print $1}')
lines=$(echo "$wc_output" | awk '{print $2}')
bytes=$(echo "$wc_output" | awk '{print $3}')

# Format the counts as JSON
json_output="{\"words\": $words, \"lines\": $lines, \"bytes\": $bytes}"

# Print the JSON output
echo "$json_output"
```
