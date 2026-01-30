"""Utility functions for awesome_tool."""

from typing import Any, Dict
import re


def validate_input(data: str) -> bool:
    """Validate input data for mypackage.

    This function checks if the input data meets the requirements
    for processing by AwesomeService.

    Args:
        data: The input string to validate

    Returns:
        bool: True if input is valid, False otherwise

    Example:
        >>> validate_input("valid data")
        True
        >>> validate_input("")
        False
    """
    if not data or not isinstance(data, str):
        return False

    if len(data) > 10000:
        return False

    # mypackage requires alphanumeric or space characters
    if not re.match(r'^[a-zA-Z0-9\s]+$', data):
        return False

    return True


def format_output(result: Dict[str, Any], pretty: bool = False) -> str:
    """Format output from AwesomeService.

    Args:
        result: The result dictionary from AwesomeService
        pretty: Whether to format output in a pretty format

    Returns:
        Formatted string representation of the result
    """
    if pretty:
        lines = [
            "=" * 50,
            "AwesomeService Result",
            "=" * 50,
            f"Status: {result.get('status', 'unknown')}",
            f"Data: {result.get('data', 'N/A')}",
        ]

        if 'metadata' in result:
            lines.append("Metadata:")
            for key, value in result['metadata'].items():
                lines.append(f"  {key}: {value}")

        lines.append("=" * 50)
        return "\n".join(lines)
    else:
        return f"AwesomeService result: {result.get('status')} - {result.get('data')}"


def sanitize_string(text: str) -> str:
    """Sanitize a string for use in awesome_tool.

    Args:
        text: The text to sanitize

    Returns:
        Sanitized string safe for processing
    """
    # Remove any non-alphanumeric characters except spaces
    sanitized = re.sub(r'[^a-zA-Z0-9\s]', '', text)

    # Normalize whitespace
    sanitized = ' '.join(sanitized.split())

    return sanitized


def create_cache_key(prefix: str, data: str) -> str:
    """Create a cache key for awesome_tool operations.

    Args:
        prefix: The prefix for the cache key (e.g., "awesome_tool")
        data: The data to include in the key

    Returns:
        A formatted cache key string
    """
    sanitized_data = sanitize_string(data)
    return f"{prefix}_{sanitized_data[:50]}"
