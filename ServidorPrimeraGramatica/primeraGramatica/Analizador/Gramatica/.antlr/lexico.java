// Generated from c:\Projects\Go\src\primeraGramatica\Analizador\Gramatica\lexico.g by ANTLR 4.8
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class lexico extends Lexer {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		SENTENCIA=1, AND=2, OR=3, IF=4, ENTONCES=5, CONSOLA=6, STRING=7, INTEGER=8, 
		REAL=9, BOOLEAN=10, TRUE=11, FALSE=12, PUBLICO=13, PRIVADO=14, MAIN=15, 
		PRINCIPAL=16, METODO=17, CLASE=18, NUMERO=19, DECIMAL=20, ID=21, DOSPUNTOS=22, 
		PUNTO=23, PUNTOCOMA=24, MAYOR_I=25, MAYOR=26, IGUAL=27, MODULO=28, MULTIPLICACION=29, 
		DIVISION=30, RESTA=31, SUMA=32, PAR_IZQ=33, PAR_DER=34, LLAVE_IZQ=35, 
		LLAVE_DER=36, CORCHETE_IZQ=37, CORCHETE_DER=38, CADENA=39, WHITESPACE=40, 
		COMMENT=41, LINE_COMMENT=42;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"SENTENCIA", "AND", "OR", "IF", "ENTONCES", "CONSOLA", "STRING", "INTEGER", 
			"REAL", "BOOLEAN", "TRUE", "FALSE", "PUBLICO", "PRIVADO", "MAIN", "PRINCIPAL", 
			"METODO", "CLASE", "NUMERO", "DECIMAL", "ID", "DOSPUNTOS", "PUNTO", "PUNTOCOMA", 
			"MAYOR_I", "MAYOR", "IGUAL", "MODULO", "MULTIPLICACION", "DIVISION", 
			"RESTA", "SUMA", "PAR_IZQ", "PAR_DER", "LLAVE_IZQ", "LLAVE_DER", "CORCHETE_IZQ", 
			"CORCHETE_DER", "CADENA", "WHITESPACE", "COMMENT", "LINE_COMMENT"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'SENTENCIA'", "'AND'", "'OR'", "'IF'", "'ENTONCES'", "'CONSOLA'", 
			"'STRING'", "'INTEGER'", "'REAL'", "'BOOLEAN'", "'TRUE'", "'FALSE'", 
			"'PUBLICO'", "'PRIVADO'", "'MAIN'", "'PRINCIPAL'", "'METODO'", "'CLASE'", 
			null, null, null, "':'", "'.'", "';'", "'>='", "'>'", "'='", "'%'", "'*'", 
			"'/'", "'-'", "'+'", "'('", "')'", "'{'", "'}'", "'['", "']'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "SENTENCIA", "AND", "OR", "IF", "ENTONCES", "CONSOLA", "STRING", 
			"INTEGER", "REAL", "BOOLEAN", "TRUE", "FALSE", "PUBLICO", "PRIVADO", 
			"MAIN", "PRINCIPAL", "METODO", "CLASE", "NUMERO", "DECIMAL", "ID", "DOSPUNTOS", 
			"PUNTO", "PUNTOCOMA", "MAYOR_I", "MAYOR", "IGUAL", "MODULO", "MULTIPLICACION", 
			"DIVISION", "RESTA", "SUMA", "PAR_IZQ", "PAR_DER", "LLAVE_IZQ", "LLAVE_DER", 
			"CORCHETE_IZQ", "CORCHETE_DER", "CADENA", "WHITESPACE", "COMMENT", "LINE_COMMENT"
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


	public lexico(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "lexico.g"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2,\u0134\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\3"+
		"\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\5"+
		"\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3"+
		"\7\3\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\n"+
		"\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3"+
		"\f\3\f\3\r\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16"+
		"\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\21"+
		"\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\3\23\3\23\3\23\3\23\3\23\3\23\3\24\6\24\u00d1\n\24\r\24\16"+
		"\24\u00d2\3\25\6\25\u00d6\n\25\r\25\16\25\u00d7\3\25\3\25\6\25\u00dc\n"+
		"\25\r\25\16\25\u00dd\3\26\5\26\u00e1\n\26\3\26\7\26\u00e4\n\26\f\26\16"+
		"\26\u00e7\13\26\3\27\3\27\3\30\3\30\3\31\3\31\3\32\3\32\3\32\3\33\3\33"+
		"\3\34\3\34\3\35\3\35\3\36\3\36\3\37\3\37\3 \3 \3!\3!\3\"\3\"\3#\3#\3$"+
		"\3$\3%\3%\3&\3&\3\'\3\'\3(\3(\7(\u010e\n(\f(\16(\u0111\13(\3(\3(\3)\6"+
		")\u0116\n)\r)\16)\u0117\3)\3)\3*\3*\3*\3*\7*\u0120\n*\f*\16*\u0123\13"+
		"*\3*\3*\3*\3*\3*\3+\3+\3+\3+\7+\u012e\n+\f+\16+\u0131\13+\3+\3+\4\u010f"+
		"\u0121\2,\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33"+
		"\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67"+
		"\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,\3\2\7\3\2\62;\5\2\62;C\\c|\6\2"+
		"\62;C\\aac|\6\2\13\f\17\17\"\"^^\4\2\f\f\17\17\2\u013b\2\3\3\2\2\2\2\5"+
		"\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2"+
		"\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33"+
		"\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2"+
		"\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2"+
		"\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2"+
		"\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K"+
		"\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\3W\3\2"+
		"\2\2\5a\3\2\2\2\7e\3\2\2\2\th\3\2\2\2\13k\3\2\2\2\rt\3\2\2\2\17|\3\2\2"+
		"\2\21\u0083\3\2\2\2\23\u008b\3\2\2\2\25\u0090\3\2\2\2\27\u0098\3\2\2\2"+
		"\31\u009d\3\2\2\2\33\u00a3\3\2\2\2\35\u00ab\3\2\2\2\37\u00b3\3\2\2\2!"+
		"\u00b8\3\2\2\2#\u00c2\3\2\2\2%\u00c9\3\2\2\2\'\u00d0\3\2\2\2)\u00d5\3"+
		"\2\2\2+\u00e0\3\2\2\2-\u00e8\3\2\2\2/\u00ea\3\2\2\2\61\u00ec\3\2\2\2\63"+
		"\u00ee\3\2\2\2\65\u00f1\3\2\2\2\67\u00f3\3\2\2\29\u00f5\3\2\2\2;\u00f7"+
		"\3\2\2\2=\u00f9\3\2\2\2?\u00fb\3\2\2\2A\u00fd\3\2\2\2C\u00ff\3\2\2\2E"+
		"\u0101\3\2\2\2G\u0103\3\2\2\2I\u0105\3\2\2\2K\u0107\3\2\2\2M\u0109\3\2"+
		"\2\2O\u010b\3\2\2\2Q\u0115\3\2\2\2S\u011b\3\2\2\2U\u0129\3\2\2\2WX\7U"+
		"\2\2XY\7G\2\2YZ\7P\2\2Z[\7V\2\2[\\\7G\2\2\\]\7P\2\2]^\7E\2\2^_\7K\2\2"+
		"_`\7C\2\2`\4\3\2\2\2ab\7C\2\2bc\7P\2\2cd\7F\2\2d\6\3\2\2\2ef\7Q\2\2fg"+
		"\7T\2\2g\b\3\2\2\2hi\7K\2\2ij\7H\2\2j\n\3\2\2\2kl\7G\2\2lm\7P\2\2mn\7"+
		"V\2\2no\7Q\2\2op\7P\2\2pq\7E\2\2qr\7G\2\2rs\7U\2\2s\f\3\2\2\2tu\7E\2\2"+
		"uv\7Q\2\2vw\7P\2\2wx\7U\2\2xy\7Q\2\2yz\7N\2\2z{\7C\2\2{\16\3\2\2\2|}\7"+
		"U\2\2}~\7V\2\2~\177\7T\2\2\177\u0080\7K\2\2\u0080\u0081\7P\2\2\u0081\u0082"+
		"\7I\2\2\u0082\20\3\2\2\2\u0083\u0084\7K\2\2\u0084\u0085\7P\2\2\u0085\u0086"+
		"\7V\2\2\u0086\u0087\7G\2\2\u0087\u0088\7I\2\2\u0088\u0089\7G\2\2\u0089"+
		"\u008a\7T\2\2\u008a\22\3\2\2\2\u008b\u008c\7T\2\2\u008c\u008d\7G\2\2\u008d"+
		"\u008e\7C\2\2\u008e\u008f\7N\2\2\u008f\24\3\2\2\2\u0090\u0091\7D\2\2\u0091"+
		"\u0092\7Q\2\2\u0092\u0093\7Q\2\2\u0093\u0094\7N\2\2\u0094\u0095\7G\2\2"+
		"\u0095\u0096\7C\2\2\u0096\u0097\7P\2\2\u0097\26\3\2\2\2\u0098\u0099\7"+
		"V\2\2\u0099\u009a\7T\2\2\u009a\u009b\7W\2\2\u009b\u009c\7G\2\2\u009c\30"+
		"\3\2\2\2\u009d\u009e\7H\2\2\u009e\u009f\7C\2\2\u009f\u00a0\7N\2\2\u00a0"+
		"\u00a1\7U\2\2\u00a1\u00a2\7G\2\2\u00a2\32\3\2\2\2\u00a3\u00a4\7R\2\2\u00a4"+
		"\u00a5\7W\2\2\u00a5\u00a6\7D\2\2\u00a6\u00a7\7N\2\2\u00a7\u00a8\7K\2\2"+
		"\u00a8\u00a9\7E\2\2\u00a9\u00aa\7Q\2\2\u00aa\34\3\2\2\2\u00ab\u00ac\7"+
		"R\2\2\u00ac\u00ad\7T\2\2\u00ad\u00ae\7K\2\2\u00ae\u00af\7X\2\2\u00af\u00b0"+
		"\7C\2\2\u00b0\u00b1\7F\2\2\u00b1\u00b2\7Q\2\2\u00b2\36\3\2\2\2\u00b3\u00b4"+
		"\7O\2\2\u00b4\u00b5\7C\2\2\u00b5\u00b6\7K\2\2\u00b6\u00b7\7P\2\2\u00b7"+
		" \3\2\2\2\u00b8\u00b9\7R\2\2\u00b9\u00ba\7T\2\2\u00ba\u00bb\7K\2\2\u00bb"+
		"\u00bc\7P\2\2\u00bc\u00bd\7E\2\2\u00bd\u00be\7K\2\2\u00be\u00bf\7R\2\2"+
		"\u00bf\u00c0\7C\2\2\u00c0\u00c1\7N\2\2\u00c1\"\3\2\2\2\u00c2\u00c3\7O"+
		"\2\2\u00c3\u00c4\7G\2\2\u00c4\u00c5\7V\2\2\u00c5\u00c6\7Q\2\2\u00c6\u00c7"+
		"\7F\2\2\u00c7\u00c8\7Q\2\2\u00c8$\3\2\2\2\u00c9\u00ca\7E\2\2\u00ca\u00cb"+
		"\7N\2\2\u00cb\u00cc\7C\2\2\u00cc\u00cd\7U\2\2\u00cd\u00ce\7G\2\2\u00ce"+
		"&\3\2\2\2\u00cf\u00d1\t\2\2\2\u00d0\u00cf\3\2\2\2\u00d1\u00d2\3\2\2\2"+
		"\u00d2\u00d0\3\2\2\2\u00d2\u00d3\3\2\2\2\u00d3(\3\2\2\2\u00d4\u00d6\t"+
		"\2\2\2\u00d5\u00d4\3\2\2\2\u00d6\u00d7\3\2\2\2\u00d7\u00d5\3\2\2\2\u00d7"+
		"\u00d8\3\2\2\2\u00d8\u00d9\3\2\2\2\u00d9\u00db\7\60\2\2\u00da\u00dc\t"+
		"\2\2\2\u00db\u00da\3\2\2\2\u00dc\u00dd\3\2\2\2\u00dd\u00db\3\2\2\2\u00dd"+
		"\u00de\3\2\2\2\u00de*\3\2\2\2\u00df\u00e1\t\3\2\2\u00e0\u00df\3\2\2\2"+
		"\u00e1\u00e5\3\2\2\2\u00e2\u00e4\t\4\2\2\u00e3\u00e2\3\2\2\2\u00e4\u00e7"+
		"\3\2\2\2\u00e5\u00e3\3\2\2\2\u00e5\u00e6\3\2\2\2\u00e6,\3\2\2\2\u00e7"+
		"\u00e5\3\2\2\2\u00e8\u00e9\7<\2\2\u00e9.\3\2\2\2\u00ea\u00eb\7\60\2\2"+
		"\u00eb\60\3\2\2\2\u00ec\u00ed\7=\2\2\u00ed\62\3\2\2\2\u00ee\u00ef\7@\2"+
		"\2\u00ef\u00f0\7?\2\2\u00f0\64\3\2\2\2\u00f1\u00f2\7@\2\2\u00f2\66\3\2"+
		"\2\2\u00f3\u00f4\7?\2\2\u00f48\3\2\2\2\u00f5\u00f6\7\'\2\2\u00f6:\3\2"+
		"\2\2\u00f7\u00f8\7,\2\2\u00f8<\3\2\2\2\u00f9\u00fa\7\61\2\2\u00fa>\3\2"+
		"\2\2\u00fb\u00fc\7/\2\2\u00fc@\3\2\2\2\u00fd\u00fe\7-\2\2\u00feB\3\2\2"+
		"\2\u00ff\u0100\7*\2\2\u0100D\3\2\2\2\u0101\u0102\7+\2\2\u0102F\3\2\2\2"+
		"\u0103\u0104\7}\2\2\u0104H\3\2\2\2\u0105\u0106\7\177\2\2\u0106J\3\2\2"+
		"\2\u0107\u0108\7]\2\2\u0108L\3\2\2\2\u0109\u010a\7_\2\2\u010aN\3\2\2\2"+
		"\u010b\u010f\7$\2\2\u010c\u010e\13\2\2\2\u010d\u010c\3\2\2\2\u010e\u0111"+
		"\3\2\2\2\u010f\u0110\3\2\2\2\u010f\u010d\3\2\2\2\u0110\u0112\3\2\2\2\u0111"+
		"\u010f\3\2\2\2\u0112\u0113\7$\2\2\u0113P\3\2\2\2\u0114\u0116\t\5\2\2\u0115"+
		"\u0114\3\2\2\2\u0116\u0117\3\2\2\2\u0117\u0115\3\2\2\2\u0117\u0118\3\2"+
		"\2\2\u0118\u0119\3\2\2\2\u0119\u011a\b)\2\2\u011aR\3\2\2\2\u011b\u011c"+
		"\7\61\2\2\u011c\u011d\7,\2\2\u011d\u0121\3\2\2\2\u011e\u0120\13\2\2\2"+
		"\u011f\u011e\3\2\2\2\u0120\u0123\3\2\2\2\u0121\u0122\3\2\2\2\u0121\u011f"+
		"\3\2\2\2\u0122\u0124\3\2\2\2\u0123\u0121\3\2\2\2\u0124\u0125\7,\2\2\u0125"+
		"\u0126\7\61\2\2\u0126\u0127\3\2\2\2\u0127\u0128\b*\2\2\u0128T\3\2\2\2"+
		"\u0129\u012a\7\61\2\2\u012a\u012b\7\61\2\2\u012b\u012f\3\2\2\2\u012c\u012e"+
		"\n\6\2\2\u012d\u012c\3\2\2\2\u012e\u0131\3\2\2\2\u012f\u012d\3\2\2\2\u012f"+
		"\u0130\3\2\2\2\u0130\u0132\3\2\2\2\u0131\u012f\3\2\2\2\u0132\u0133\b+"+
		"\2\2\u0133V\3\2\2\2\r\2\u00d2\u00d7\u00dd\u00e0\u00e3\u00e5\u010f\u0117"+
		"\u0121\u012f\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}