# 📌 What does "pinning" mean?

In the context of IPFS, "pinning" refers to the act of marking a file or a set of files to be permanently stored and retained within a node or a network. When you pin a file in IPFS, you ensure that the file is always available and won't be automatically removed due to garbage collection or other mechanisms that manage storage resources.

Pinning is an essential concept in IPFS because, by default, the network operates on a content-addressable system where files are stored based on demand. When a file is requested by a user, it gets temporarily cached on the nodes that served it. However, if there is no ongoing demand for a file, it may eventually get evicted from the cache to make room for more frequently accessed content.

By pinning a file, you tell the IPFS network to retain that file in your local node's storage or in other nodes participating in the network. Pinning can be done at various levels, from individual files to entire directories or even large datasets. Pinning ensures the persistence of files and allows them to be accessible even when there is no immediate demand for them.

Pinning is particularly useful when you want to ensure that specific content remains available in IPFS for as long as you need it, even if the file is not frequently accessed by others. It provides control over the lifecycle and availability of data in the IPFS network.
