<?php

namespace AppBundle\Twig;

class DateExtension extends \Twig_Extension
{
    public function getFilters()
    {
        return array(
            new \Twig_SimpleFilter('date', array($this, 'dateFilter')),
        );
    }

    public function dateFilter($time, $timeZone = null)
    {
        $tz = null;
        if (null !== $timeZone) {
            $tz = new \DateTimeZone($timeZone);
        }
        $d = new \DateTime($time, $tz);
        return $d->format('Y-m-d H:i:s');
    }

    public function getName()
    {
        return 'date_extension';
    }
}
