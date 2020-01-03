<?php

use One\Text;
use One\Foo;
use Kahlan\Plugin\Stub;
use Kahlan\Plugin\Double;

describe('Text', function () {
    describe('->get()', function() {

        it('test', function () {
            $t = new Text();
            $r = $t->get();
            expect($r)->toBe('BAR');
        });

        it('mock', function () {
            Stub::on(Foo::class)->method('bar', function () {
                return 'MyMockedText';
            });
            $r = (new Text())->get();
            expect($r)->toBe('MyMockedText');
        });

        it('mock 2', function () {
             allow(Foo::class)->toReceive('bar')->andReturn('MyAnotherMockedText');
             $r = (new Text())->get();
             expect($r)->toBe('MyAnotherMockedText');
        });

        it('time', function () {
            allow('time')->toBeCalled()->andReturn(100);
            $r = (new Text())->time();
            expect($r)->toBe(100);
        });

        it('dt', function () {
            allow(DateTime::class)->toReceive('::createFromFormat')->andReturn('FormatED');
            $r = (new Text())->dt();
            expect($r)->toBe('FormatED');
        });
    });
});
