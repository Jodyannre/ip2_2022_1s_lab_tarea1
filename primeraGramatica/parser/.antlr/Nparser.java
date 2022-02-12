// Generated from c:\Projects\Go\src\holaMundo\parser\Nparser.g4 by ANTLR 4.8

  

import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class Nparser extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		SENTENCIA=1, AND=2, OR=3, IF=4, ENTONCES=5, CONSOLA=6, STRING=7, INTEGER=8, 
		REAL=9, BOOLEAN=10, TRUE=11, FALSE=12, PUBLICO=13, PRIVADO=14, MAIN=15, 
		PRINCIPAL=16, METODO=17, CLASE=18, DECLARAR=19, NUMERO=20, DECIMAL=21, 
		ID=22, DOSPUNTOS=23, PUNTO=24, PUNTOCOMA=25, MAYOR_I=26, MAYOR=27, MENOR=28, 
		IGUAL=29, MODULO=30, MULTIPLICACION=31, DIVISION=32, RESTA=33, SUMA=34, 
		PAR_IZQ=35, PAR_DER=36, LLAVE_IZQ=37, LLAVE_DER=38, CORCHETE_IZQ=39, CORCHETE_DER=40, 
		CADENA=41, WHITESPACE=42, COMMENT=43, LINE_COMMENT=44;
	public static final int
		RULE_inicio = 0, RULE_expresion = 1, RULE_operador_aritmetico = 2;
	private static String[] makeRuleNames() {
		return new String[] {
			"inicio", "expresion", "operador_aritmetico"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, "':'", 
			"'.'", "';'", "'>='", "'>'", "'<'", "'='", "'%'", "'*'", "'/'", "'-'", 
			"'+'", "'('", "')'", "'{'", "'}'", "'['", "']'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "SENTENCIA", "AND", "OR", "IF", "ENTONCES", "CONSOLA", "STRING", 
			"INTEGER", "REAL", "BOOLEAN", "TRUE", "FALSE", "PUBLICO", "PRIVADO", 
			"MAIN", "PRINCIPAL", "METODO", "CLASE", "DECLARAR", "NUMERO", "DECIMAL", 
			"ID", "DOSPUNTOS", "PUNTO", "PUNTOCOMA", "MAYOR_I", "MAYOR", "MENOR", 
			"IGUAL", "MODULO", "MULTIPLICACION", "DIVISION", "RESTA", "SUMA", "PAR_IZQ", 
			"PAR_DER", "LLAVE_IZQ", "LLAVE_DER", "CORCHETE_IZQ", "CORCHETE_DER", 
			"CADENA", "WHITESPACE", "COMMENT", "LINE_COMMENT"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Nparser.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }




	public Nparser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class InicioContext extends ParserRuleContext {
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public InicioContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_inicio; }
	}

	public final InicioContext inicio() throws RecognitionException {
		InicioContext _localctx = new InicioContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_inicio);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(6);
			expresion(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExpresionContext extends ParserRuleContext {
		public ExpresionContext op_left;
		public ExpresionContext op_right;
		public TerminalNode RESTA() { return getToken(Nparser.RESTA, 0); }
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public TerminalNode NUMERO() { return getToken(Nparser.NUMERO, 0); }
		public TerminalNode DECIMAL() { return getToken(Nparser.DECIMAL, 0); }
		public Operador_aritmeticoContext operador_aritmetico() {
			return getRuleContext(Operador_aritmeticoContext.class,0);
		}
		public ExpresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expresion; }
	}

	public final ExpresionContext expresion() throws RecognitionException {
		return expresion(0);
	}

	private ExpresionContext expresion(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpresionContext _localctx = new ExpresionContext(_ctx, _parentState);
		ExpresionContext _prevctx = _localctx;
		int _startState = 2;
		enterRecursionRule(_localctx, 2, RULE_expresion, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(13);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case RESTA:
				{
				setState(9);
				match(RESTA);
				setState(10);
				expresion(4);
				}
				break;
			case NUMERO:
				{
				setState(11);
				match(NUMERO);
				}
				break;
			case DECIMAL:
				{
				setState(12);
				match(DECIMAL);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(21);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,1,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ExpresionContext(_parentctx, _parentState);
					_localctx.op_left = _prevctx;
					_localctx.op_left = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_expresion);
					setState(15);
					if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
					setState(16);
					operador_aritmetico();
					setState(17);
					((ExpresionContext)_localctx).op_right = expresion(4);
					}
					} 
				}
				setState(23);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,1,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class Operador_aritmeticoContext extends ParserRuleContext {
		public TerminalNode MULTIPLICACION() { return getToken(Nparser.MULTIPLICACION, 0); }
		public TerminalNode DIVISION() { return getToken(Nparser.DIVISION, 0); }
		public TerminalNode MODULO() { return getToken(Nparser.MODULO, 0); }
		public TerminalNode SUMA() { return getToken(Nparser.SUMA, 0); }
		public TerminalNode RESTA() { return getToken(Nparser.RESTA, 0); }
		public Operador_aritmeticoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operador_aritmetico; }
	}

	public final Operador_aritmeticoContext operador_aritmetico() throws RecognitionException {
		Operador_aritmeticoContext _localctx = new Operador_aritmeticoContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_operador_aritmetico);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(24);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << MODULO) | (1L << MULTIPLICACION) | (1L << DIVISION) | (1L << RESTA) | (1L << SUMA))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 1:
			return expresion_sempred((ExpresionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expresion_sempred(ExpresionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 3);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3.\35\4\2\t\2\4\3\t"+
		"\3\4\4\t\4\3\2\3\2\3\3\3\3\3\3\3\3\3\3\5\3\20\n\3\3\3\3\3\3\3\3\3\7\3"+
		"\26\n\3\f\3\16\3\31\13\3\3\4\3\4\3\4\2\3\4\5\2\4\6\2\3\3\2 $\2\34\2\b"+
		"\3\2\2\2\4\17\3\2\2\2\6\32\3\2\2\2\b\t\5\4\3\2\t\3\3\2\2\2\n\13\b\3\1"+
		"\2\13\f\7#\2\2\f\20\5\4\3\6\r\20\7\26\2\2\16\20\7\27\2\2\17\n\3\2\2\2"+
		"\17\r\3\2\2\2\17\16\3\2\2\2\20\27\3\2\2\2\21\22\f\5\2\2\22\23\5\6\4\2"+
		"\23\24\5\4\3\6\24\26\3\2\2\2\25\21\3\2\2\2\26\31\3\2\2\2\27\25\3\2\2\2"+
		"\27\30\3\2\2\2\30\5\3\2\2\2\31\27\3\2\2\2\32\33\t\2\2\2\33\7\3\2\2\2\4"+
		"\17\27";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}