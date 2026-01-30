"""Basic usage example for awesome_tool."""

from awesome_tool import AwesomeService, Config, format_output


def main():
    """Demonstrate basic usage of awesome_tool."""
    print("=" * 60)
    print("awesome_tool - Basic Usage Example")
    print("=" * 60)
    print()

    # Create a custom configuration
    config = Config(
        database_name="awesome_db",
        debug_mode=True,
        service_name="awesome_tool"
    )

    # Initialize MyService
    service = AwesomeService(config)
    print(f"Initialized AwesomeService with database: {config.database_name}")
    print()

    # Process single item
    print("Processing single item:")
    result = service.process_data("hello world")
    print(format_output(result, pretty=True))
    print()

    # Process batch
    print("Processing batch:")
    items = ["first item", "second item", "third item"]
    results = service.batch_process(items)

    for i, result in enumerate(results, 1):
        print(f"\nItem {i}:")
        print(format_output(result, pretty=False))

    # Get statistics
    print("\nService Statistics:")
    stats = service.get_stats()
    for key, value in stats.items():
        print(f"  {key}: {value}")


if __name__ == "__main__":
    main()
