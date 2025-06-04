# 🚀 spotctl

A modern CLI for managing Rackspace Spot resources with style and simplicity.

## ✨ Features

- 🎯 **Resource Management** - Organizations, regions, server classes, and cloudspaces
- ⚙️ **Smart Configuration** - Interactive setup with multiple auth methods
- 📊 **Flexible Output** - Table, JSON, and YAML formats with intelligent paging
- 🌍 **Cross-Platform** - Linux, macOS, and Windows support

## 📦 Installation

### Quick Start

```bash
go install github.com/georgetaylor/spotctl@latest
```

### From Source

```bash
git clone https://github.com/georgetaylor/spotctl.git
cd spotctl
make build
```

## 🔧 Configuration

### 1. Get Your Token

1. Visit [Rackspace Spot Console](https://spot.rackspace.com) 🌐
2. Go to **API Access > Terraform**
3. Click **Get New Token**
4. Copy your refresh token

### 2. Configure spotctl

```bash
# Interactive setup (recommended)
spotctl config init

# Or set values directly
spotctl config set refresh-token your-token-here
spotctl config set region uk-lon-1
```

### Alternative Methods

```bash
# Environment variables
export SPOTCTL_REFRESH_TOKEN=your-token
export SPOTCTL_REGION=uk-lon-1

# Command flags
spotctl --refresh-token your-token --region uk-lon-1 regions list
```

## 🎮 Usage

### Quick Commands

```bash
# List available regions
spotctl regions list

# List your organizations
spotctl organizations list

# List server classes with details
spotctl serverclasses list --details

# List cloudspaces in a namespace
spotctl cloudspaces list my-namespace

# Get specific server class info
spotctl serverclasses get standard-2-4-80
```

### Output Formats

```bash
# Default table view
spotctl regions list

# JSON output for automation
spotctl regions list --output json

# YAML for configuration
spotctl regions list --output yaml

# Detailed view with extra columns
spotctl regions list --details
```

### Global Options

| Flag           | Description                            |
| -------------- | -------------------------------------- |
| `--output, -o` | Output format: `table`, `json`, `yaml` |
| `--details`    | Show additional columns                |
| `--no-pager`   | Disable automatic paging               |
| `--debug`      | Enable debug output                    |

## 🛠️ Development

### Prerequisites

- Go 1.24+
- Make (optional)

### Commands

```bash
make build      # Build binary
make test       # Run tests
make dev        # Format, lint, test, build
```

## 📁 Project Structure

```
spotctl/
├── cmd/          # CLI commands
├── pkg/          # Public packages
│   ├── client/   # API client
│   ├── config/   # Configuration
│   ├── output/   # Formatters
│   └── pager/    # Output paging
├── internal/     # Private utilities
└── main.go       # Entry point
```

## 🤝 Contributing

1. Fork the repo
2. Create a feature branch
3. Make your changes
4. Run `make dev`
5. Submit a PR

## 📄 License

MIT License - see [LICENSE](LICENSE) for details.
