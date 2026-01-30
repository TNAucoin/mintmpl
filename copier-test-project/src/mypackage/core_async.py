"""Async core functionality for awesome_tool."""

import asyncio
from typing import Any, Dict, List, Optional
import logging
from awesome_tool.config import Config
from awesome_tool.utils import validate_input


logger = logging.getLogger("awesome_tool")


class AwesomeServiceAsync:
    """Async version of AwesomeService for mypackage.

    This class provides async/await functionality for processing data
    asynchronously in high-performance scenarios.

    Attributes:
        config: Configuration instance for the service
        _cache: Internal cache for processed results
    """

    def __init__(self, config: Optional[Config] = None):
        """Initialize AwesomeServiceAsync.

        Args:
            config: Optional configuration instance
        """
        self.config = config or Config()
        self._cache: Dict[str, Any] = {}

        logger.info("AwesomeServiceAsync initialized for mypackage")

    async def process_data(self, data: str) -> Dict[str, Any]:
        """Asynchronously process input data.

        Args:
            data: The input data to process

        Returns:
            Dict containing the processed results

        Raises:
            ValueError: If input data is invalid
        """
        if not validate_input(data):
            raise ValueError("Invalid input data provided to AwesomeServiceAsync")

        cache_key = f"awesome_tool_async_{data}"
        if cache_key in self._cache:
            return self._cache[cache_key]

        # Simulate async processing
        await asyncio.sleep(0.1)

        logger.info("AwesomeServiceAsync processing: %s", data)

        processed = {
            "status": "success",
            "data": data.upper(),
            "metadata": {
                "service": "awesome_tool",
                "mode": "async",
                "database": self.config.database_name,
            }
        }

        self._cache[cache_key] = processed
        return processed

    async def batch_process(self, items: List[str]) -> List[Dict[str, Any]]:
        """Process multiple items concurrently.

        Args:
            items: List of items to process

        Returns:
            List of processed results
        """
        logger.info("AwesomeServiceAsync batch processing %d items", len(items))

        tasks = [self.process_data(item) for item in items]
        results = await asyncio.gather(*tasks, return_exceptions=True)

        # Handle exceptions
        processed_results = []
        for result in results:
            if isinstance(result, Exception):
                processed_results.append({
                    "status": "error",
                    "data": None,
                    "error": str(result)
                })
            else:
                processed_results.append(result)

        return processed_results
