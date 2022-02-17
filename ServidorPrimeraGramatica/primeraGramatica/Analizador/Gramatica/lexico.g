lexer grammar lexico;


SENTENCIA       : 'SENTENCIA';
AND             : 'AND';
OR              : 'OR';
IF              : 'IF';
ENTONCES        : 'ENTONCES';
CONSOLA         : 'CONSOLA';
STRING          : 'STRING';
INTEGER         : 'INTEGER';
REAL            : 'REAL';
BOOLEAN         : 'BOOLEAN';
TRUE            : 'TRUE';
FALSE           : 'FALSE';
PUBLICO         : 'PUBLICO';
PRIVADO         : 'PRIVADO';
MAIN            : 'MAIN';
PRINCIPAL       : 'PRINCIPAL';
METODO          : 'METODO';
CLASE           : 'CLASE';

NUMERO          : [0-9]+;
DECIMAL         : [0-9]+'.'[0-9]+;
ID              : ([a-zA-Z]|[0-9])([a-zA-Z]|[0-9]|'_')*;


DOSPUNTOS       : ':';
PUNTO           : '.';
PUNTOCOMA       : ';';
MAYOR_I         : '>=';
MAYOR           : '>';
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

