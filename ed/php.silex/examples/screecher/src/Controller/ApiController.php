<?php

use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Validator\Constraints as Assert;

$apis = $app['controllers_factory'];

$apis->get('/', function () use ($app) {
    try {
        $sql = 'SELECT * FROM api';
        $sth = $app['database']->prepare($sql);
        if (!$sth->execute()) {
            throw new RuntimeException($sth->errorInfo());
        }
        $result = [
            'success' => true,
            'results' => $sth->fetchAll(PDO::FETCH_ASSOC),
        ];
        return $app->json($result, 200);
    } catch (RuntimeException $e) {
        $result = [
            'success' => false,
            'error' => 'Runtime error.',
        ];
        return $app->json($result, 500);
    } catch (Exception $e) {
        $result = [
            'success' => false,
            'error' => 'Internal server error.',
        ];
        return $app->json($result, 500);
    }
});

$apis->get('/{id}', function ($id) use ($app) {
    try {
        // Validation.
        $errors = $app['validator']->validateValue($id, new Assert\Regex('/^\d+$/'));
        if (count($errors) > 0) {
            throw new InvalidArgumentException('id invalid.');
        }
        // Select data from db.
        $sql = 'SELECT * FROM api WHERE id = :id';
        $sth = $app['database']->prepare($sql);
        $sth->bindValue(':id', $id, PDO::PARAM_STR);
        if (!$sth->execute()) {
            throw new RuntimeException($sth->errorInfo());
        }
        $result = [
            'success' => true,
            'results' => $sth->fetch(PDO::FETCH_ASSOC),
        ];
        return $app->json($result, 200);
    } catch (InvalidArgumentException $e) {
        $result = [
            'success' => false,
            'error' => $e->getMessage(),
        ];
        return $app->json($result, 400);
    } catch (RuntimeException $e) {
        $result = [
            'success' => false,
            'error' => 'Runtime error.',
        ];
        return $app->json($result, 500);
    } catch (Exception $e) {
        $result = [
            'success' => false,
            'error' => 'Internal server error.',
        ];
        return $app->json($result, 500);
    }
});

$apis->post('/', function (Request $request) use ($app) {
    try {
        $parameters = json_decode($request->getContent(), true);
        if (json_last_error() !== JSON_ERROR_NONE) {
            throw new InvalidArgumentException('Invalid input parameters.');
        }
        // Validation.
        if (!isset($parameters['name'])) {
            throw new InvalidArgumentException('Name cannot be blank.');
        }
        $errors = $app['validator']->validateValue($parameters['name'], new Assert\Regex('/^[\w\s]+$/'));
        if (count($errors) > 0) {
            throw new InvalidArgumentException('Name invalid.');
        }
        // Find duplicates.
        $sql = 'SELECT id FROM api WHERE name = :name';
        $sth = $app['database']->prepare($sql);
        $sth->bindValue(':name', $parameters['name'], PDO::PARAM_STR);
        if (!$sth->execute()) {
            throw new RuntimeException($sth->errorInfo());
        }
        $id = $sth->fetchColumn();
        if ($id !== false) {
            throw new InvalidArgumentException('Name already present in db.');
        }
        // Add to db.
        $sql = 'INSERT INTO api SET name = :name';
        $sth = $app['database']->prepare($sql);
        $sth->bindValue(':name', $parameters['name'], PDO::PARAM_STR);
        if (!$sth->execute()) {
            throw new RuntimeException($sth->errorInfo());
        }
        $result = [
            'success' => true,
            'id' => $app['database']->lastInsertId(),
        ];
        $result += $parameters;
        return $app->json($result, 201);
    } catch (InvalidArgumentException $e) {
        $result = [
            'success' => false,
            'error' => $e->getMessage(),
        ];
        return $app->json($result, 400);
    } catch (RuntimeException $e) {
        $result = [
            'success' => false,
            'error' => 'Runtime error.',
        ];
        return $app->json($result, 500);
    } catch (Exception $e) {
        $result = [
            'success' => false,
            'error' => 'Internal server error.',
        ];
        return $app->json($result, 500);
    }
});

$apis->put('/{id}', function (Request $request, $id) use ($app) {
    try {
        $parameters = json_decode($request->getContent(), true);
        // Validation.
        $errors = $app['validator']->validateValue($id, new Assert\Regex('/^\d+$/'));
        if (count($errors) > 0) {
            throw new InvalidArgumentException('id invalid.');
        }
        if (!isset($parameters['name'])) {
            throw new InvalidArgumentException('Name cannot be blank.');
        }
        $errors = $app['validator']->validateValue($parameters['name'], new Assert\Regex('/^[\w\s]+$/'));
        if (count($errors) > 0) {
            throw new InvalidArgumentException('Name invalid.');
        }
        // Update at db.
        $sql = 'UPDATE api SET name = :name WHERE id = :id';
        $sth = $app['database']->prepare($sql);
        $sth->bindValue(':id', $id, PDO::PARAM_STR);
        $sth->bindValue(':name', $parameters['name'], PDO::PARAM_STR);
        if (!$sth->execute()) {
            throw new RuntimeException($sth->errorInfo());
        }
        $result = [
            'success' => true,
            'id' => $id,
        ];
        $result += $parameters;
        return $app->json($result, 200);
    } catch (InvalidArgumentException $e) {
        $result = [
            'success' => false,
            'error' => $e->getMessage(),
        ];
        return $app->json($result, 400);
    } catch (RuntimeException $e) {
        $result = [
            'success' => false,
            'error' => 'Runtime error.',
        ];
        return $app->json($result, 500);
    } catch (Exception $e) {
        $result = [
            'success' => false,
            'error' => 'Internal server error.',
        ];
        return $app->json($result, 500);
    }
});

$apis->delete('/{id}', function (Request $request, $id) use ($app) {
    try {
        // Validation.
        $errors = $app['validator']->validateValue($id, new Assert\Regex('/^\d+$/'));
        if (count($errors) > 0) {
            throw new InvalidArgumentException('id invalid.');
        }
        // Delete from db.
        $sql = 'DELETE FROM api WHERE id = :id';
        $sth = $app['database']->prepare($sql);
        $sth->bindValue(':id', $id, PDO::PARAM_STR);
        if (!$sth->execute()) {
            throw new RuntimeException($sth->errorInfo());
        }
        $result = [
            'success' => true,
        ];
        return $app->json($result, 200);
    } catch (InvalidArgumentException $e) {
        $result = [
            'success' => false,
            'error' => $e->getMessage(),
        ];
        return $app->json($result, 400);
    } catch (RuntimeException $e) {
        $result = [
            'success' => false,
            'error' => 'Runtime error.',
        ];
        return $app->json($result, 500);
    } catch (Exception $e) {
        $result = [
            'success' => false,
            'error' => 'Internal server error.',
        ];
        return $app->json($result, 500);
    }
});

return $apis;
