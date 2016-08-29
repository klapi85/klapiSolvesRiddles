use strict;

## W. Kuchta 2016
## https://www.hackerrank.com/challenges/mars-exploration

my $s = <STDIN>;

my $sum = 0;

while ($s =~ /(...)([A-Za-z0-9]*)/) {
    $s = $2;
    my $result = checkTriple($1);
    $sum += $result;
}
print $sum;

sub checkTriple() {
        my ($t) = @_;
        my $result = 0;

        if ('S' ne substr($t, 0, 1)) { $result++; }
        if ('O' ne substr($t, 1, 1)) { $result++; }
        if ('S' ne substr($t, 2, 1)) { $result++; }

        return $result;
}
