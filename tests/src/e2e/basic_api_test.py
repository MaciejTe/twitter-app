import datetime

import pytest
import requests
from bson.objectid import ObjectId

from src.core.config import get_settings

settings = get_settings()
URL = f"http://localhost:{settings.API_PORT}/messages"


def test_get_no_messages():
    """Try to get messages on empty collection. """
    response = requests.get(URL)
    assert response.status_code == 200
    assert response.json() == {"status": True, "count": 0}

def test_insert_message(db_client):
    "Create sample message and check the result in database."
    payload = {
        "text": "sample message",
        "tags": ["tag1", "tag2"]
    }
    response = requests.post(URL, data=payload)
    assert response.status_code == 201
    assert response.json()["count"] == 1
    assert response.json()["result"]["text"] == "sample message"
    
    doc_id = response.json()["result"]["id"]
    count = db_client[settings.DB_COLLECTION_NAME].count_documents({"_id": ObjectId(doc_id)})
    assert count == 1

def test_get_messages():
    """Get messages using tag query parameter. """
    query_params = {
        "tags": "tag1,tag2"
    }
    response = requests.get(URL, params=query_params)
    assert response.status_code == 200
    assert response.json() == {"status": True, "count": 1}


def test_get_messages_query_param_to():
    """Get messages using `to` query parameter. """
    end = datetime.datetime.utcnow().isoformat("T") + "Z"
    query_params = {
        "to": end
    }
    response = requests.get(URL, params=query_params)
    assert response.status_code == 200
    assert response.json() == {"status": True, "count": 1}

def test_get_messages_query_param_from():
    """Get messages using `from` query parameter. """
    start = datetime.datetime.utcnow().isoformat("T") + "Z"
    query_params = {
        "from": start
    }
    response = requests.get(URL, params=query_params)
    assert response.status_code == 200
    assert response.json() == {"status": True, "count": 0}

def test_get_messages_improper_query_param_from():
    """Get messages using improper `from` query parameter. """
    # end = datetime.datetime.utcnow().isoformat("T") + "Z"
    query_params = {
        "from": "2912-12-54-a98wj"
    }
    response = requests.get(URL, params=query_params)
    assert response.status_code == 422
    assert response.json() == {'error': 'failed to convert start filtering date to RFC3339 time format'}

def test_get_messages_improper_query_param_to():
    """Get messages using improper `from` query parameter. """
    # end = datetime.datetime.utcnow().isoformat("T") + "Z"
    query_params = {
        "to": "2912-12-54-a98wj"
    }
    response = requests.get(URL, params=query_params)
    assert response.status_code == 422
    assert response.json() == {'error': 'failed to convert stop filtering date to RFC3339 time format'}
