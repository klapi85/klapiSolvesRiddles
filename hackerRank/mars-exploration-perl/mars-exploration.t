use strict;
use warnings;
use Test::More tests => 3;

require './mars-exploration.pl';

is(checkTriple('SOS'), 0, 'SOS has no errors');
is(checkTriple('SZS'), 1, 'SZS has one error');
is(checkTriple('SOSSOS'), 0, 'SOSSOS has no errors');
