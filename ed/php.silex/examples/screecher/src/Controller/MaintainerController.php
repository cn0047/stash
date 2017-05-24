<?php

use Doctrine\DBAL\Exception\ForeignKeyConstraintViolationException;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Validator\Constraints as Assert;

$maintainers = $app['controllers_factory'];

$maintainers->get('/', function () use ($app) {
    $q = $app['entityManager']
        ->createQuery('SELECT m FROM Screecher\Entity\Maintainer m');
    $result = [
        'success' => true,
        'results' => $q->getArrayResult(),
    ];
    return $app->json($result, 200);
});

$maintainers->get('/{id}', function ($id) use ($app) {
    try {
        // Validation.
        $errors = $app['validator']->validateValue($id, new Assert\Regex('/^\d+$/'));
        if (count($errors) > 0) {
            throw new InvalidArgumentException('id invalid.');
        }
        // Select data from db.
        $q = $app['entityManager']
            ->createQuery('SELECT m FROM Screecher\Entity\Maintainer m WHERE m.id = ?1');
        $q->setParameter(1, $id);
        $result = [
            'success' => true,
            'results' => $q->getArrayResult(),
        ];
        return $app->json($result, 200);
    } catch (InvalidArgumentException $e) {
        $result = [
            'success' => false,
            'error' => $e->getMessage(),
        ];
        return $app->json($result, 400);
    } catch (Exception $e) {
        $result = [
            'success' => false,
            'error' => 'Internal server error.',
        ];
        return $app->json($result, 500);
    }
});

$maintainers->post('/', function (Request $request) use ($app) {
    try {
        $parameters = json_decode($request->getContent(), true);
        $maintainer = new Screecher\Entity\Maintainer();
        $maintainer->setProperties($parameters);
        // Validation.
        $errors = $app['validator']->validate($maintainer);
        if (count($errors) > 0) {
            $message = '';
            foreach ($errors->getIterator() as $obj) {
                $message .= $obj->getPropertyPath().' - '.$obj->getMessage();
            }
            throw new InvalidArgumentException($message);
        }
        // Add to db.
        $app['entityManager']->persist($maintainer);
        $app['entityManager']->flush();
        $result = [
            'success' => true,
            'id' => $maintainer->getId(),
        ];
        $result += $parameters;
        return $app->json($result, 201);
    } catch (InvalidArgumentException $e) {
        $result = [
            'success' => false,
            'error' => $e->getMessage(),
        ];
        return $app->json($result, 400);
    } catch (ForeignKeyConstraintViolationException $e) {
        $result = [
            'success' => false,
            'error' => 'Api id not found.',
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

$maintainers->put('/{id}', function (Request $request, $id) use ($app) {
    return 'Not implemented yet.';
});

$maintainers->delete('/{id}', function (Request $request, $id) use ($app) {
    return 'Not implemented yet.';
});

return $maintainers;
