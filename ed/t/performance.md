Performance
-

[page speed](www.bytecheck.com) [and](www.webpagetest.org)

Resilience - is the ability of a system to adapt or keep working when challenges occur.

#### High Performance Web Sites:

1. Make Fewer HTTP Requests
  * Image Maps
  * ~~CSS Sprites~~
  * Inline Images
  * Combined Scripts and Stylesheets

2. Use a Content Delivery Network

3. Add an Expires Header
  * Expires Header
  * Max-Age and mod_expires

4. Gzip Components (scripts and stylesheets)

5. Put Stylesheets at the Top (in the document HEAD using the LINK tag)

6. Put Scripts at the Bottom (before </body>)
  * Don't forget about parallel(async)|subsequent|defer(page parsed) downloads
  * Scripts Block Downloads

7. Avoid CSS Expressions

8. Make JavaScript and CSS External (put it to external files)

9. Reduce DNS Lookups (using Keep-Alive and fewer domains)

10. Minify JavaScript
  * Minification
  * Obfuscation

11. Avoid Redirects
  * Alternatives to Redirects (Missing Trailing Slash, and else)
  * 301 HTTP Coode for any resource request

12. Remove Duplicate Scripts

13. Configure or remove ETags

14. Make Ajax Cacheable

#### General

* Gzip JSON responses.
* Use HTTP2.
* If site doesn't have favicon - 2 overhead request per every request.
* `site.com & api.site.com` better than `site.com & site.com/api`
  because of separation between web & api traffic (easier to scale and deal with LBs).

#### Scaling

![horizontal vs vertical 1](https://gist.github.com/cn007b/2c63c4b626be598166d5bce28b82552e/raw/633cd2dd905093aa584c2dc2fe44678804e96d14/1.2.jpeg)
![horizontal vs vertical 2](https://gist.github.com/cn007b/2c63c4b626be598166d5bce28b82552e/raw/633cd2dd905093aa584c2dc2fe44678804e96d14/1.jpeg)
![horizontal vs vertical 3](https://gist.github.com/cn007b/2c63c4b626be598166d5bce28b82552e/raw/633cd2dd905093aa584c2dc2fe44678804e96d14/2.jpeg)
![horizontal](https://gist.github.com/cn007b/2c63c4b626be598166d5bce28b82552e/raw/633cd2dd905093aa584c2dc2fe44678804e96d14/3.jpeg)
