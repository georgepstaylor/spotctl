# Rackspace Spot CLI

A command-line interface for managing and interacting with Rackspace Spot resources through their public API.

## Features

- **Instance Management**: Create, list, delete, and monitor spot instances
- **Pricing Information**: Get current and historical spot pricing data
- **Configuration Management**: Easy setup and management of API credentials
- **Multiple Output Formats**: Support for JSON and table output formats
- **Cross-Platform**: Available for Linux, macOS, and Windows

## Installation

### From Source

```bash
git clone https://github.com/georgetaylor/rackspace-spot-cli.git
cd rackspace-spot-cli
make build
# Binary will be available in bin/rackspace-spot
```

### Using Go Install

```bash
go install github.com/georgetaylor/rackspace-spot-cli@latest
```

## Configuration

Before using the CLI, you need to configure your Rackspace Spot refresh token.

### Getting Your Refresh Token

1. Log in to the [Rackspace Spot Console](https://spot.rackspace.com)
2. Navigate to **API Access > Terraform** in the sidebar
3. Click **Get New Token** to generate a refresh token
4. Copy the generated refresh token for use with the CLI

### Using Command Line Flags

```bash
rackspace-spot --refresh-token your-refresh-token --region us-east-1 instances list
```

### Using Environment Variables

```bash
export RACKSPACE_SPOT_REFRESH_TOKEN=your-refresh-token
export RACKSPACE_SPOT_REGION=us-east-1
rackspace-spot instances list
```

### Using Configuration File

Create a configuration file at `~/.rackspace-spot.yaml`:

```yaml
refresh-token: your-refresh-token
region: us-east-1
base-url: https://spot.rackspace.com/apis/ngpc.rxt.io/v1
debug: false
timeout: 30
```

### Interactive Configuration

Initialize your configuration interactively:

```bash
rackspace-spot config init
```

Or use the config command to set individual values:

```bash
rackspace-spot config set refresh-token your-refresh-token
rackspace-spot config set region us-east-1
```

View current configuration:

```bash
rackspace-spot config show
```

## Usage

### Region Management

```bash
# List all available regions
rackspace-spot regions list

# List regions with detailed information
rackspace-spot regions list --details

# List regions with JSON output
rackspace-spot regions list --output json
```

### Instance Management

```bash
# List all instances (implementation pending API docs)
rackspace-spot instances list

# List instances with JSON output
rackspace-spot instances list --output json

# Create a new instance (implementation pending API docs)
rackspace-spot instances create

# Delete an instance
rackspace-spot instances delete instance-id
```

### Pricing Information

```bash
# Get current pricing
rackspace-spot pricing current

# Get pricing for specific region
rackspace-spot pricing current --region us-west-2

# Get pricing history
rackspace-spot pricing history --start-date 2024-01-01 --end-date 2024-01-31
```

### Global Options

- `--refresh-token`: Your Rackspace Spot refresh token
- `--region`: Rackspace region
- `--config`: Path to config file
- `--debug`: Enable debug output
- `--output, -o`: Output format (json, table)

## Development

### Prerequisites

- Go 1.21 or later
- Make (optional, for using Makefile)

### Building

```bash
# Build for current platform
make build

# Build for all platforms
make build-all

# Install locally
make install
```

### Testing

```bash
# Run tests
make test

# Run tests with coverage
make test-coverage
```

### Development Workflow

```bash
# Format, lint, test, and build
make dev
```

## Project Structure

```
.
├── cmd/                    # Command definitions
│   ├── root.go            # Root command and global flags
│   ├── version.go         # Version command
│   ├── config.go          # Configuration management
│   ├── instances.go       # Instance management commands
│   └── pricing.go         # Pricing commands
├── pkg/
│   ├── client/            # API client
│   └── config/            # Configuration management
├── internal/
│   └── utils/             # Utility functions
├── main.go                # Application entry point
├── Makefile              # Build automation
└── README.md             # This file
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Run `make dev` to ensure code quality
6. Submit a pull request

## License

See [LICENSE](LICENSE) file for details.
