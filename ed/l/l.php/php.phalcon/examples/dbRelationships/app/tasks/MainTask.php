<?php

class MainTask extends \Phalcon\Cli\Task
{
    public function mainAction()
    {
        $robot = Robots::findFirst(1);
        printf("Robot: %s \n", $robot->name);
        $robotsParts = $robot->getRobotsParts();
        foreach ($robotsParts as $rp) {
            $part = $rp->getParts();
            printf("  robot id: %s, part id: %s, part name: %s \n", $rp->robots_id, $rp->parts_id, $part->name);
        }
    }
}
