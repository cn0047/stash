<?php

namespace spec\App;

use App\Foo;
use App\Boo;
use PhpSpec\ObjectBehavior;
use PhpSpec\Wrapper\Subject;

/**
 * @method Subject bar
 * @method Subject sbar
 * @method Subject baz
 */
class FooSpec extends ObjectBehavior
{
    public function it_is_initializable()
    {
        $this->shouldHaveType(Foo::class);
    }

    public function it_is_bar()
    {
        $this->bar()->shouldReturn('bar');
    }

    public function it_is_sbar()
    {
        $this->sbar()->shouldReturn('sbar');
    }

    public function it_is_baz()
    {
        $this->baz()->shouldReturn('baz');
    }

    public function it_is_baz_mock(Boo $boo)
    {
        // PROBLEM
        $boo->baz()->willReturn('yes');
        $this->baz()->shouldReturn('no');
    }

    public function it_is_sbaz_mock(Boo $boo)
    {
        // PROBLEM
        $boo->sbaz()->willReturn('syes');
        $this->sbaz()->shouldReturn('sno');
    }
}
