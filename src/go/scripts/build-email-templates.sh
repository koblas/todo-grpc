#!/usr/bin/env bash

for f in services/core/*/[a-zA-Z]*.mjml ; do
  npx mjml $f -o `echo $f | sed -e "s/-mjml//" | sed -e "s/mjml/html/"`
done
