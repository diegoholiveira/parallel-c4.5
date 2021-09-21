#!/usr/bin/env bash

(
  cd go || exit 1
  ./c45-generator
)

(
  cd python || exit 1
  poetry run classify
)
