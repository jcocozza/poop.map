#!/bin/bash
#
# Quick script to build and run for web

flutter build web
cd build/web
python3 -m http.server 8011
