grammar PricingDSL;

start : rule+ EOF;

rule
    : RULE STRING IF expression THEN modifier
      (PRIORITY NUMBER)?
      (GROUP groupType)?
    ;

modifier
    : modifierType NUMBER
    ;

modifierType
    : PERCENT
    | FIXED
    | FINAL
    ;

groupType
    : COUPON
    | PROMOCODE
    | CERT
    ;

expression
    : comparison ((AND | OR) comparison)*
    | TRUE
    ;

comparison
    : identifier comparator value
    ;

comparator
    : EQ
    | NE
    | GE
    | GT
    | LE
    | LT
    ;

identifier
    : ID ('.' ID)*
    ;

value
    : STRING
    | NUMBER
    ;

// Keywords
RULE      : 'RULE';
IF        : 'IF';
THEN      : 'THEN';
PRIORITY  : 'PRIORITY';
GROUP     : 'GROUP';

PERCENT   : 'percent';
FIXED     : 'fixed';
FINAL     : 'final';

COUPON    : 'coupon';
PROMOCODE : 'promocode';
CERT      : 'cert';

AND       : 'AND';
OR        : 'OR';
TRUE      : 'true';

// Operators
EQ        : '==';
NE        : '!=';
GE        : '>=';
GT        : '>';
LE        : '<=';
LT        : '<';

// Values
ID        : [a-zA-Z] [a-zA-Z0-9_]*;
NUMBER    : [0-9]+ ('.' [0-9]+)?;
STRING    : '"' (ESCAPE | ~["\\])* '"';

fragment ESCAPE
    : '\\' ('"' | '\\')
    ;

WS
    : [ \t\r\n]+ -> skip
    ;

