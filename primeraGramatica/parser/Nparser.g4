parser grammar Nparser;


options {
  tokenVocab = Nlexer;
  language = Go;
}

@header {
  import "primeraGramatica/Analizador/Ast"
  import "primeraGramatica/Analizador/Expresiones"
  import "primeraGramatica/Analizador/Instrucciones"
  import "github.com/colegno/arraylist"
}

@members{

}

inicio returns[*arraylist.List lista] 
				: instrucciones	
				{
					$lista = $instrucciones.list
					//fmt.Println($lista)
				}
;
/* 
inicio : PUBLICO MAIN DOSPUNTOS ID CORCHETE_IZQ ID PAR_IZQ STRING LLAVE_IZQ LLAVE_DER PAR_DER DOSPUNTOS ID CORCHETE_IZQ instrucciones CORCHETE_DER CORCHETE_DER 
;
*/

instrucciones returns [*arraylist.List list]
			@init{
				$list =  arraylist.New()
			}
			: e += instruccion*  {
				listInt := localctx.(*InstruccionesContext).GetE()
						for _, e := range listInt {
						$list.Add(e.GetEx())
					}
				//fmt.Printf("tipo %T",localctx.(*InstruccionesContext).GetE())
			}
;

instruccion returns[interface{} ex] 
				:expresion 			{$ex = $expresion.ex}
				|control_IF_simple	{$ex = $control_IF_simple.ex} 
			    |asignacion 		{$ex = $asignacion.ex} 	
				|declaracion		{$ex = $declaracion.ex}
			    |funcion_imprimir 	{$ex = $funcion_imprimir.ex}
			    
			     
;

declaracion returns[Ast.Instruccion ex] 	
				:DECLARAR ID tipo IGUAL MENOR expresion 
				{
					$ex = Expresiones.NewDeclaracion($ID.text,$tipo.ex,$expresion.ex)
				}
				MAYOR PUNTOCOMA
				|DECLARAR ID tipo 
				{
					$ex = Expresiones.NewDeclaracion($ID.text,$tipo.ex,nil)
				}
				PUNTOCOMA
;


asignacion returns[Ast.Instruccion ex]
				: <assoc=right> ID IGUAL MENOR expresion 
					{
						//linea = strconv.Atoi($ID.line)
						//columna = strconv.Atoi($ID.pos)
						$ex = Expresiones.NewAsignacion($ID.text, $expresion.ex)
					}
				MAYOR PUNTOCOMA
;


expresion returns[Ast.Expresion ex] 
				:RESTA op_left=expresion
						{
							operador := $RESTA.text
							$ex = Expresiones.NewOperation($op_left.ex,operador,nil,true)

						}	
			  	|op_left=expresion operador_logico op_right=expresion
				  		{
							operador := $operador_logico.operador
							$ex = Expresiones.NewOperation($op_left.ex,operador,$op_right.ex,false)

						}	
			  	|op_left=expresion operador_comparacion op_right=expresion
				  		{
							operador := $operador_comparacion.operador
							$ex = Expresiones.NewOperation($op_left.ex,operador,$op_right.ex,false)

						}	
        		|op_left=expresion operador_aritmetico op_right=expresion	
						{
							operador := $operador_aritmetico.operador
							$ex = Expresiones.NewOperation($op_left.ex,operador,$op_right.ex,false)

						}	
				|NUMERO	{
							num,err := strconv.Atoi($NUMERO.text)
							if err!= nil{
								fmt.Println("Hay un error en el get número")
								fmt.Println(err)
							}
							$ex = Expresiones.NewPrimitivo (num,Ast.INTEGER)
						}	  
				|DECIMAL{
							num,err := strconv.ParseFloat($DECIMAL.text,64)
							if err!= nil{
								fmt.Println("Hay un error en el get número")
								fmt.Println(err)
							}
							$ex = Expresiones.NewPrimitivo (num,Ast.REAL)
						}	
			  	|CADENA {
							texto := $CADENA.text
							texto = texto[1:len(texto)-1]
							$ex = Expresiones.NewPrimitivo (texto,Ast.STRING)
						}	
			  	|TRUE	{
							valor := true
							$ex = Expresiones.NewPrimitivo (valor,Ast.BOOLEAN)
						}	
			  	|FALSE	{
							valor := false
							$ex = Expresiones.NewPrimitivo (valor,Ast.BOOLEAN)
						}	
			  	|ID		{
							id := $ID.text
							$ex = Expresiones.NewIdentificador (id,Ast.IDENTIFICADOR)
						}	
			  	|PAR_IZQ expresion PAR_DER
				  		{
							$ex = $expresion.ex	  
						}	
;


tipo returns[Ast.TipoDato ex]  	
		:INTEGER 	{$ex = Ast.INTEGER}
		|REAL		{$ex = Ast.REAL}
		|BOOLEAN	{$ex = Ast.BOOLEAN}
		|STRING		{$ex = Ast.STRING}
;

 
operador_comparacion returns[string operador] 
					:MAYOR			{$operador = $MAYOR.text}
					|MAYOR IGUAL	{$operador = $MAYOR.text + $IGUAL.text}
					//|IGUAL			{$operador = $IGUAL.text}
;
						
operador_logico 	returns[string operador] 
					:AND 			{$operador = $AND.text}
					|OR 			{$operador = $OR.text}
;

				
operador_aritmetico returns[string operador] 
					:MULTIPLICACION	{$operador = $MULTIPLICACION.text}
                    |DIVISION		{$operador = $DIVISION.text}	
                    |MODULO			{$operador = $MODULO.text}
                    |SUMA			{$operador = $SUMA.text}
                    |RESTA 			{$operador = $RESTA.text}
;
				


//ESTRUCTURA DE CONTROL IF, IF ENTONCES, IF ENTONCES 


control_IF_simple returns[Ast.Instruccion ex]
	:IF DOSPUNTOS PAR_IZQ expresion PAR_DER CORCHETE_IZQ instrucciones CORCHETE_DER control_else PUNTOCOMA
	{
		lista_entonces := arraylist.New()
		lista_entonces.Add($control_else.ex)
		$ex = Instrucciones.NewIF($expresion.ex,$instrucciones.list,lista_entonces,Ast.IF)
	}	
	|IF DOSPUNTOS PAR_IZQ expresion PAR_DER CORCHETE_IZQ
	instrucciones CORCHETE_DER PUNTOCOMA
	{
		lista_null := arraylist.New()
		$ex = Instrucciones.NewIF($expresion.ex,$instrucciones.list,lista_null,Ast.IF)
	}
	|IF DOSPUNTOS PAR_IZQ expresion PAR_DER CORCHETE_IZQ instrucciones CORCHETE_DER control_IF_entonces PUNTOCOMA
	{
		lista_entonces := $control_IF_entonces.list
		$ex = Instrucciones.NewIF($expresion.ex,$instrucciones.list,lista_entonces,Ast.IF)		
	}
	|IF DOSPUNTOS PAR_IZQ expresion PAR_DER CORCHETE_IZQ instrucciones CORCHETE_DER control_IF_entonces control_else PUNTOCOMA
		{
		lista_entonces := $control_IF_entonces.list
		lista_entonces.Add($control_else.ex)
		$ex = Instrucciones.NewIF($expresion.ex,$instrucciones.list,lista_entonces,Ast.IF)		
	}
;

control_IF_entonces returns [*arraylist.List list]
			@init{
				$list =  arraylist.New()
			}
			: e += control_if_entonces*  {
				listInt := localctx.(*Control_IF_entoncesContext).GetE()
						for _, e := range listInt {
						$list.Add(e.GetEx())
					}
				//fmt.Printf("tipo %T",localctx.(*Control_IF_entoncesContext).GetE())
			}
;


control_if_entonces returns[Ast.Instruccion ex]
	:ENTONCES IF DOSPUNTOS PAR_IZQ expresion PAR_DER CORCHETE_IZQ instrucciones CORCHETE_DER 
	{
		lista_null := arraylist.New()
		$ex = Instrucciones.NewIF($expresion.ex,$instrucciones.list,lista_null,Ast.IF_ENTONCES)
	}
;

control_else returns[Ast.Instruccion ex] 	
	: ENTONCES CORCHETE_IZQ instrucciones
	{
		expresionNull := Expresiones.NewPrimitivo(nil,Ast.NULL)
		lista_null := arraylist.New()
		$ex = Instrucciones.NewIF(expresionNull,$instrucciones.list,lista_null,Ast.ENTONCES)
	}
	CORCHETE_DER
;


//FUNCIONES NATIVAS -- SOLO PRINT

funcion_imprimir returns[Ast.Instruccion ex] 
			:SENTENCIA PUNTO CONSOLA PAR_IZQ expresion 
			{
				$ex = Instrucciones.NewImprimir($expresion.ex)
			}
			PAR_DER PUNTOCOMA
;




