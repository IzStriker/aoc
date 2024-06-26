use strict;
use Class::Struct;
use List::Util qw(max first);

struct Coord => { x  => '$',     y  => '$' };
struct Pair  => { g1 => 'Coord', g2 => 'Coord' };

sub part1() {
    printf( "part 1 total: %d\n", get_total_lenghts(1) );
}

sub part2() {
    printf( "part 2 total: %d\n", get_total_lenghts(999999) );
}

sub get_total_lenghts($) {
    my ($expansion_amount) = @_;
    my $filename = "2023/11/input.txt";
    open( fh, $filename ) or die "Failed to open $filename";

    my @lines;
    my @galaxies;

    my $i = 0;
    while (<fh>) {
        push( @lines, $_ );
        @galaxies = ( @galaxies, get_all_occurences( $_, $i ) );
        $i += 1;
    }
    @galaxies = expand_universe( $expansion_amount, @galaxies );

    my @pairs = get_unique_pairs(@galaxies);

    my $total = 0;
    foreach (@pairs) {
        my $val = abs( $_->g1->x - $_->g2->x ) + abs( $_->g1->y - $_->g2->y );
        $total += $val;
    }

    return $total;
}

sub get_all_occurences( $$ ) {
    my ( $line, $y ) = @_;
    my $result = index( $line, "#" );
    my @xs;

    while ( $result != -1 ) {
        if ( $result != -1 ) {
            my $coord = Coord->new();
            $coord->x($result);
            $coord->y($y);

            push( @xs, $coord );
        }
        $result = index( $line, "#", $result + 1 );
    }

    return (@xs);
}

sub compare_coords( $$ ) {
    my ( $p1, $p2 ) = @_;

    return ( $p1->x == $p2->x and $p1->y == $p2->y );
}

sub get_unique_pairs(@) {
    my (@galaxies) = @_;
    my @pairs;
    my %seen_pairs;

    my $galaxy_count = scalar(@galaxies);

    for ( my $i = 0 ; $i < $galaxy_count ; $i++ ) {
        for ( my $j = $i + 1 ; $j < $galaxy_count ; $j++ ) {
            my $g1 = $galaxies[$i];
            my $g2 = $galaxies[$j];

            my $pair_key1 = join( ',', $g1, $g2 );
            my $pair_key2 = join( ',', $g2, $g1 );

            next if $seen_pairs{$pair_key1} or $seen_pairs{$pair_key2};

            if ( not compare_coords( $g1, $g2 ) ) {
                push( @pairs, Pair->new( g1 => $g1, g2 => $g2 ) );
                $seen_pairs{$pair_key1} = 1;
                $seen_pairs{$pair_key2} = 1;
            }
        }
    }

    return @pairs;
}

sub expand_universe($@) {
    my ( $expansion_amount, @galaxies ) = @_;

    my @empty_xs = missing_values( sort( map { $_->x } @galaxies ) );
    my @empty_ys = missing_values( sort ( map { $_->y } @galaxies ) );

    my %x_changes;
    my %y_changes;

    foreach my $x (@empty_xs) {
        foreach my $coord ( grep { $_->x > $x } @galaxies ) {
            my $key = sprintf( "%d-%d", $coord->x, $coord->y );
            $x_changes{$key} = $x_changes{$key} + $expansion_amount;
        }
    }

    foreach my $y (@empty_ys) {
        foreach my $coord ( grep { $_->y > $y } @galaxies ) {
            my $key = sprintf( "%d-%d", $coord->x, $coord->y );
            $y_changes{$key} = $y_changes{$key} + $expansion_amount;
        }
    }

    foreach my $coord (@galaxies) {
        my $key = sprintf( "%d-%d", $coord->x, $coord->y );
        $coord->x( $coord->x + $x_changes{$key} );
        $coord->y( $coord->y + $y_changes{$key} );
    }

    return @galaxies;
}

sub missing_values(@) {
    my (@values) = @_;
    my $max_val = max(@values);
    my @missing_vals;

    for ( my $i = 0 ; $i < $max_val ; $i++ ) {
        if ( not grep( { $_ == $i } @values ) ) {
            push( @missing_vals, $i );
        }
    }
    return @missing_vals;
}

part1();
part2();
