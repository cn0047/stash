<?php

$a = [
    [
        'id' => 1,
        'name' => 'node-1',
        'parent_id' => null,
    ],
    [
        'id' => 2,
        'name' => 'node-2',
        'parent_id' => null,
    ],
    [
        'id' => 3,
        'name' => 'node-3',
        'parent_id' => 1,
    ],
    [
        'id' => 4,
        'name' => 'node-4',
        'parent_id' => 1,
    ],
    [
        'id' => 5,
        'name' => 'node-5',
        'parent_id' => 1,
    ],
    [
        'id' => 6,
        'name' => 'node-6',
        'parent_id' => 2,
    ],
    [
        'id' => 7,
        'name' => 'node-7',
        'parent_id' => 6,
    ],
    [
        'id' => 8,
        'name' => 'node-8',
        'parent_id' => 7,
    ],
    [
        'id' => 9,
        'name' => 'node-9',
        'parent_id' => 3,
    ],
];
$tree = [];
while ($el = array_shift($a)) {
    if (empty($el['parent_id'])) {
        $tree[] = $el;
    } else {
        attachToParent($el, $tree);
    }
}
function attachToParent($el, &$tree) {
    foreach ($tree as &$parent) {
        if ($parent['id'] === $el['parent_id']) {
            $parent['childrens'][] = $el;
            return;
        }
        if (isset($parent['childrens'])) {
            $branch = &$parent['childrens'];
            attachToParent($el, $branch);
        }
    }
}

var_export($tree);

/*
array (
  0 =>
  array (
    'id' => 1,
    'name' => 'node-1',
    'parent_id' => NULL,
    'childrens' =>
    array (
      0 =>
      array (
        'id' => 3,
        'name' => 'node-3',
        'parent_id' => 1,
        'childrens' =>
        array (
          0 =>
          array (
            'id' => 9,
            'name' => 'node-9',
            'parent_id' => 3,
          ),
        ),
      ),
      1 =>
      array (
        'id' => 4,
        'name' => 'node-4',
        'parent_id' => 1,
      ),
      2 =>
      array (
        'id' => 5,
        'name' => 'node-5',
        'parent_id' => 1,
      ),
    ),
  ),
  1 =>
  array (
    'id' => 2,
    'name' => 'node-2',
    'parent_id' => NULL,
    'childrens' =>
    array (
      0 =>
      array (
        'id' => 6,
        'name' => 'node-6',
        'parent_id' => 2,
        'childrens' =>
        array (
          0 =>
          array (
            'id' => 7,
            'name' => 'node-7',
            'parent_id' => 6,
            'childrens' =>
            array (
              0 =>
              array (
                'id' => 8,
                'name' => 'node-8',
                'parent_id' => 7,
              ),
            ),
          ),
        ),
      ),
    ),
  ),
)
*/
