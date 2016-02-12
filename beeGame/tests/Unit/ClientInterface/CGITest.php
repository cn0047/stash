<?php

namespace Test\Unit\ClientInterface;

use ClientInterface\CGI;

class CGITest extends \PHPUnit_Framework_TestCase
{
    /**
     * @expectedException \DomainException
     */
    public function testConstruct()
    {
        new CGI();
    }
}
