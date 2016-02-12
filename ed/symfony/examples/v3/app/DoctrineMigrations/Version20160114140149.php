<?php

namespace Application\Migrations;

use Doctrine\DBAL\Migrations\AbstractMigration;
use Doctrine\DBAL\Schema\Schema;

/**
 * Auto-generated Migration: Please modify to your needs!
 */
class Version20160114140149 extends AbstractMigration
{
    // password = guest
    private $password = '$2y$13$dBieYBsXQkfB1X8Uq9WjTO.Do6Z79PV6OoW7fCSBX2y6yd4XRdFFu';

    /**
     * @param Schema $schema
     */
    public function up(Schema $schema)
    {
        $this->addSql("
            INSERT INTO fos_user SET
                 username              = 'guest'
                ,username_canonical    = 'guest'
                ,email                 = 'guest@localhost.com'
                ,email_canonical       = 'guest@localhost.com'
                ,enabled               = 1
                ,salt                  = 'm96wm4iincgoggs0wcwwwksksscswkc'
                ,password              = '{$this->password}'
                ,locked                = 0
                ,expired               = 0
                ,expires_at            = NULL
                ,confirmation_token    = NULL
                ,password_requested_at = NULL
                ,roles                 = 'a:0:{}'
                ,credentials_expired   = 0
                ,credentials_expire_at = NULL
        ");
    }

    /**
     * @param Schema $schema
     */
    public function down(Schema $schema)
    {
        $this->addSql("
            DELETE FROM fos_user WHERE
                    username              =  'guest'
                AND username_canonical    =  'guest'
                AND email                 =  'guest@localhost.com'
                AND email_canonical       =  'guest@localhost.com'
                AND enabled               =  1
                AND salt                  =  'm96wm4iincgoggs0wcwwwksksscswkc'
                AND password              =  '{$this->password}'
                AND locked                =  0
                AND expired               =  0
                AND expires_at            IS NULL
                AND confirmation_token    IS NULL
                AND password_requested_at IS NULL
                AND roles                 =  'a:0:{}'
                AND credentials_expired   =  0
                AND credentials_expire_at IS NULL
        ");
    }
}
