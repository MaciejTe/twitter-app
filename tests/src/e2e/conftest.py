import pytest
from pymongo import MongoClient

from src.core.config import get_settings

settings = get_settings()


@pytest.fixture()
def db_client() :
    """Create database client for testing purposes."""
    mongo_client = MongoClient(settings.DB_URI)
    db_client = mongo_client[settings.DB_NAME]
    return db_client

@pytest.fixture(autouse=True, scope="session")
def reset_database():
    """Reset database before tests."""
    client = MongoClient(settings.DB_URI)
    client.drop_database(settings.DB_NAME)
