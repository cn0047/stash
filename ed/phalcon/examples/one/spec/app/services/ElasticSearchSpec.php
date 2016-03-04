<?php

describe('ElasticSearch', function() {
    describe('getServiceName', function() {
        it('one', function() {
             $s = new \Services\ElasticSearch();
             $r = $s->getServiceName();
             expect($r)->toEqual('Services\\ElasticSearch');
        });
    });
});
