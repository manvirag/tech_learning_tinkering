### Functional Requirement
1. Document collaboration: Multiple users should be able to edit a document simultaneously ( assume it is text based only).
2. Conflict resolution: The system should push the edits done by one user to all the other collaborators. The system should also resolve conflicts between users if they’re editing the same portion of the document.

### Non-functional Requirement
1. Concurrency: A lot of people are working on the same document 
2. Latency: Different users can be connected to collaborate on the same document. Maintaining low latency is challenging for users connected from different regions.

### Pre-requisites:

#### Operational Transformation 
#### Conflict-free Replicated Data Type (CRDT)


### High level design

#### References:

1. https://medium.com/@sureshpodeti/system-design-google-docs-93e12133a979