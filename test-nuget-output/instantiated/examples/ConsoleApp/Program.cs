using System;
using System.Threading.Tasks;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Logging;
using MyLibrary;
using MyLibrary.Models;
using MyLibrary.Utilities;

namespace ConsoleApp;

/// <summary>
/// Example console application demonstrating MyLibrary usage.
/// </summary>
internal static class Program
{
    private static async Task Main(string[] args)
    {
        Console.WriteLine("MyLibrary Example Application");
        Console.WriteLine("==============================\n");

        // Setup dependency injection and logging
        var services = new ServiceCollection();
        services.AddLogging(builder =>
        {
            builder.AddConsole();
            builder.SetMinimumLevel(LogLevel.Information);
        });

        var serviceProvider = services.BuildServiceProvider();
        var logger = serviceProvider.GetRequiredService<ILogger<ServiceClient>>();

        // Create and validate configuration
        var config = new Configuration
        {
            ApiKey = "demo-api-key-12345",
            BaseUrl = "https://api.example.com",
            Timeout = TimeSpan.FromSeconds(30),
            MaxRetries = 3,
            UseCompression = true
        };

        Console.WriteLine("Configuration:");
        Console.WriteLine($"  Base URL: {config.BaseUrl}");
        Console.WriteLine($"  Timeout: {config.Timeout.TotalSeconds}s");
        Console.WriteLine($"  Max Retries: {config.MaxRetries}\n");

        // Validate configuration
        try
        {
            config.Validate();
            Console.WriteLine("✓ Configuration validated successfully\n");
        }
        catch (InvalidOperationException ex)
        {
            Console.WriteLine($"✗ Configuration error: {ex.Message}");
            return;
        }

        // Use the service client
        using (var client = new ServiceClient(config, logger))
        {
            Console.WriteLine("Service Client Examples");
            Console.WriteLine("----------------------\n");

            // Validate connection
            Console.WriteLine("1. Validating connection...");
            var isValid = await client.ValidateConnectionAsync();
            Console.WriteLine($"   Connection valid: {isValid}\n");

            // Get data
            Console.WriteLine("2. Fetching data...");
            var data = await client.GetDataAsync("resource-123");
            Console.WriteLine($"   Retrieved: {data}\n");

            // Post data
            Console.WriteLine("3. Posting data...");
            var postResult = await client.PostDataAsync("resource-456", "sample-data");
            Console.WriteLine($"   Post successful: {postResult}\n");
        }

        // Demonstrate utility functions
        Console.WriteLine("\nString Utility Examples");
        Console.WriteLine("----------------------\n");

        var sampleText = "hello world from MyLibrary";

        Console.WriteLine($"Original: {sampleText}");
        Console.WriteLine($"Title Case: {StringHelper.ToTitleCase(sampleText)}");
        Console.WriteLine($"Camel Case: {StringHelper.ToCamelCase(sampleText)}");
        Console.WriteLine($"Pascal Case: {StringHelper.ToPascalCase(sampleText)}");
        Console.WriteLine($"Truncated (20 chars): {StringHelper.Truncate(sampleText, 20)}");

        Console.WriteLine("\n✓ All examples completed successfully!");
    }
}
