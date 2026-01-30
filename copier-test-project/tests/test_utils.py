"""Tests for awesome_tool utility functions."""

import pytest
from awesome_tool.utils import (
    validate_input,
    format_output,
    sanitize_string,
    create_cache_key,
)


class TestValidateInput:
    """Test suite for validate_input function."""

    def test_valid_input(self):
        """Test validation with valid input."""
        assert validate_input("valid data") is True
        assert validate_input("test123") is True
        assert validate_input("ABC XYZ 789") is True

    def test_invalid_input_empty(self):
        """Test validation with empty input."""
        assert validate_input("") is False

    def test_invalid_input_special_chars(self):
        """Test validation with special characters."""
        assert validate_input("invalid@#$") is False
        assert validate_input("test!data") is False

    def test_invalid_input_too_long(self):
        """Test validation with input that's too long."""
        long_input = "a" * 10001
        assert validate_input(long_input) is False


class TestFormatOutput:
    """Test suite for format_output function."""

    def test_format_output_simple(self):
        """Test simple output formatting."""
        result = {
            "status": "success",
            "data": "test data"
        }
        output = format_output(result, pretty=False)
        assert "success" in output
        assert "test data" in output

    def test_format_output_pretty(self):
        """Test pretty output formatting."""
        result = {
            "status": "success",
            "data": "test data",
            "metadata": {
                "service": "awesome_tool",
                "version": "0.1.0"
            }
        }
        output = format_output(result, pretty=True)
        assert "AwesomeService Result" in output
        assert "success" in output
        assert "awesome_tool" in output


class TestSanitizeString:
    """Test suite for sanitize_string function."""

    def test_sanitize_removes_special_chars(self):
        """Test that special characters are removed."""
        result = sanitize_string("test@#$data")
        assert result == "testdata"

    def test_sanitize_normalizes_whitespace(self):
        """Test that whitespace is normalized."""
        result = sanitize_string("test   multiple   spaces")
        assert result == "test multiple spaces"

    def test_sanitize_preserves_alphanumeric(self):
        """Test that alphanumeric characters are preserved."""
        result = sanitize_string("test123 data456")
        assert result == "test123 data456"


class TestCreateCacheKey:
    """Test suite for create_cache_key function."""

    def test_create_cache_key_basic(self):
        """Test basic cache key creation."""
        key = create_cache_key("awesome_tool", "test data")
        assert key.startswith("awesome_tool_")
        assert "test data" in key

    def test_create_cache_key_truncates_long_data(self):
        """Test that long data is truncated."""
        long_data = "a" * 100
        key = create_cache_key("awesome_tool", long_data)
        assert len(key) <= 50 + len("awesome_tool_")
