{$skill_{$skill_abbreviation}}
// or
{assign var='myVar' value=$skill_{$skill_abbreviation}}
===
{foreach from=$skill_abbreviations item=abbr}
  {$skill_{$abbr}}
{/foreach}
