lexer grammar Nlexer;


options {
    language = Go;
}

fragment A      : ('A'|'a');
fragment B      : ('B'|'b');
fragment C      : ('C'|'c');
fragment D      : ('D'|'d');
fragment E      : ('E'|'e');
fragment F      : ('F'|'f');
fragment G      : ('G'|'g');
fragment H      : ('H'|'h');
fragment I      : ('I'|'i');
fragment J      : ('J'|'j');
fragment K      : ('K'|'k');
fragment L      : ('L'|'l');
fragment M      : ('M'|'m');      
fragment N      : ('N'|'n');
fragment O      : ('O'|'o');
fragment P      : ('P'|'p');
fragment Q      : ('Q'|'q');
fragment R      : ('R'|'r');
fragment S      : ('S'|'s');
fragment T      : ('T'|'t');
fragment U      : ('U'|'u');
fragment V      : ('V'|'v');
fragment W      : ('W'|'w');
fragment X      : ('X'|'x');
fragment Y      : ('Y'|'y');
fragment Z      : ('Z'|'z');



SENTENCIA       : S E N T E N C I A;
AND             : A N D;
OR              : O R;
IF              : I F;
ENTONCES        : E N T O N C E S;
CONSOLA         : C O N S O L A;
STRING          : S T R I N G;
INTEGER         : I N T E G E R;
REAL            : R E A L;
BOOLEAN         : B O O L E A N;
TRUE            : T R U E;
FALSE           : F A L S E;
PUBLICO         : P U B L I C O;
PRIVADO         : P R I V A D O;
MAIN            : M A I N;
PRINCIPAL       : P R I N C I P A L;
METODO          : M E T O D O;
CLASE           : C L A S E;
DECLARAR        : D E C L A R A R;

NUMERO          : [0-9]+;
DECIMAL         : [0-9]+'.'[0-9]+;
ID              : ([a-zA-Z]|[0-9])([a-zA-Z]|[0-9]|'_')*;

DOSPUNTOS       : ':';
PUNTO           : '.';
PUNTOCOMA       : ';';
MAYOR_I         : '>=';
MAYOR           : '>';
MENOR           : '<';
IGUAL           : '=';
MODULO          : '%';
MULTIPLICACION  : '*' ;
DIVISION        : '/' ;
RESTA           : '-' ;
SUMA            : '+' ;
PAR_IZQ         : '(';
PAR_DER         : ')';
LLAVE_IZQ       : '{';
LLAVE_DER       : '}';
CORCHETE_IZQ    : '[';
CORCHETE_DER    : ']';
CADENA          : '"' .*? '"';
WHITESPACE      : [ \\\r\n\t]+ -> skip;
COMMENT         : '/*' .*? '*/' -> skip;
LINE_COMMENT    : '//' ~[\r\n]* -> skip;

