using System;

namespace MyLibrary.Models;

/// <summary>
/// Configuration settings for the service client.
/// </summary>
public sealed class Configuration
{
    /// <summary>
    /// Gets or sets the API key for authentication.
    /// </summary>
    public string ApiKey { get; set; } = string.Empty;

    /// <summary>
    /// Gets or sets the base URL for the service.
    /// </summary>
    public string BaseUrl { get; set; } = "https://api.example.com";

    /// <summary>
    /// Gets or sets the request timeout duration.
    /// </summary>
    public TimeSpan Timeout { get; set; } = TimeSpan.FromSeconds(30);

    /// <summary>
    /// Gets or sets the maximum number of retry attempts.
    /// </summary>
    public int MaxRetries { get; set; } = 3;

    /// <summary>
    /// Gets or sets a value indicating whether to use compression.
    /// </summary>
    public bool UseCompression { get; set; } = true;

    /// <summary>
    /// Gets or sets additional custom headers.
    /// </summary>
    public System.Collections.Generic.Dictionary<string, string>? CustomHeaders { get; set; }

    /// <summary>
    /// Validates the configuration settings.
    /// </summary>
    /// <exception cref="InvalidOperationException">Thrown when configuration is invalid.</exception>
    public void Validate()
    {
        if (string.IsNullOrWhiteSpace(ApiKey))
        {
            throw new InvalidOperationException("API key is required.");
        }

        if (string.IsNullOrWhiteSpace(BaseUrl))
        {
            throw new InvalidOperationException("Base URL is required.");
        }

        if (!Uri.TryCreate(BaseUrl, UriKind.Absolute, out _))
        {
            throw new InvalidOperationException("Base URL must be a valid absolute URI.");
        }

        if (Timeout <= TimeSpan.Zero)
        {
            throw new InvalidOperationException("Timeout must be greater than zero.");
        }

        if (MaxRetries < 0)
        {
            throw new InvalidOperationException("Max retries cannot be negative.");
        }
    }
}
