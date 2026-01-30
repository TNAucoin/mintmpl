# My Library

A modern .NET library for awesome functionality

## Features

- Clean, maintainable code following .NET best practices
- Comprehensive unit tests
- Full XML documentation for IntelliSense
- NuGet package ready with proper metadata
- Async/await support
- Dependency injection friendly

## Installation

```bash
dotnet add package MyLibrary
```

## Usage

```csharp
using MyLibrary;

// Create the service client
var client = new ServiceClient(new Configuration
{
    ApiKey = "your-api-key",
    BaseUrl = "https://api.example.com"
});

// Use the client
var result = await client.GetDataAsync("resource-id");
Console.WriteLine(result);
```

## Configuration

The library uses a `Configuration` object for setup:

```csharp
var config = new Configuration
{
    ApiKey = "your-api-key",
    BaseUrl = "https://api.example.com",
    Timeout = TimeSpan.FromSeconds(30),
    MaxRetries = 3
};
```

## Development

### Building

```bash
dotnet build
```

### Running Tests

```bash
dotnet test
```

### Creating NuGet Package

```bash
dotnet pack -c Release
```

## License

MIT License - see LICENSE file for details

## Author

John Smith (john.smith@example.com)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
