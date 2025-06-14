# LSEQ CRDT — Collaborative Text Editing

This project explores how to use **LSEQ (List Sequence)** — a CRDT (Conflict-free Replicated Data Type) algorithm — for building collaborative text editors where multiple users can edit the same document at the same time, even when offline or with network delays.

## What is LSEQ?

LSEQ stands for **List Sequence**. It’s a CRDT designed to maintain the correct order of elements (like characters in a document) without central coordination. The key idea is that every character gets assigned a **position identifier (posID)** that ensures:

- All replicas (clients) can insert concurrently and still agree on the final order.
- The posID space has room for infinite inserts between any two positions.
- The identifiers stay compact even after many edits (logarithmic growth).

## How does LSEQ position ID work?

Every character in the document has a posID — a list of integers (e.g., `[2]`, `[2, 5]`, `[3]`). When you insert between two characters (say `[2]` and `[3]`):

1. The client sends a request specifying the left and right posIDs.
2. The backend generates a posID between them (e.g., `[2, 5]`).
3. The backend broadcasts this new posID with the character and metadata.
4. All clients apply the change consistently.

## Example: concurrent inserts

Imagine two clients are editing the document `"hello"` at the same time:

- Client A inserts `'X'` between `'e'` and `'l'`.
- Client B inserts `'Y'` at the same spot.

Each sends:

```json
{
  "op": "insert",
  "char": "X",
  "left_pos_id": [2],
  "right_pos_id": [3],
  "site_id": "clientA",
  "seq": 1
}
````

```json
{
  "op": "insert",
  "char": "Y",
  "left_pos_id": [2],
  "right_pos_id": [3],
  "site_id": "clientB",
  "seq": 1
}
```

The backend generates posIDs:

````
{
    "op": "insert",
    "char": "X",
    "pos_id": [2, 5],
    "site_id": "clientA",
    "seq": 1
}

{
    "op": "insert",
    "char": "Y",
    "pos_id": [2, 3],
    "site_id": "clientB",
    "seq": 1
}
````

The backend generates posIDs using LSEQ logic between given left/right posIDs.
It broadcasts these operations to all clients — no further conflict resolution needed.

#### How collisions are handled
If somehow the backend generates the same posID (very unlikely because it controls generation), or if tie-breaking is needed:

- Site ID (client ID): lexicographic order is used.
- Sequence number: breaks ties further.

This ensures all clients apply operations consistently.
