from functools import lru_cache

from pydantic import BaseSettings


class Settings(BaseSettings):
    """Settings class holds Twitter test configuration. Reads from environment variables."""

    # TWITTER API SETTINGS
    API_PORT: str = "3000"

    # MONGODB SETTINGS
    DB_NAME: str = "twitter"
    DB_URI: str
    DB_COLLECTION_NAME: str = "messages"


@lru_cache()
def get_settings() -> Settings:
    """Gets the application settings."""
    return Settings()
