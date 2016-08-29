use strict;

## W. Kuchta 2016
## https://www.hackerrank.com/challenges/pangrams

my $s = <STDIN>;
$s = lc($s);

my @chars = split //, $s;
my @sorted = sort @chars;
my @unique = uniq(@sorted);
my $short = "@unique";
$short =~ s/ //g;

if ($short eq 'abcdefghijklmnopqrstuvwxyz') {
    print "pangram";
} else {
    print "not pangram";
}

sub uniq {
  my %seen;
  return grep { !$seen{$_}++ } @_;
}
