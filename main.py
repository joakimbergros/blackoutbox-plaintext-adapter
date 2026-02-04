import os
import sys
import requests

API_BASE_URL = os.getenv("API_BASE_URL")  # e.g. http://localhost:3000
SYSTEM_ID = os.getenv("SYSTEM_ID")  # e.g. system-123

if not API_BASE_URL or not SYSTEM_ID:
    print("API_BASE_URL and SYSTEM_ID must be set", file=sys.stderr)
    sys.exit(1)

text_document = """Hello,
This is a text document coming from an external system.
It will be stored as a .txt file.
"""

documents = [
    {
        "file_id": "DOC-001",
        "document": text_document,
        "tags": ["example", "text"],
    }
]

url = f"{API_BASE_URL}/systems/{SYSTEM_ID}/sync"

try:
    response = requests.post(
        url,
        json=documents,  # use json= to auto-set header
        timeout=10,  # seconds
    )
except requests.RequestException as e:
    print(f"HTTP request failed: {e}", file=sys.stderr)
    sys.exit(1)

if response.status_code != 204:
    print(
        f"sync failed: {response.status_code} {response.text}", file=sys.stderr
    )
    sys.exit(1)

print("Adapter: text document synced successfully")
