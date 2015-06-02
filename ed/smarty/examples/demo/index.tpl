{foreach from=$data key=k item=i name=n}
    {if $smarty.foreach.n.index is div by 5}
        <div class="thumb-holder"></div>
    {/if}
    {$i}<br>
{/foreach}
