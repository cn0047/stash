Smarty
-

version 2.6.26

````js
{$smarty.now|date_format:'%Y-%m-%d %H:%M:%S'}

{assign var='i' value=$k+1}
{$skill_{$skill_abbreviation}}

<pre>{$var|@var_export}<pre>

{','|implode:$options.errors}

{get_assoc array=$matrix key=$id assoc="other"}

{if in_array($brand.wl.wl_site_id, array(3818, 3788, 4047, 4262, 5656))}

{capture name=c}id={$row.userid}{/capture}
{$smarty.capture.c}
````
