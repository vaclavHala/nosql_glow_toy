#!/usr/bin/perl -pi
#multi-line in place substitute - subs.pl
use strict;
use warnings;

BEGIN {undef $/;}

s/\[?\s*\{\s*"_id": ([0-9]*),[^{]*\{\s*"ratingValue": ([0-9.]*),\s*"reviewCount": ([0-9]*)\s*\},\s*"description": "([a-zA-Z. ]+)",\s*"name": "([a-zA-Z]+)",[^a]*availability": "([a-zA-Z]+)",\s*"price": ([0-9.]+),\s*"[^}]*\}\s*\},?\s*\]?/INSERT INTO product (id, name, description, price, availability, currency, rating, ratingCount) VALUES ($1, '$5', '$4', $7, '$6', 'USD', $2, $3);\n/smg;
