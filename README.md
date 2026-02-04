# Blackoutbox – Example Text Adapter

This repository is an **example (exemplar) adapter implementation** for the
**Blackoutbox project**:

👉 [https://github.com/joakimbergros/blackoutbox](https://github.com/joakimbergros/blackoutbox)

It demonstrates the **simplest possible way** for an external system to send
**text documents and metadata** to a Blackoutbox deployment.

---

## Purpose of this repository

This repository exists as a **reference implementation**, not as a production
component.

It answers the question:

> “How should an external system talk to Blackoutbox?”

By keeping the adapter:

* small
* explicit
* easy to read

it becomes a **learning and starting point** for building real adapters.

---

## What this adapter represents

* An **external system** that produces documents as **plain text**
* A **translator** that reshapes that data into Blackoutbox’s API contract
* A **client** that sends the data over HTTP

It intentionally does **not**:

* share code with Blackoutbox
* access the Blackoutbox database
* write to the Blackoutbox filesystem
* implement authentication or retries

Those concerns belong to real adapters, not to this example.

---

## How it fits into the Blackoutbox architecture

```
External system
    ↓
Example adapter (this repo)
    ↓  HTTP
Blackoutbox API
    ↓
Stored as .txt + metadata
```

This adapter communicates with Blackoutbox **only through its public API**.

---

## API contract used

The adapter sends data to the Blackoutbox endpoint:

```
POST /systems/{system_id}/sync
```

With a JSON payload:

```json
[
  {
    "file_id": "DOC-001",
    "document": "This is the document content as plain text.",
    "tags": ["example", "text"]
  }
]
```

* `document` is treated as a **text (`.txt`) document**
* Storage details are handled entirely by Blackoutbox

---

## Why this adapter uses plain text

This example uses **plain text documents** to keep things simple:

* no binary data
* no multipart uploads
* easy to inspect
* easy to reproduce
* easy to port to other languages

The same pattern can be extended to other formats in real adapters.

---

## Configuration

The adapter uses environment variables:

| Variable       | Description                                |
| -------------- | ------------------------------------------ |
| `API_BASE_URL` | Base URL of a running Blackoutbox instance |
| `SYSTEM_ID`    | System identifier used by Blackoutbox      |

Example:

```bash
export API_BASE_URL=http://localhost:3000
export SYSTEM_ID=example-system
```

---

## Running the example adapter

```bash
pip install requests
python adapter.py
```

If successful, the adapter sends a sample text document and exits.

---

## Who this repository is for

This repository is intended for:

* developers integrating external systems with Blackoutbox
* contributors who want to understand the adapter boundary
* anyone building their own adapter in another language

It is **not** intended to be deployed as-is.

---

## Design principles demonstrated

* **API as the boundary**
* **Loose coupling**
* **Single responsibility**
* **Replaceability**
* **Clarity over cleverness**

---

## Summary

> This repository is an **example adapter** for the Blackoutbox project
> showing the simplest possible integration using plain text documents.

Use it as a reference, copy what you need, and build your own adapter on top.

---
