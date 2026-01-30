"""Core functionality for awesome_tool."""

from typing import Any, Dict, List, Optional
import logging
from awesome_tool.config import Config
from awesome_tool.utils import validate_input, format_output


# Configure logging for mypackage
logger = logging.getLogger("awesome_tool")


class AwesomeService:
    """Main service class for mypackage.

    This class provides the core functionality for processing data
    and interacting with external services.

    Attributes:
        config: Configuration instance for the service
        _cache: Internal cache for processed results

    Example:
        >>> service = AwesomeService()
        >>> result = service.process_data("example input")
        >>> print(result)
    """

    def __init__(self, config: Optional[Config] = None):
        """Initialize AwesomeService.

        Args:
            config: Optional configuration instance. If not provided,
                   default configuration will be used.
        """
        self.config = config or Config()
        self._cache: Dict[str, Any] = {}

        logger.info("AwesomeService initialized with database: %s", self.config.database_name)

    def process_data(self, data: str) -> Dict[str, Any]:
        """Process input data and return results.

        This method validates the input, processes it according to
        the service configuration, and returns formatted results.

        Args:
            data: The input data to process

        Returns:
            Dict containing the processed results with keys:
                - 'status': Processing status
                - 'data': Processed data
                - 'metadata': Additional metadata

        Raises:
            ValueError: If input data is invalid
        """
        # Validate input
        if not validate_input(data):
            raise ValueError("Invalid input data provided to AwesomeService")

        # Check cache
        cache_key = f"awesome_tool_{data}"
        if cache_key in self._cache:
            logger.debug("Returning cached result for: %s", data)
            return self._cache[cache_key]

        # Process data
        logger.info("Processing data in AwesomeService: %s", data)

        processed = {
            "status": "success",
            "data": data.upper(),
            "metadata": {
                "service": "awesome_tool",
                "version": "0.1.0",
                "database": self.config.database_name,
            }
        }

        # Cache result
        self._cache[cache_key] = processed

        return processed

    def batch_process(self, items: List[str]) -> List[Dict[str, Any]]:
        """Process multiple items in batch.

        Args:
            items: List of items to process

        Returns:
            List of processed results
        """
        logger.info("AwesomeService batch processing %d items", len(items))

        results = []
        for item in items:
            try:
                result = self.process_data(item)
                results.append(result)
            except ValueError as e:
                logger.error("Error processing item in AwesomeService: %s", e)
                results.append({
                    "status": "error",
                    "data": None,
                    "error": str(e)
                })

        return results

    def clear_cache(self) -> None:
        """Clear the internal cache."""
        logger.info("Clearing AwesomeService cache")
        self._cache.clear()

    def get_stats(self) -> Dict[str, Any]:
        """Get service statistics.

        Returns:
            Dictionary containing service statistics
        """
        return {
            "service_name": "awesome_tool",
            "cache_size": len(self._cache),
            "database": self.config.database_name,
            "debug_mode": self.config.debug_mode,
        }
