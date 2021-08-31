robots.txt
-

robots.txt - text file with instructions for search engine crawlers.

````html
<meta name="robots" content="noindex">
````

````sh
User-agent: *
Disallow: /login/
Disallow: /card/
Disallow: /fotos/
Disallow: /temp/
Disallow: /search/
Disallow: /*.pdf$

Sitemap: https://www.example.com/sitemap.xml
````
