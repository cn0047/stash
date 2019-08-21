HTML (HyperText Markup Language)
-
5

````
<meta name="viewport" content="width=device-width, initial-scale=1.0">
````

#### Special Characters

````
&nbsp;     # space
&#34;      # double quotes
&quot;     # double quotes

# trik:
<table dir="ltr" border="1" cellspacing="0" cellpadding="0"><colgroup><col width="76"/><col width="90"/></colgroup><tbody><tr><td data-sheets-value="{&#34;1&#34;:2,&#34;2&#34;:&#34;Active&#34;}">Active</td><td data-sheets-value="{&#34;1&#34;:2,&#34;2&#34;:&#34;Betsy&#34;}">Betsyx</td></tr></tbody></table></textarea>

<p>
<span data-sheets-value="{&#34;1&#34;:2,&#34;2&#34;:&#34;I would describe my style as a mix of classic elements with current items on trend. I let what i&#39;m reading at the moment influence how I see the world and that is often perceived through my style of the moment. &#34;}" 
data-sheets-userformat="{&#34;2&#34;:513,&#34;3&#34;:{&#34;1&#34;:0},&#34;12&#34;:0}">
<strong><span data-sheets-value="{&#34;1&#34;:2,&#34;2&#34;:&#34;Colorful, sporty and feminine always with a bit of vintage&#34;}" 
data-sheets-userformat="{&#34;2&#34;:513,&#34;3&#34;:{&#34;1&#34;:0},&#34;12&#34;:0}">The one trend she&#39;s keen on: </span></strong>
<span data-sheets-value="{&#34;1&#34;:2,&#34;2&#34;:&#34;Colorful, sporty and feminine always with a bit of vintage&#34;}" 
data-sheets-userformat="{&#34;2&#34;:513,&#34;3&#34;:{&#34;1&#34;:0},&#34;12&#34;:0}">
<span data-sheets-value="{&#34;1&#34;:2,&#34;2&#34;:&#34;Leather capris are on my radar &#34;}" 
data-sheets-userformat="{&#34;2&#34;:513,&#34;3&#34;:{&#34;1&#34;:0},&#34;12&#34;:0}">Leather capris are on my radar.</span></span></span>
</p>
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

#### Microdata

````
<span itemprop="name">Elizabeth</span>
````

#### Form

"multipart/form-data"

````
<form action="file" method="post" enctype="multipart/form-data">
  <input type="text" id="msg" name="msg">
  <input type="file" id="file" name="file">
  <input type="submit" name="submit" value="Upload">
</form>
````
