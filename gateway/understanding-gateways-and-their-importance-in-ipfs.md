---
description: >-
  This article aims to explain the purpose of gateways, why they are needed, and
  how they facilitate the retrieval of content from IPFS.
---

# Understanding Gateways and their Importance in IPFS

## The Purpose of Gateways:

Gateways serve as portals or intermediaries that enable users to access content pinned on IPFS through their web browsers or other applications that do not natively support IPFS. Since most applications do not yet have built-in IPFS support, gateways bridge the gap and allow users to retrieve IPFS content using familiar HTTP-based protocols.

## Retrieving Content from IPFS:

When you pin your content to the W3IPFS pinning service, it becomes available on the public IPFS network. This means that anyone running an IPFS node can access the content through their own IPFS node. However, utilizing gateways provides a more convenient and accessible way to retrieve IPFS content.

## Gateways and their Structure:

Gateways typically follow a specific structure when requesting data from the IPFS network. The format commonly used is:

```
https://gateway-url/ipfs/<CID>
```

Here, the CID represents the content identifier or hash of the specific content you want to retrieve from IPFS.

## Importance of Gateways:

Gateways act as translators, allowing IPFS content to be understood and served appropriately on modern platforms that are not IPFS-compatible. By leveraging the HTTP protocol and providing a web URL, gateways enable users to request IPFS content using their regular web browsers or applications.

{% hint style="info" %}
Gateways are the portal to access content pinned on IPFS on modern-day applications.
{% endhint %}

## W3IPFS Pinning Service:

When you pin your content to W3IPFS, it becomes accessible on the public IPFS network. To retrieve content pinned on W3IPFS, you can utilize an IPFS gateway by constructing the URL in the format mentioned above, replacing the `<gateway-url>` with the appropriate W3IPFS gateway URL.

## Conclusion:

Gateways are essential components of the IPFS ecosystem, enabling users to access content pinned on IPFS through modern-day applications. By leveraging HTTP-based protocols, gateways bridge the compatibility gap between IPFS and applications that do not natively support it.
