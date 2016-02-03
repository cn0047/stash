<?php

namespace State;

interface StateInterface
{
    public function getPromptMessage();

    public function getPromptedCommand();

    public function getNotPromptedCommand();
}
