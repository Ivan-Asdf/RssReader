# RssReader

### Build
run `go build` in project directory

### Usage
<pre>RssReader [-t] | <ins>url</ins>...</pre>.
If given flag `-t` it will use predefined test urls.

Otherwise give a list of urls.


### Notes
Some sites(reddit, youtube) have ""rss"" feed links, but they are not actually RSS specification:
* `<entry>` instead of `<item>`
* `<published>` instead of `<pubData>`
* etc...

These feeds are a different standard - Atom https://en.wikipedia.org/wiki/Atom_(Web_standard) which is an alternative to RSS.
**This program does not support them.**
