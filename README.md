# nitter-rss-proxy

[Twitter stopped providing RSS feeds] in 2011. [Nitter], an alternate Twitter
frontend, [serves ATOM feeds], and there are various [public Nitter instances].
However, the public instances are constantly going up and down or getting
rate-limited by Twitter, in my experience.

This repository provides a small [Go] HTTP server that accepts requests for RSS
feeds and iterates through multiple Nitter instances until it finds one that's
working. It also attempts to rewrite GUIDs and links within the feed to refer to
`https://twitter.com` rather than the particular Nitter instance that generated
the feed, so that feed readers won't show duplicate items from different
instances.

[Twitter stopped providing RSS feeds]: https://sociable.co/social-media/twitter-removes-all-search-rss-links-from-its-site-now-users-must-resort-to-hacks-to-get-feeds/
[Nitter]: https://github.com/zedeus/nitter
[serves ATOM feeds]: https://github.com/zedeus/nitter/issues/5
[public Nitter instances]: https://github.com/zedeus/nitter/wiki/Instances
[Go]: https://golang.org/

## Usage

use config in [`./config/config.yaml`](./config/config.yaml)
```
proxy:
  addr: 0.0.0.0:8080
  instances:
    - https://.../
```

The server passes `GET` request paths to the Nitter instance:

*   `http://localhost:8080/USPS`
*   `http://localhost:8080/USPS,NWS`
*   `http://localhost:8080/USPS/media`
*   `http://localhost:8080/USPS/search`
*   `http://localhost:8080/USPS/with_replies`

### Docker
```
docker run -p 8080:8080 -v `pwd`/data:/data dezhishen/nitter-rss-proxy
```
