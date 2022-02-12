// Generated from c:\Projects\Go\src\holaMundo\parser\Nlexer.g4 by ANTLR 4.8
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class Nlexer extends Lexer {
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
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", 
			"O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "SENTENCIA", 
			"AND", "OR", "IF", "ENTONCES", "CONSOLA", "STRING", "INTEGER", "REAL", 
			"BOOLEAN", "TRUE", "FALSE", "PUBLICO", "PRIVADO", "MAIN", "PRINCIPAL", 
			"METODO", "CLASE", "DECLARAR", "NUMERO", "DECIMAL", "ID", "DOSPUNTOS", 
			"PUNTO", "PUNTOCOMA", "MAYOR_I", "MAYOR", "MENOR", "IGUAL", "MODULO", 
			"MULTIPLICACION", "DIVISION", "RESTA", "SUMA", "PAR_IZQ", "PAR_DER", 
			"LLAVE_IZQ", "LLAVE_DER", "CORCHETE_IZQ", "CORCHETE_DER", "CADENA", "WHITESPACE", 
			"COMMENT", "LINE_COMMENT"
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


	public Nlexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Nlexer.g4"; }

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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2.\u01ab\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\4D\tD\4E\tE\4F\tF\4G\tG\3\2\3\2\3"+
		"\3\3\3\3\4\3\4\3\5\3\5\3\6\3\6\3\7\3\7\3\b\3\b\3\t\3\t\3\n\3\n\3\13\3"+
		"\13\3\f\3\f\3\r\3\r\3\16\3\16\3\17\3\17\3\20\3\20\3\21\3\21\3\22\3\22"+
		"\3\23\3\23\3\24\3\24\3\25\3\25\3\26\3\26\3\27\3\27\3\30\3\30\3\31\3\31"+
		"\3\32\3\32\3\33\3\33\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34"+
		"\3\35\3\35\3\35\3\35\3\36\3\36\3\36\3\37\3\37\3\37\3 \3 \3 \3 \3 \3 \3"+
		" \3 \3 \3!\3!\3!\3!\3!\3!\3!\3!\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3#\3#\3#\3"+
		"#\3#\3#\3#\3#\3$\3$\3$\3$\3$\3%\3%\3%\3%\3%\3%\3%\3%\3&\3&\3&\3&\3&\3"+
		"\'\3\'\3\'\3\'\3\'\3\'\3(\3(\3(\3(\3(\3(\3(\3(\3)\3)\3)\3)\3)\3)\3)\3"+
		")\3*\3*\3*\3*\3*\3+\3+\3+\3+\3+\3+\3+\3+\3+\3+\3,\3,\3,\3,\3,\3,\3,\3"+
		"-\3-\3-\3-\3-\3-\3.\3.\3.\3.\3.\3.\3.\3.\3.\3/\6/\u0146\n/\r/\16/\u0147"+
		"\3\60\6\60\u014b\n\60\r\60\16\60\u014c\3\60\3\60\6\60\u0151\n\60\r\60"+
		"\16\60\u0152\3\61\5\61\u0156\n\61\3\61\7\61\u0159\n\61\f\61\16\61\u015c"+
		"\13\61\3\62\3\62\3\63\3\63\3\64\3\64\3\65\3\65\3\65\3\66\3\66\3\67\3\67"+
		"\38\38\39\39\3:\3:\3;\3;\3<\3<\3=\3=\3>\3>\3?\3?\3@\3@\3A\3A\3B\3B\3C"+
		"\3C\3D\3D\7D\u0185\nD\fD\16D\u0188\13D\3D\3D\3E\6E\u018d\nE\rE\16E\u018e"+
		"\3E\3E\3F\3F\3F\3F\7F\u0197\nF\fF\16F\u019a\13F\3F\3F\3F\3F\3F\3G\3G\3"+
		"G\3G\7G\u01a5\nG\fG\16G\u01a8\13G\3G\3G\4\u0186\u0198\2H\3\2\5\2\7\2\t"+
		"\2\13\2\r\2\17\2\21\2\23\2\25\2\27\2\31\2\33\2\35\2\37\2!\2#\2%\2\'\2"+
		")\2+\2-\2/\2\61\2\63\2\65\2\67\39\4;\5=\6?\7A\bC\tE\nG\13I\fK\rM\16O\17"+
		"Q\20S\21U\22W\23Y\24[\25]\26_\27a\30c\31e\32g\33i\34k\35m\36o\37q s!u"+
		"\"w#y${%}&\177\'\u0081(\u0083)\u0085*\u0087+\u0089,\u008b-\u008d.\3\2"+
		"!\4\2CCcc\4\2DDdd\4\2EEee\4\2FFff\4\2GGgg\4\2HHhh\4\2IIii\4\2JJjj\4\2"+
		"KKkk\4\2LLll\4\2MMmm\4\2NNnn\4\2OOoo\4\2PPpp\4\2QQqq\4\2RRrr\4\2SSss\4"+
		"\2TTtt\4\2UUuu\4\2VVvv\4\2WWww\4\2XXxx\4\2YYyy\4\2ZZzz\4\2[[{{\4\2\\\\"+
		"||\3\2\62;\5\2\62;C\\c|\6\2\62;C\\aac|\6\2\13\f\17\17\"\"^^\4\2\f\f\17"+
		"\17\2\u0198\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2"+
		"A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3"+
		"\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2"+
		"\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2"+
		"g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2\2m\3\2\2\2\2o\3\2\2\2\2q\3\2\2\2\2s\3"+
		"\2\2\2\2u\3\2\2\2\2w\3\2\2\2\2y\3\2\2\2\2{\3\2\2\2\2}\3\2\2\2\2\177\3"+
		"\2\2\2\2\u0081\3\2\2\2\2\u0083\3\2\2\2\2\u0085\3\2\2\2\2\u0087\3\2\2\2"+
		"\2\u0089\3\2\2\2\2\u008b\3\2\2\2\2\u008d\3\2\2\2\3\u008f\3\2\2\2\5\u0091"+
		"\3\2\2\2\7\u0093\3\2\2\2\t\u0095\3\2\2\2\13\u0097\3\2\2\2\r\u0099\3\2"+
		"\2\2\17\u009b\3\2\2\2\21\u009d\3\2\2\2\23\u009f\3\2\2\2\25\u00a1\3\2\2"+
		"\2\27\u00a3\3\2\2\2\31\u00a5\3\2\2\2\33\u00a7\3\2\2\2\35\u00a9\3\2\2\2"+
		"\37\u00ab\3\2\2\2!\u00ad\3\2\2\2#\u00af\3\2\2\2%\u00b1\3\2\2\2\'\u00b3"+
		"\3\2\2\2)\u00b5\3\2\2\2+\u00b7\3\2\2\2-\u00b9\3\2\2\2/\u00bb\3\2\2\2\61"+
		"\u00bd\3\2\2\2\63\u00bf\3\2\2\2\65\u00c1\3\2\2\2\67\u00c3\3\2\2\29\u00cd"+
		"\3\2\2\2;\u00d1\3\2\2\2=\u00d4\3\2\2\2?\u00d7\3\2\2\2A\u00e0\3\2\2\2C"+
		"\u00e8\3\2\2\2E\u00ef\3\2\2\2G\u00f7\3\2\2\2I\u00fc\3\2\2\2K\u0104\3\2"+
		"\2\2M\u0109\3\2\2\2O\u010f\3\2\2\2Q\u0117\3\2\2\2S\u011f\3\2\2\2U\u0124"+
		"\3\2\2\2W\u012e\3\2\2\2Y\u0135\3\2\2\2[\u013b\3\2\2\2]\u0145\3\2\2\2_"+
		"\u014a\3\2\2\2a\u0155\3\2\2\2c\u015d\3\2\2\2e\u015f\3\2\2\2g\u0161\3\2"+
		"\2\2i\u0163\3\2\2\2k\u0166\3\2\2\2m\u0168\3\2\2\2o\u016a\3\2\2\2q\u016c"+
		"\3\2\2\2s\u016e\3\2\2\2u\u0170\3\2\2\2w\u0172\3\2\2\2y\u0174\3\2\2\2{"+
		"\u0176\3\2\2\2}\u0178\3\2\2\2\177\u017a\3\2\2\2\u0081\u017c\3\2\2\2\u0083"+
		"\u017e\3\2\2\2\u0085\u0180\3\2\2\2\u0087\u0182\3\2\2\2\u0089\u018c\3\2"+
		"\2\2\u008b\u0192\3\2\2\2\u008d\u01a0\3\2\2\2\u008f\u0090\t\2\2\2\u0090"+
		"\4\3\2\2\2\u0091\u0092\t\3\2\2\u0092\6\3\2\2\2\u0093\u0094\t\4\2\2\u0094"+
		"\b\3\2\2\2\u0095\u0096\t\5\2\2\u0096\n\3\2\2\2\u0097\u0098\t\6\2\2\u0098"+
		"\f\3\2\2\2\u0099\u009a\t\7\2\2\u009a\16\3\2\2\2\u009b\u009c\t\b\2\2\u009c"+
		"\20\3\2\2\2\u009d\u009e\t\t\2\2\u009e\22\3\2\2\2\u009f\u00a0\t\n\2\2\u00a0"+
		"\24\3\2\2\2\u00a1\u00a2\t\13\2\2\u00a2\26\3\2\2\2\u00a3\u00a4\t\f\2\2"+
		"\u00a4\30\3\2\2\2\u00a5\u00a6\t\r\2\2\u00a6\32\3\2\2\2\u00a7\u00a8\t\16"+
		"\2\2\u00a8\34\3\2\2\2\u00a9\u00aa\t\17\2\2\u00aa\36\3\2\2\2\u00ab\u00ac"+
		"\t\20\2\2\u00ac \3\2\2\2\u00ad\u00ae\t\21\2\2\u00ae\"\3\2\2\2\u00af\u00b0"+
		"\t\22\2\2\u00b0$\3\2\2\2\u00b1\u00b2\t\23\2\2\u00b2&\3\2\2\2\u00b3\u00b4"+
		"\t\24\2\2\u00b4(\3\2\2\2\u00b5\u00b6\t\25\2\2\u00b6*\3\2\2\2\u00b7\u00b8"+
		"\t\26\2\2\u00b8,\3\2\2\2\u00b9\u00ba\t\27\2\2\u00ba.\3\2\2\2\u00bb\u00bc"+
		"\t\30\2\2\u00bc\60\3\2\2\2\u00bd\u00be\t\31\2\2\u00be\62\3\2\2\2\u00bf"+
		"\u00c0\t\32\2\2\u00c0\64\3\2\2\2\u00c1\u00c2\t\33\2\2\u00c2\66\3\2\2\2"+
		"\u00c3\u00c4\5\'\24\2\u00c4\u00c5\5\13\6\2\u00c5\u00c6\5\35\17\2\u00c6"+
		"\u00c7\5)\25\2\u00c7\u00c8\5\13\6\2\u00c8\u00c9\5\35\17\2\u00c9\u00ca"+
		"\5\7\4\2\u00ca\u00cb\5\23\n\2\u00cb\u00cc\5\3\2\2\u00cc8\3\2\2\2\u00cd"+
		"\u00ce\5\3\2\2\u00ce\u00cf\5\35\17\2\u00cf\u00d0\5\t\5\2\u00d0:\3\2\2"+
		"\2\u00d1\u00d2\5\37\20\2\u00d2\u00d3\5%\23\2\u00d3<\3\2\2\2\u00d4\u00d5"+
		"\5\23\n\2\u00d5\u00d6\5\r\7\2\u00d6>\3\2\2\2\u00d7\u00d8\5\13\6\2\u00d8"+
		"\u00d9\5\35\17\2\u00d9\u00da\5)\25\2\u00da\u00db\5\37\20\2\u00db\u00dc"+
		"\5\35\17\2\u00dc\u00dd\5\7\4\2\u00dd\u00de\5\13\6\2\u00de\u00df\5\'\24"+
		"\2\u00df@\3\2\2\2\u00e0\u00e1\5\7\4\2\u00e1\u00e2\5\37\20\2\u00e2\u00e3"+
		"\5\35\17\2\u00e3\u00e4\5\'\24\2\u00e4\u00e5\5\37\20\2\u00e5\u00e6\5\31"+
		"\r\2\u00e6\u00e7\5\3\2\2\u00e7B\3\2\2\2\u00e8\u00e9\5\'\24\2\u00e9\u00ea"+
		"\5)\25\2\u00ea\u00eb\5%\23\2\u00eb\u00ec\5\23\n\2\u00ec\u00ed\5\35\17"+
		"\2\u00ed\u00ee\5\17\b\2\u00eeD\3\2\2\2\u00ef\u00f0\5\23\n\2\u00f0\u00f1"+
		"\5\35\17\2\u00f1\u00f2\5)\25\2\u00f2\u00f3\5\13\6\2\u00f3\u00f4\5\17\b"+
		"\2\u00f4\u00f5\5\13\6\2\u00f5\u00f6\5%\23\2\u00f6F\3\2\2\2\u00f7\u00f8"+
		"\5%\23\2\u00f8\u00f9\5\13\6\2\u00f9\u00fa\5\3\2\2\u00fa\u00fb\5\31\r\2"+
		"\u00fbH\3\2\2\2\u00fc\u00fd\5\5\3\2\u00fd\u00fe\5\37\20\2\u00fe\u00ff"+
		"\5\37\20\2\u00ff\u0100\5\31\r\2\u0100\u0101\5\13\6\2\u0101\u0102\5\3\2"+
		"\2\u0102\u0103\5\35\17\2\u0103J\3\2\2\2\u0104\u0105\5)\25\2\u0105\u0106"+
		"\5%\23\2\u0106\u0107\5+\26\2\u0107\u0108\5\13\6\2\u0108L\3\2\2\2\u0109"+
		"\u010a\5\r\7\2\u010a\u010b\5\3\2\2\u010b\u010c\5\31\r\2\u010c\u010d\5"+
		"\'\24\2\u010d\u010e\5\13\6\2\u010eN\3\2\2\2\u010f\u0110\5!\21\2\u0110"+
		"\u0111\5+\26\2\u0111\u0112\5\5\3\2\u0112\u0113\5\31\r\2\u0113\u0114\5"+
		"\23\n\2\u0114\u0115\5\7\4\2\u0115\u0116\5\37\20\2\u0116P\3\2\2\2\u0117"+
		"\u0118\5!\21\2\u0118\u0119\5%\23\2\u0119\u011a\5\23\n\2\u011a\u011b\5"+
		"-\27\2\u011b\u011c\5\3\2\2\u011c\u011d\5\t\5\2\u011d\u011e\5\37\20\2\u011e"+
		"R\3\2\2\2\u011f\u0120\5\33\16\2\u0120\u0121\5\3\2\2\u0121\u0122\5\23\n"+
		"\2\u0122\u0123\5\35\17\2\u0123T\3\2\2\2\u0124\u0125\5!\21\2\u0125\u0126"+
		"\5%\23\2\u0126\u0127\5\23\n\2\u0127\u0128\5\35\17\2\u0128\u0129\5\7\4"+
		"\2\u0129\u012a\5\23\n\2\u012a\u012b\5!\21\2\u012b\u012c\5\3\2\2\u012c"+
		"\u012d\5\31\r\2\u012dV\3\2\2\2\u012e\u012f\5\33\16\2\u012f\u0130\5\13"+
		"\6\2\u0130\u0131\5)\25\2\u0131\u0132\5\37\20\2\u0132\u0133\5\t\5\2\u0133"+
		"\u0134\5\37\20\2\u0134X\3\2\2\2\u0135\u0136\5\7\4\2\u0136\u0137\5\31\r"+
		"\2\u0137\u0138\5\3\2\2\u0138\u0139\5\'\24\2\u0139\u013a\5\13\6\2\u013a"+
		"Z\3\2\2\2\u013b\u013c\5\t\5\2\u013c\u013d\5\13\6\2\u013d\u013e\5\7\4\2"+
		"\u013e\u013f\5\31\r\2\u013f\u0140\5\3\2\2\u0140\u0141\5%\23\2\u0141\u0142"+
		"\5\3\2\2\u0142\u0143\5%\23\2\u0143\\\3\2\2\2\u0144\u0146\t\34\2\2\u0145"+
		"\u0144\3\2\2\2\u0146\u0147\3\2\2\2\u0147\u0145\3\2\2\2\u0147\u0148\3\2"+
		"\2\2\u0148^\3\2\2\2\u0149\u014b\t\34\2\2\u014a\u0149\3\2\2\2\u014b\u014c"+
		"\3\2\2\2\u014c\u014a\3\2\2\2\u014c\u014d\3\2\2\2\u014d\u014e\3\2\2\2\u014e"+
		"\u0150\7\60\2\2\u014f\u0151\t\34\2\2\u0150\u014f\3\2\2\2\u0151\u0152\3"+
		"\2\2\2\u0152\u0150\3\2\2\2\u0152\u0153\3\2\2\2\u0153`\3\2\2\2\u0154\u0156"+
		"\t\35\2\2\u0155\u0154\3\2\2\2\u0156\u015a\3\2\2\2\u0157\u0159\t\36\2\2"+
		"\u0158\u0157\3\2\2\2\u0159\u015c\3\2\2\2\u015a\u0158\3\2\2\2\u015a\u015b"+
		"\3\2\2\2\u015bb\3\2\2\2\u015c\u015a\3\2\2\2\u015d\u015e\7<\2\2\u015ed"+
		"\3\2\2\2\u015f\u0160\7\60\2\2\u0160f\3\2\2\2\u0161\u0162\7=\2\2\u0162"+
		"h\3\2\2\2\u0163\u0164\7@\2\2\u0164\u0165\7?\2\2\u0165j\3\2\2\2\u0166\u0167"+
		"\7@\2\2\u0167l\3\2\2\2\u0168\u0169\7>\2\2\u0169n\3\2\2\2\u016a\u016b\7"+
		"?\2\2\u016bp\3\2\2\2\u016c\u016d\7\'\2\2\u016dr\3\2\2\2\u016e\u016f\7"+
		",\2\2\u016ft\3\2\2\2\u0170\u0171\7\61\2\2\u0171v\3\2\2\2\u0172\u0173\7"+
		"/\2\2\u0173x\3\2\2\2\u0174\u0175\7-\2\2\u0175z\3\2\2\2\u0176\u0177\7*"+
		"\2\2\u0177|\3\2\2\2\u0178\u0179\7+\2\2\u0179~\3\2\2\2\u017a\u017b\7}\2"+
		"\2\u017b\u0080\3\2\2\2\u017c\u017d\7\177\2\2\u017d\u0082\3\2\2\2\u017e"+
		"\u017f\7]\2\2\u017f\u0084\3\2\2\2\u0180\u0181\7_\2\2\u0181\u0086\3\2\2"+
		"\2\u0182\u0186\7$\2\2\u0183\u0185\13\2\2\2\u0184\u0183\3\2\2\2\u0185\u0188"+
		"\3\2\2\2\u0186\u0187\3\2\2\2\u0186\u0184\3\2\2\2\u0187\u0189\3\2\2\2\u0188"+
		"\u0186\3\2\2\2\u0189\u018a\7$\2\2\u018a\u0088\3\2\2\2\u018b\u018d\t\37"+
		"\2\2\u018c\u018b\3\2\2\2\u018d\u018e\3\2\2\2\u018e\u018c\3\2\2\2\u018e"+
		"\u018f\3\2\2\2\u018f\u0190\3\2\2\2\u0190\u0191\bE\2\2\u0191\u008a\3\2"+
		"\2\2\u0192\u0193\7\61\2\2\u0193\u0194\7,\2\2\u0194\u0198\3\2\2\2\u0195"+
		"\u0197\13\2\2\2\u0196\u0195\3\2\2\2\u0197\u019a\3\2\2\2\u0198\u0199\3"+
		"\2\2\2\u0198\u0196\3\2\2\2\u0199\u019b\3\2\2\2\u019a\u0198\3\2\2\2\u019b"+
		"\u019c\7,\2\2\u019c\u019d\7\61\2\2\u019d\u019e\3\2\2\2\u019e\u019f\bF"+
		"\2\2\u019f\u008c\3\2\2\2\u01a0\u01a1\7\61\2\2\u01a1\u01a2\7\61\2\2\u01a2"+
		"\u01a6\3\2\2\2\u01a3\u01a5\n \2\2\u01a4\u01a3\3\2\2\2\u01a5\u01a8\3\2"+
		"\2\2\u01a6\u01a4\3\2\2\2\u01a6\u01a7\3\2\2\2\u01a7\u01a9\3\2\2\2\u01a8"+
		"\u01a6\3\2\2\2\u01a9\u01aa\bG\2\2\u01aa\u008e\3\2\2\2\r\2\u0147\u014c"+
		"\u0152\u0155\u0158\u015a\u0186\u018e\u0198\u01a6\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}