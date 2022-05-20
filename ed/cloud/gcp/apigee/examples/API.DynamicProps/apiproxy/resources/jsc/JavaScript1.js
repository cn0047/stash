var strategy = context.getVariable('propertyset.main.strategy');

if (strategy==='v1') {
    var p = 'v1/messages/x';
    context.setVariable('target.url', 'https://httpbin.org/post?v=v1&name=1&target_path='+p);
} else if (strategy==='fromV1toV2') {
    var c = context.getVariable('request.header.c');
    var l = context.getVariable('request.formparam.l');
    var p = 'v2/c/'+c+'/l/'+l+'/x';
    context.setVariable('target.url', 'https://httpbin.org/post?v=v2&name=2&target_path='+p);
} else {
    context.setVariable('target.url', 'https://httpbin.org/post?strategy='+strategy);
}
