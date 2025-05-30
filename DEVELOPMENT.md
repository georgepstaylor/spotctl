# Spotctl Development Guide

## Project Structure

The project follows Go best practices with a clear separation of concerns:

```
├── cmd/                    # CLI commands using Cobra
│   ├── root.go            # Root command with global flags
│   ├── version.go         # Version information
│   ├── config.go          # Configuration management
│   ├── instances.go       # Instance management (placeholder)
│   └── pricing.go         # Pricing commands (placeholder)
├── pkg/                   # Public packages
│   ├── client/            # HTTP client for Rackspace API
│   └── config/            # Configuration management
├── internal/              # Private packages
│   └── utils/             # Utility functions
├── main.go                # Application entry point
└── Makefile              # Build automation
```

## Adding New Commands

When you receive the API documentation, follow this pattern to add new commands:

1. **Define the API models** in `pkg/client/` (e.g., `types.go`)
2. **Add API methods** to the client in `pkg/client/client.go`
3. **Create command files** in `cmd/` following the pattern of existing commands
4. **Add the command** to the root command in the `init()` function

### Example: Adding a new "servers" command

1. Create `cmd/servers.go`:

```go
package cmd

import (
    "context"
    "fmt"

    "github.com/georgetaylor/rackspace-spot-cli/internal/utils"
    "github.com/georgetaylor/rackspace-spot-cli/pkg/client"
    "github.com/georgetaylor/rackspace-spot-cli/pkg/config"
    "github.com/spf13/cobra"
)

var serversCmd = &cobra.Command{
    Use:   "servers",
    Short: "Manage servers",
    Long:  "Manage Rackspace spot servers.",
}

var serversListCmd = &cobra.Command{
    Use:   "list",
    Short: "List servers",
    Run: func(cmd *cobra.Command, args []string) {
        cfg, err := config.GetConfig()
        utils.CheckError(err)

        client := client.NewClient(cfg)

        // Call API method
        servers, err := client.ListServers(context.Background())
        utils.CheckError(err)

        // Format output
        output := utils.GetOutputFormat(cmd)
        utils.CheckError(utils.FormatOutput(servers, output))
    },
}

func init() {
    rootCmd.AddCommand(serversCmd)
    serversCmd.AddCommand(serversListCmd)
    utils.AddOutputFlag(serversListCmd)
}
```

2. Add API method to `pkg/client/client.go`:

```go
type Server struct {
    ID     string `json:"id"`
    Name   string `json:"name"`
    Status string `json:"status"`
    // ... other fields based on API
}

func (c *Client) ListServers(ctx context.Context) ([]Server, error) {
    resp, err := c.Get(ctx, "/servers")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if err := c.HandleAPIError(resp); err != nil {
        return nil, err
    }

    var servers []Server
    if err := json.NewDecoder(resp.Body).Decode(&servers); err != nil {
        return nil, fmt.Errorf("failed to decode response: %w", err)
    }

    return servers, nil
}
```

## Testing

- Write unit tests for all new functionality
- Place tests next to the code they test with `_test.go` suffix
- Use `make test` to run all tests
- Use `make test-coverage` to generate coverage reports

## Building and Distribution

- `make build` - Build for current platform
- `make build-all` - Build for all supported platforms
- `make install` - Install locally
- `make clean` - Clean build artifacts

## Code Quality

- Run `make fmt` to format code
- Run `make lint` to check for issues (requires golangci-lint)
- Run `make dev` for the full development workflow

## Configuration

The CLI supports multiple configuration methods:

1. Command-line flags (highest priority)
2. Environment variables (RACKSPACE*SPOT*\*)
3. Configuration file (~/.spot/config.yaml)
4. Defaults (lowest priority)

## Error Handling

- Use the `utils.CheckError()` function for fatal errors
- Return errors from functions and handle them at the command level
- Use the `APIError` type for API-specific errors
- Provide helpful error messages to users

## Output Formatting

- Support both JSON and table output formats
- Use `utils.AddOutputFlag()` to add the output flag to commands
- Use `utils.FormatOutput()` to format the output
- Default to table format for better UX

## Next Steps

1. **Receive API Documentation**: Once you have the API docs, you can:

   - Define proper data structures in `pkg/client/types.go`
   - Implement real API calls in the client
   - Replace placeholder commands with real functionality

2. **Add Authentication**: Implement proper authentication based on the API requirements

3. **Add More Commands**: Based on the API capabilities, add commands for:

   - Instance lifecycle management
   - Pricing and billing
   - Monitoring and logs
   - Resource management

4. **Enhance Output**: Add proper table formatting, colors, and better UX

5. **Add Validation**: Implement input validation and better error messages

6. **Documentation**: Add more examples and usage documentation
