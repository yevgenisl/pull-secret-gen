# Pull Secret Generator

A simple Go tool that combines multiple OpenShift/Kubernetes registry pull secret JSON files into a single consolidated pull secret.

## Installation
```bash
go install github.com/yevgenisl/pull-secret-gen@latest
```

## Usage
```bash
pull-secret-gen -dir <directory> [-output <output-file>]
```

### Flags

- `-dir`: (Required) Directory containing JSON pull secret files
- `-output`: (Optional) Output file path (default: "combined-pull-secret.json")

### Example

```bash
pull-secret-gen -dir ./pull-secrets -output combined.json
```

## Input Format

Each JSON file in the input directory should contain pull secrets in the standard format:

```json
{
  "auths": {
    "registry.example.com": {
      "auth": "base64-encoded-credentials",
      ...
    }
  }
}
```

## Output

The tool will generate a single JSON file combining all registry authentications from the input files. If multiple files contain credentials for the same registry, the last processed file's credentials will take precedence.

## License

MIT

