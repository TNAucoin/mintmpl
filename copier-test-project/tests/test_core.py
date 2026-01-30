"""Tests for awesome_tool core functionality."""

import pytest
from awesome_tool import AwesomeService
from awesome_tool.config import Config


class TestAwesomeService:
    """Test suite for AwesomeService class."""

    def test_initialization(self):
        """Test AwesomeService initialization."""
        service = AwesomeService()
        assert service is not None
        assert isinstance(service.config, Config)

    def test_initialization_with_config(self):
        """Test AwesomeService initialization with custom config."""
        config = Config(database_name="test_db")
        service = AwesomeService(config)
        assert service.config.database_name == "test_db"

    def test_process_data(self):
        """Test basic data processing in AwesomeService."""
        service = AwesomeService()
        result = service.process_data("test input")

        assert result["status"] == "success"
        assert result["data"] == "TEST INPUT"
        assert result["metadata"]["service"] == "awesome_tool"
        assert result["metadata"]["database"] == "awesome_db"

    def test_process_data_invalid_input(self):
        """Test AwesomeService with invalid input."""
        service = AwesomeService()

        with pytest.raises(ValueError):
            service.process_data("invalid@#$%input")

    def test_process_data_caching(self):
        """Test that AwesomeService caches results."""
        service = AwesomeService()

        result1 = service.process_data("cached data")
        result2 = service.process_data("cached data")

        # Results should be identical (from cache)
        assert result1 is result2

    def test_batch_process(self):
        """Test batch processing in AwesomeService."""
        service = AwesomeService()
        items = ["item1", "item2", "item3"]

        results = service.batch_process(items)

        assert len(results) == 3
        assert all(r["status"] == "success" for r in results)
        assert results[0]["data"] == "ITEM1"
        assert results[1]["data"] == "ITEM2"
        assert results[2]["data"] == "ITEM3"

    def test_clear_cache(self):
        """Test cache clearing in AwesomeService."""
        service = AwesomeService()

        # Process some data to populate cache
        service.process_data("test")
        assert len(service._cache) > 0

        # Clear cache
        service.clear_cache()
        assert len(service._cache) == 0

    def test_get_stats(self):
        """Test getting service statistics."""
        service = AwesomeService()
        stats = service.get_stats()

        assert stats["service_name"] == "awesome_tool"
        assert stats["database"] == "awesome_db"
        assert "cache_size" in stats
