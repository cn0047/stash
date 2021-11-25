Analytics
-

https://ga-dev-tools.appspot.com/campaign-url-builder
https://console.developers.google.com/cloud-resource-manager

````html
<a href="" onclick="_gtag.push(['_trackPageview', 'myLink'])">myLink</a>
````

````js
_gtag.push(['_trackPageview', 'myVirtualPage'])
_gtag.push(['_trackEvent', 'action'])

_gtag.push(['_tackSocial', 'facebook', 'like', 'fx.ua'])
_gtag.push(['_setSustomVar', 'slot', 'name', 'value', 'scope'])
_gtag.push(['_setSustomVar', 1, 'role', 'customer', 1])
// where scopes: 3 - page; 2 - session; 1 - visitor.
````

````js
gtag('config', '#{googleAnalyticsId}', {'custom_map': {'metric1': 'streamIdMismatch'}});
gtag('event', 'streamIdMismatchCount', {'streamIdMismatch': 1});
````

Event Tracking:
````sh
v=1              # Version.
&tid=UA-XXXXX-Y  # Tracking ID / Property ID.
&cid=555         # Anonymous Client ID (random integer). If this value not changed - total count in GA report will not change.
&t=event         # Event hit type
&ec=video        # Event Category. Required.
&ea=play         # Event Action. Required.
&el=holiday      # Event label.
&ev=300          # Event value.

tid='UA-109758529-1'
cid=`php -r 'echo rand();'`
curl -i 'https://www.google-analytics.com/collect?v=1&t=event&tid='$tid'&cid='$cid'&ec=shellClick&ea=click.shell&el=cli&ev=1'
````
