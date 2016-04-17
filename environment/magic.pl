#!/usr/bin/perl -pi
#multi-line in place substitute - subs.pl
use strict;
use warnings;

BEGIN {undef $/;}

s/\[?\s*\{\s*"_id": ([0-9]*),[^{]*\{\s*"ratingValue": ([0-9.]*),\s*"reviewCount": ([0-9]*)\s*\},\s*"description": "([a-zA-Z. ]+)",\s*"name": "([a-zA-Z]+)",[^a]*availability": "([a-zA-Z]+)",\s*"price": ([0-9.]+),\s*"[^}]*\}\s*\},?\s*\]?/$1,$5,$6,$7,USD,$2,$3,$4\n/smg;
