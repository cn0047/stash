<?php

namespace spec\App;

use App\Boo;
use PhpSpec\ObjectBehavior;
use Prophecy\Argument;

class BooSpec extends ObjectBehavior
{
    function it_is_initializable()
    {
        $this->shouldHaveType(Boo::class);
    }
}
