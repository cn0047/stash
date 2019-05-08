#!/bin/bash

while IFS='$\n' read -r v; do
    echo "${v}"
done
