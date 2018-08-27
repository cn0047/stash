HTML (HyperText Markup Language)
-
5

````
<meta name="viewport" content="width=device-width, initial-scale=1.0">
````

#### Special Characters

````
&nbsp;     # space
````

#### Script

````
<script src="demo_async.js" async defer></script>
````

`async` - script will be executed asynchronously
as soon as it is available (only for external scripts).

`defer` - script won't run until page loaded.

If neither `async` or `defer` is present - script is fetched and executed immediately,
before the browser continues parsing the page.
