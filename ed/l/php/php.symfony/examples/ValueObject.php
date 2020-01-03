<?php
/**
 * To use this class just run:
 * @example `$filter = PlayerFilter::createFromRequest($request);`
 */

declare(strict_types=1);

namespace AppBundle\Model;

use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\Validator\Constraints as Assert;

class PlayerFilter
{
    /**
     * @Assert\NotNull()
     * @Assert\GreaterThanOrEqual(1)
     */
    private $page;

    /**
     * @Assert\NotNull()
     * @Assert\GreaterThanOrEqual(1)
     * @Assert\LessThanOrEqual(200)
     */
    private $limit;

    /**
     * @Assert\Choice(choices = {
     *     "id",
     *     "-id",
     *     "externalId",
     *     "-externalId",
     *     "country",
     *     "-country",
     *     "sex",
     *     "-sex",
     *     "currency",
     *     "-currency",
     *     "jurisdiction",
     *     "-jurisdiction",
     *     null
     * },
     *      message = "Invalid query parameter."
     * )
     */
    private $sortBy;

    /**
     * @Assert\Choice(choices = {"-1","0","1"}, message = "Invalid query parameter.")
     */
    private $active;

    public function __construct(int $active, int $page, int $limit, string $sortBy = null)
    {
        $this->active = $active;
        $this->page = $page;
        $this->limit = $limit;
        $this->sortBy = $sortBy;
    }

    public static function createFromRequest(Request $request): PlayerFilter
    {
        $active = (int) $request->query->get('active', 0);
        $page = (int) $request->query->get('page', 1);
        $limit = (int) $request->query->get('limit', 10);
        $sortBy = (string) $request->query->get('sortBy');

        return new self($active, $page, $limit, $sortBy);
    }

    public function getOffset(): int
    {
        return ($this->page - 1) * $this->limit;
    }

    public function getLimit(): int
    {
        return $this->limit;
    }

    public function getOrder(): string
    {
        if (empty($this->sortBy)) {
            return 'ASC';
        }

        if ($this->sortBy[0] === '-') {
            return 'DESC';
        }

        return 'ASC';
    }

    public function getSort(string $alias): string
    {
        $sortBy = $this->sortBy;

        $sort = 'id';
        if (!empty($this->sortBy)) {
            $sort = $sortBy[0] === '-' ? substr($sortBy, 1) : $sortBy;
        }

        return sprintf('%s.%s', $alias, $sort);
    }

    public function isActive(): ?bool
    {
        if ($this->active === -1) {
            return null;
        }

        return (bool) $this->active;
    }
}
