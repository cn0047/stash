<?php

namespace VISA;

use Domain\Interfaces\VISA\EuropeanUnion as VISAInterface;

class EuropeanUnion implements VISAInterface
{
    public function approve(): string
    {
        return 'Your request is approved!';
    }
}
