# spotctl

A command-line interface for managing and interacting with Rackspace Spot resources through their public API.

## Features

- **Resource Management**: List and manage organizations, regions, servercla# Output in JSON format
  spotctl cloudspaces list <namespace> --output json

## Examples

- **Configuration Management**: Easy setup and management of API credentials
- **Multiple Output Formats**: Support for JSON, YAML, and table output formats
- **Flexible Display Options**: Wide and detail views with customizable table columns
- **Cross-Platform**: Available for Linux, macOS, and Windows

## Available Commands

- **Organizations management**: List organizations in your account
- **Regions management**: List available Rackspace regions
- **Serverclasses management**: List and get details of available server classes
- **Cloudspaces management**: List cloudspaces in a namespace
- **Configuration management**: Setup and manage API credentials
- **Multiple output formats**: Table, JSON, and YAML output formats

## Installation

### From Source

```bash
git clone https://github.com/georgetaylor/spotctl.git
cd spotctl
make build
# Binary will be available in bin/spotctl
```

### Using Go Install

```bash
go install github.com/georgetaylor/spotctl@latest
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
spotctl --refresh-token your-refresh-token --region uk-lon-1 instances list
```

### Using Environment Variables

```bash
export SPOTCTL_REFRESH_TOKEN=your-refresh-token
export SPOTCTL_REGION=uk-lon-1
spotctl instances list
```

### Using Configuration File

Create a configuration file at `~/.spot/config.yaml`:

```yaml
refresh-token: your-refresh-token
region: uk-lon-1
base-url: https://spot.rackspace.com/apis/ngpc.rxt.io/v1
debug: false
timeout: 30
```

### Interactive Configuration

Initialize your configuration interactively:

```bash
spotctl config init
```

Or use the config command to set individual values:

```bash
spotctl config set refresh-token your-refresh-token
spotctl config set region uk-lon-1
```

View current configuration:

```bash
spotctl config show
```

## Output Formatting and Paging

The CLI supports multiple output formats and includes intelligent paging for long output, similar to AWS CLI.

### Output Formats

- **Table** (default): Human-readable tabular output
- **JSON**: Machine-readable JSON format
- **YAML**: YAML format for configuration files

```bash
# Table format (default)
spotctl regions list

# JSON format
spotctl regions list --output json

# YAML format
spotctl regions list --output yaml
```

### Table Display Options

```bash
# Basic table (default)
spotctl regions list

# Detailed table with additional columns
spotctl regions list --details

# Wide table with all available columns
spotctl regions list --wide
```

### Automatic Paging

For long output, the CLI automatically uses a pager (like `less` or `more`) when:

- Output is longer than the terminal height
- Output is going to a terminal (not piped to a file)

#### Pager Configuration

```bash
# Disable pager with flag
spotctl regions list --no-pager

# Disable pager with environment variable
export SPOTCTL_NO_PAGER=true
spotctl regions list

# Configure custom pager
export PAGER="less -R"  # Color-preserving pager
export PAGER="more"     # Simple pager
export PAGER="cat"      # No paging (direct output)
```

The pager respects your `$PAGER` environment variable. If the configured pager is not available, the CLI will display a warning and output directly to the terminal. The CLI does not fall back to other pagers automatically.

## Usage

### Region Management

```bash
# List all available regions
spotctl regions list

# List regions with detailed information
spotctl regions list --details

# List regions with JSON output
spotctl regions list --output json
```

### Organization Management

```bash
# List all organizations
spotctl organizations list

# List organizations with detailed information
spotctl organizations list --details

# List organizations with JSON output
spotctl organizations list --output json
```

### ServerClass Management

```bash
# List all serverclasses
spotctl serverclasses list

# List serverclasses with detailed information
spotctl serverclasses list --details

# Get details of a specific serverclass
spotctl serverclasses get <serverclass-name>

# Output in JSON format
spotctl serverclasses list --output json
```

### CloudSpace Management

```bash
# List cloudspaces in a namespace
spotctl cloudspaces list <namespace>

# List cloudspaces with detailed information
spotctl cloudspaces list <namespace> --details

# List cloudspaces with wide view (all columns)
spotctl cloudspaces list <namespace> --wide

# Output in JSON format
spotctl cloudspaces list <namespace> --output json
```

### Planned Features (API Documentation Pending)

The following commands are available but display placeholder messages pending API documentation:

```bash
# Instance management (not yet implemented)
spotctl instances list
spotctl instances create
spotctl instances delete <instance-id>

# Pricing information (not yet implemented)
spotctl pricing current
spotctl pricing history
```

### Global Options

- `--refresh-token`: Your Rackspace Spot refresh token
- `--region`: Rackspace region
- `--config`: Path to config file
- `--debug`: Enable debug output
- `--output, -o`: Output format (json, yaml, table)
- `--details`: Show detailed information with additional columns
- `--wide`: Show wide view with all available columns
- `--no-pager`: Disable automatic paging for long output

## Examples

Here are some examples of the currently working commands:

```bash
# Basic resource listing
spotctl regions list
spotctl organizations list
spotctl serverclasses list

# Detailed views with additional information
spotctl regions list --details
spotctl serverclasses list --wide

# Different output formats
spotctl organizations list --output json
spotctl regions list --output yaml

# CloudSpace management (requires namespace)
spotctl cloudspaces list my-namespace
spotctl cloudspaces list my-namespace --wide --output json

# Get specific serverclass details
spotctl serverclasses get standard-2-4-80
```

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
│   ├── regions.go         # Region management commands
│   ├── organizations.go   # Organization management commands
│   ├── serverclasses.go   # ServerClass management commands
│   └── cloudspaces.go     # CloudSpace management commands
├── pkg/
│   ├── client/            # API client and types
│   │   ├── client.go      # Main API client
│   │   └── types.go       # API response types
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
