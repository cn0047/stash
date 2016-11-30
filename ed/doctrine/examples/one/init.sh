#!/bin/bash

# create tables in db from entities
vendor/bin/doctrine orm:schema:update --force
