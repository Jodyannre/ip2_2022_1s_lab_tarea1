parser grammar Nparser;


options {
  tokenVocab = Nlexer;
  language = Go;
}

@header {
  
}

@members{

}

inicio : expresion	#prod_inicio
;
/* 
inicio : PUBLICO MAIN DOSPUNTOS ID CORCHETE_IZQ ID PAR_IZQ STRING LLAVE_IZQ LLAVE_DER PAR_DER DOSPUNTOS ID CORCHETE_IZQ instrucciones CORCHETE_DER CORCHETE_DER 
;

instrucciones : instrucciones instruccion
				      |instruccion;
*/

/* 
instruccion : declaracion 
			      |asignacion 
			      |funcion_imprimir 
			      |control_if 
			      |expresion
;

declaracion : DECLARAR ID tipo IGUAL MENOR expresion MAYOR PUNTOCOMA
;

asignacion : ID IGUAL MENOR expresion MAYOR PUNTOCOMA
;
*/

expresion : RESTA expresion	#unario
			  //|expresion operador_logico expresion
			  //|expresion operador_comparacion expresion
        		|op_left=expresion operador_aritmetico op_right=expresion	#op_arit
				|NUMERO		#prod_numero  
				|DECIMAL 	#prod_decimal
			  //|CADENA
			  //|TRUE
			  //|FALSE
			  //|ID
			  //|PAR_DER expresion PAR_IZQ
;

/* 
tipo  : INTEGER
		  |REAL
		  |BOOLEAN
		  |STRING
;
*/
/* 
operador_comparacion : MAYOR
						| MAYOR IGUAL
						| IGUAL
;
						
operador_logico : AND
						    |OR
;
*/
				
operador_aritmetico : MULTIPLICACION	
                    |DIVISION	
                    |MODULO	
                    |SUMA	
                    |RESTA	
;
				
/*

//ESTRUCTURA DE CONTROL IF, IF ENTONCES, IF ENTONCES 

control_if : control_if IF ENTONCES DOSPUNTOS PAR_DER expresion PAR_IZQ CORCHETE_DER instruccion CORCHETE_IZQ PUNTOCOMA
 | control_if IF ENTONCES DOSPUNTOS PAR_DER expresion PAR_IZQ CORCHETE_DER instruccion CORCHETE_IZQ control_ENTONCES
 | IF DOSPUNTOS PAR_DER expresion PAR_IZQ CORCHETE_DER instruccion CORCHETE_IZQ PUNTOCOMA
 | IF DOSPUNTOS PAR_DER expresion PAR_IZQ CORCHETE_DER instruccion CORCHETE_IZQ control_ENTONCES
 ;

control_ENTONCES : ENTONCES CORCHETE_DER instruccion CORCHETE_IZQ PUNTOCOMA
;


//FUNCIONES NATIVAS -- SOLO PRINT

funcion_imprimir : SENTENCIA PUNTO CONSOLA PAR_DER instruccion PAR_IZQ PUNTOCOMA
;
 */



