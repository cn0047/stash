jQuery
-

1.4.2

````js
jQuery.fn.jquery // version of jQuery
````

[jQuery in Action 2nd Edition](http://www.bibeault.org/jqia2/)

####Flashback
````js
// Selectors:
E>F         // element F that child of E
E+F         // all F before that is present E
E~F         // all F below E
:only-child // single child el
:eq(n)      // el with index n
:gt(n)      // el after n el
:lt(n)      // el before n el
:last
:animated
:checked
:parent
:selected
:visible

$('#el:hidden')
$('#el:not(.el)')
$('#el:has(.el)')
$('#el:contains(Bond)')
$('table tr td[data-field!=""]') // Not empty data-field.

.children()
.closest()
.next()
.parent()
.prev()
.siblings()
.is(selector)
.index()
.add()
.find(selector)
.not(expression)
.filter(expression)
.map(callback)
.each(iterator)
/*
$.each(yourObject, function (index, value) {
});
 */
.end() // prev selector in chain
.addSelf()
.data(name, value)
/*
console.log($(this).attr('data-pseudo-field'));
console.log($(this).data('pseudo-field'));
*/
.removeData(name)

jQuery.noCoflict();

$(function () {});
(function ($) {})(jQuery);
$(document).ready(function () {});

.removeAttr(attr);
.hasClass(name);
.addClass(name);
.removeClass(name);
.toggleClass(name);
.remove();
.detach();
.empty();

.bind(eventType, data, handler);
.one(eventType, data, listener);
.unbind(eventType, listener);
.live(eventType, data, listener);
.trigger(eventType, data);
.queue(name, func);
.dequeue();
.clearQueue();
.preventDefault();
.stopPropagation();
.isDefaultPrevented();
.isPropagationStopped();

.scroll();
.hover();
.toggle();
.serialize();
.serializeArray();

$.fx.off = true or false // Animation allowed

$.map(array, callback)
$.inArray(needle, haystack)
$.isArray(array)
$.isEmptyObject(object)
$.isPlainObject(object)
$.isFunction(func)
$.merge(array, array)
$.extend(deep, target, source1, sourceN)

$('*').attr('title', function (index, previousValue) {
    return previousValue + ' i am el:' + index;
});

var settings = $.extend({x: 0}, options || {});
````

####UI
````js
// Speeds
slow|normal|fast
// Effects
blind|bounce|clip|drop|explode|fade|fold|highlight|puff|pulsate|scale|shake|size|slide|transfer

$('*').show('fast');
$('*').hide('puff');

draggable(options, optionName, value)
droppable(options, optionName, value)
sortable(options, optionName, value)
resizable(options, optionName, value)
selectable(options, optionName, value)
````

####Ajax
````js
$('form').ajaxForm({});

jQuery.getJSON(url, [data], [callback]);

$.ajax({
    url: 'index.php',
    dataType: 'json',
    type: 'post',
    data: {},
    success: function (r) {
    },
    error: function (xhr, status, error) {
    },
    complete: function () {
    },
});

load(url, parameters, function (xhr, status, error) {});

$.get(url, parameters, callback, type);

$.getJSON(url, parameters, callback);

$.post(url, parameters, callback, type);
````
