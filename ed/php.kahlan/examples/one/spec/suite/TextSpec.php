<?php

use One\Text;
use Kahlan\Plugin\Stub;

describe('Text', function () {
    describe('->get()', function() {

        it('test get 1', function () {
            $t = new Text();
            $r = $t->get('-test-');
            expect($r)->toBe('One\\Text::get-test-');
        });

        it('mock', function () {
            Stub::on('Tex')->method('get', function ($str) {
                return 'mocked method result';
            });
            $t = new Text();
            $r = $t->get('-test-2-');
            expect($r)->toBe('One\\Text::get-test-2-');
        });

    });
});
