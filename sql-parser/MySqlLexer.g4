lexer grammar MySqlLexer;

channels { MYSQLCOMMENT, ERRORCHANNEL }

// SKIP

SPACE:                               [ \t\r\n]+    -> channel(HIDDEN);
SPEC_MYSQL_COMMENT:                  '/*!' .+? '*/' -> channel(MYSQLCOMMENT);
COMMENT_INPUT:                       '/*' .*? '*/' -> channel(HIDDEN);
LINE_COMMENT:                        (
                                       ('-- ' | '#') ~[\r\n]* ('\r'? '\n' | EOF)
                                       | '--' ('\r'? '\n' | EOF)
                                     ) -> channel(HIDDEN);

ALL:                                 A L L;
AND:                                 A N D;
AS:                                  A S;
ASC:                                 A S C;
BETWEEN:                             B E T W E E N;
BY:                                  B Y;
CASE:                                C A S E;
CROSS:                               C R O S S;
DESC:                                D E S C;
DISTINCT:                            D I S T I N C T;
ELSE:                                E L S E;
EXISTS:                              E X I S T S;
FALSE:                               F A L S E;
FROM:                                F R O M;
GROUP:                               G R O U P;
HAVING:                              H A V I N G;
IN:                                  I N;
INNER:                               I N N E R;
INTO:                                I N T O;
IS:                                  I S;
JOIN:                                J O I N;
LEFT:                                L E F T;
LIKE:                                L I K E;
LIMIT:                               L I M I T;
NOT:                                 N O T;
NULL_LITERAL:                        N U L L;
ON:                                  O N;
OR:                                  O R;
ORDER:                               O R D E R;
OUTER:                               O U T E R;
REGEXP:                              R E G E X P;
RIGHT:                               R I G H T;
RLIKE:                               R L I K E;
SELECT:                              S E L E C T;
THEN:                                T H E N;
TRUE:                                T R U E;
UNION:                               U N I O N;
USING:                               U S I N G;
VALUES:                              V A L U E S;
WHEN:                                W H E N;
WHERE:                               W H E R E;
DELETE:                              D E L E T E;
INSERT:                              I N S E R T;
SET:                                 S E T;
UPDATE:                              U P D A T E;
VALUE:                               V A L U E;
DEFAULT:                             D E F A U L T;
IGNORE:                              I G N O R E;

//// Common Keywords, but can be ID
//
ANY:                                 A N Y;
END:                                 E N D;
OFFSET:                              O F F S E T;
SOME:                                S O M E;
UNKNOWN:                             U N K N O W N;

//// Operators. Arithmetics

STAR:                                '*';
DIVIDE:                              '/';
MODULE:                              '%';
PLUS:                                '+';
MINUSMINUS:                          '--';
MINUS:                               '-';
DIV:                                 D I V;
MOD:                                 M O D;

//// Operators. Comparation

EQUAL_SYMBOL:                        '=';
GREATER_SYMBOL:                      '>';
LESS_SYMBOL:                         '<';
EXCLAMATION_SYMBOL:                  '!';

//// Operators. Bit

BIT_NOT_OP:                          '~';
BIT_OR_OP:                           '|';
BIT_AND_OP:                          '&';
BIT_XOR_OP:                          '^';

//// Constructors symbols

DOT:                                 '.';
LR_BRACKET:                          '(';
RR_BRACKET:                          ')';
COMMA:                               ',';

//// Literal Primitives

STRING_LITERAL:                      DQUOTA_STRING | SQUOTA_STRING;
DECIMAL_LITERAL:                     DEC_DIGIT+;

REAL_LITERAL:                        (DEC_DIGIT+)? '.' DEC_DIGIT+;
NULL_SPEC_LITERAL:                   '\\' 'N';

DOT_ID:                              '.' ID_LITERAL;

//// Identifiers

ID:                                  ID_LITERAL | BQUOTA_STRING;

fragment ID_LITERAL:                 [A-Za-z_$0-9]*?[A-Za-z_$]+?[A-Za-z_$0-9]*;
fragment DQUOTA_STRING:              '"' ( '\\'. | '""' | ~('"'| '\\') )* '"';
fragment SQUOTA_STRING:              '\'' ('\\'. | '\'\'' | ~('\'' | '\\'))* '\'';
fragment BQUOTA_STRING:              '`' ( '\\'. | '``' | ~('`'|'\\'))* '`';
fragment DEC_DIGIT:                  [0-9];

fragment A : [aA];
fragment B : [bB];
fragment C : [cC];
fragment D : [dD];
fragment E : [eE];
fragment F : [fF];
fragment G : [gG];
fragment H : [hH];
fragment I : [iI];
fragment J : [jJ];
fragment K : [kK];
fragment L : [lL];
fragment M : [mM];
fragment N : [nN];
fragment O : [oO];
fragment P : [pP];
fragment Q : [qQ];
fragment R : [rR];
fragment S : [sS];
fragment T : [tT];
fragment U : [uU];
fragment V : [vV];
fragment W : [wW];
fragment X : [xX];
fragment Y : [yY];
fragment Z : [zZ];

// Last tokens must generate Errors
ERROR_RECONGNIGION:                  .    -> channel(ERRORCHANNEL);
