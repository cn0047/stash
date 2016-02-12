<?php

namespace State;

use Command\CommandInterface;

/**
 * State of game interface.
 *
 * State encapsulates rules for game in particular moment,
 * and can lead game in one from few available ways.
 * State helps maintain game by providing fixed set of possible next commands.
 */
interface StateInterface
{
    /**
     * Gets message to interact with client (gamer).
     *
     * Usually this message is the question about next action (command).
     *
     * @return string Message for client (gamer).
     */
    public function getPromptMessage();

    /**
     * Command that will be executed in case when user pass "YES" to question.
     *
     * @return CommandInterface Command.
     */
    public function getPromptedCommand();

    /**
     * Command that will be executed in case when user pass "NO" to question.
     *
     * @return CommandInterface Command.
     */
    public function getNotPromptedCommand();
}
