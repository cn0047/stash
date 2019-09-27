<?php

namespace AppBundle\Security;

use Symfony\Component\Security\Core\Authentication\Token\TokenInterface;
use Symfony\Component\Security\Core\Authorization\Voter\Voter;
use Symfony\Component\Security\Core\User\User;

class PostVoter extends Voter
{
    protected function supports($attribute, $subject)
    {
        if ($attribute === 'view') {
            return true;
        }
        if ($attribute === 'edit') {
            return true;
        }

        return false;
    }

    protected function voteOnAttribute($attribute, $subject, TokenInterface $token)
    {
        $user = $token->getUser();

        if (!$user instanceof User) {
            return false;
        }

        switch ($attribute) {
            case 'view':
                return $this->canView($subject, $user);
            case 'edit':
                return $this->canEdit($subject, $user);
        }

        throw new \LogicException('PostVoterFail.');
    }

    private function canView(array $subject, User $user)
    {
        return true === in_array('ROLE_USER', $user->getRoles(), true);
    }

    private function canEdit(array $subject, User $user)
    {
        return true === in_array('ROLE_ADMIN', $user->getRoles(), true);
    }
}
