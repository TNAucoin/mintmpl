using System;
using System.Threading;
using System.Threading.Tasks;
using Microsoft.Extensions.Logging;
using MyLibrary.Models;

namespace MyLibrary;

/// <summary>
/// Main service client for interacting with the API.
/// </summary>
public sealed class ServiceClient : IDisposable
{
    private readonly Configuration _configuration;
    private readonly ILogger<ServiceClient> _logger;
    private bool _disposed;

    /// <summary>
    /// Initializes a new instance of the <see cref="ServiceClient"/> class.
    /// </summary>
    /// <param name="configuration">The configuration settings.</param>
    /// <param name="logger">Optional logger instance.</param>
    /// <exception cref="ArgumentNullException">Thrown when configuration is null.</exception>
    public ServiceClient(Configuration configuration, ILogger<ServiceClient>? logger = null)
    {
        _configuration = configuration ?? throw new ArgumentNullException(nameof(configuration));
        _logger = logger ?? Microsoft.Extensions.Logging.Abstractions.NullLogger<ServiceClient>.Instance;

        _logger.LogInformation("ServiceClient initialized with base URL: {BaseUrl}", _configuration.BaseUrl);
    }

    /// <summary>
    /// Gets data from the service asynchronously.
    /// </summary>
    /// <param name="resourceId">The resource identifier.</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    /// <returns>A task representing the asynchronous operation with the retrieved data.</returns>
    /// <exception cref="ArgumentException">Thrown when resourceId is null or empty.</exception>
    /// <exception cref="ObjectDisposedException">Thrown when the client has been disposed.</exception>
    public async Task<string> GetDataAsync(string resourceId, CancellationToken cancellationToken = default)
    {
        ObjectDisposedException.ThrowIf(_disposed, this);

        if (string.IsNullOrWhiteSpace(resourceId))
        {
            throw new ArgumentException("Resource ID cannot be null or empty.", nameof(resourceId));
        }

        _logger.LogDebug("Fetching data for resource: {ResourceId}", resourceId);

        // Simulate async operation
        await Task.Delay(100, cancellationToken).ConfigureAwait(false);

        var result = $"Data for {resourceId} from {_configuration.BaseUrl}";

        _logger.LogInformation("Successfully retrieved data for resource: {ResourceId}", resourceId);

        return result;
    }

    /// <summary>
    /// Posts data to the service asynchronously.
    /// </summary>
    /// <param name="resourceId">The resource identifier.</param>
    /// <param name="data">The data to post.</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    /// <returns>A task representing the asynchronous operation with the result.</returns>
    /// <exception cref="ArgumentException">Thrown when parameters are invalid.</exception>
    /// <exception cref="ObjectDisposedException">Thrown when the client has been disposed.</exception>
    public async Task<bool> PostDataAsync(string resourceId, string data, CancellationToken cancellationToken = default)
    {
        ObjectDisposedException.ThrowIf(_disposed, this);

        if (string.IsNullOrWhiteSpace(resourceId))
        {
            throw new ArgumentException("Resource ID cannot be null or empty.", nameof(resourceId));
        }

        if (string.IsNullOrWhiteSpace(data))
        {
            throw new ArgumentException("Data cannot be null or empty.", nameof(data));
        }

        _logger.LogDebug("Posting data to resource: {ResourceId}", resourceId);

        // Simulate async operation
        await Task.Delay(150, cancellationToken).ConfigureAwait(false);

        _logger.LogInformation("Successfully posted data to resource: {ResourceId}", resourceId);

        return true;
    }

    /// <summary>
    /// Validates the connection to the service.
    /// </summary>
    /// <param name="cancellationToken">Cancellation token.</param>
    /// <returns>A task representing the asynchronous operation with the validation result.</returns>
    public async Task<bool> ValidateConnectionAsync(CancellationToken cancellationToken = default)
    {
        ObjectDisposedException.ThrowIf(_disposed, this);

        _logger.LogDebug("Validating connection to {BaseUrl}", _configuration.BaseUrl);

        // Simulate connection check
        await Task.Delay(50, cancellationToken).ConfigureAwait(false);

        _logger.LogInformation("Connection validation successful");

        return true;
    }

    /// <summary>
    /// Disposes the service client and releases resources.
    /// </summary>
    public void Dispose()
    {
        if (_disposed)
        {
            return;
        }

        _logger.LogDebug("Disposing ServiceClient");
        _disposed = true;
    }
}
