<?php

class Auth
{
    private $login = '';
    private $password = '';
    private $passwordHash = '';

    public function IsAuthenticated()
    {
        try {
            if (!isset($_SERVER['PHP_AUTH_USER']) or !isset($_SERVER['PHP_AUTH_PW'])) {
                throw new Exception('Authorization Required.');
            }
            if ($_SERVER['PHP_AUTH_USER'] != $this->login) {
                throw new Exception('Incorrect login.');
            }
            if (!password_verify($_SERVER['PHP_AUTH_PW'], $this->passwordHash)) {
                throw new Exception('Incorrect password.');
            }
            return ture;
        } catch (Exception $e) {
            header('WWW-Authenticate: Basic realm="Authorization Required"');
            header('HTTP/1.0 401 Unauthorized');
            echo $e->getMessage();
            return false;
        }
    }
}
