package faker

/*
 * RULES
 * 1. The elements must be alphabetically ordered
 * 2. Maximun of 120 chars per lines
 */

var (
	// countryCodes is an intertational ISO code, this list should not be in locales.
	countryCodes = []string{
		"AF", "AL", "DZ", "AS", "AD", "AO", "AI", "AQ", "AG", "AR", "AM", "AW", "AU", "AT", "AZ", "BS", "BH",
		"BD", "BB", "BY", "BE", "BZ", "BJ", "BM", "BT", "BO", "BQ", "BA", "BW", "BV", "BR", "IO", "BN", "BG",
		"BF", "BI", "KH", "CM", "CA", "CV", "KY", "CF", "TD", "CL", "CN", "CX", "CC", "CO", "KM", "CG", "CD",
		"CK", "CR", "HR", "CU", "CW", "CY", "CZ", "CI", "DK", "DJ", "DM", "DO", "EC", "EG", "SV", "GQ", "ER",
		"EE", "ET", "FK", "FO", "FJ", "FI", "FR", "GF", "PF", "TF", "GA", "GM", "GE", "DE", "GH", "GI", "GR",
		"GL", "GD", "GP", "GU", "GT", "GG", "GN", "GW", "GY", "HT", "HM", "VA", "HN", "HK", "HU", "IS", "IN",
		"ID", "IR", "IQ", "IE", "IM", "IL", "IT", "JM", "JP", "JE", "JO", "KZ", "KE", "KI", "KP", "KR", "KW",
		"KG", "LA", "LV", "LB", "LS", "LR", "LY", "LI", "LT", "LU", "MO", "MK", "MG", "MW", "MY", "MV", "ML",
		"MT", "MH", "MQ", "MR", "MU", "YT", "MX", "FM", "MD", "MC", "MN", "ME", "MS", "MA", "MZ", "MM", "NA",
		"NR", "NP", "NL", "NC", "NZ", "NI", "NE", "NG", "NU", "NF", "MP", "NO", "OM", "PK", "PW", "PS", "PA",
		"PG", "PY", "PE", "PH", "PN", "PL", "PT", "PR", "QA", "RO", "RU", "RW", "RE", "BL", "SH", "KN", "LC",
		"MF", "PM", "VC", "WS", "SM", "ST", "SA", "SN", "RS", "SC", "SL", "SG", "SX", "SK", "SI", "SB", "SO",
		"ZA", "GS", "SS", "ES", "LK", "SD", "SR", "SJ", "SZ", "SE", "CH", "SY", "TW", "TJ", "TZ", "TH", "TL",
		"TG", "TK", "TO", "TT", "TN", "TR", "TM", "TC", "TV", "UG", "UA", "AE", "GB", "US", "UM", "UY", "UZ",
		"VU", "VE", "VN", "VG", "VI", "WF", "EH", "YE", "ZM", "ZW", "AX"}

	// Common used operating systems
	operatingSystems = []string{
		"AmigaOS", "Android", "BeOS", "Chrome OS", "DOS", "FreeBSD", "GNU/Linux", "HP-UX", "IOS", "MacOS", "Minix",
		"NetBSD", "OpenBSD", "OS/2", "Plan 9", "Solaris", "Windows",
	}

	// Well known programming languages
	programmingLanguages = []string{
		"Ada", "Algol", "AWK", "BASIC", "C", "C#", "C++", "Clojure", "COBOL", "D", "Dart", "Delphi", "Elixir",
		"Emacs Lisp", "Erlang", "F#", "Fortran", "Go", "Groovy", "Hack", "Haskell", "Haxe", "Java", "JavaScript",
		"Julia", "Lua", "Modula-2", "Modula-3", "Nim", "Objective-C", "OCaml", "Pascal", "Perl", "PHP", "PL/SQL",
		"Prolog", "Python", "R", "REBOL", "Ruby", "Rust", "Scala", "Scheme", "Self", "Simula", "Smalltalk", "SQL",
		"Swift", "TECO", "TEX", "Visual Basic",
	}

	// Well known Text editors
	textEditors = []string{
		"Acme", "Atom", "BBEdit", "Brackets", "Emacs", "Geany", "Gedit", "jEdit", "Kate", "Nano", "Notepad++", "SciTE",
		"Sublime Text", "TextMate", "VIM",
	}

	// List of web frameworks
	// See http://hotframeworks.com/
	// See https://en.wikipedia.org/wiki/Comparison_of_web_frameworks
	webFrameworks = []string{
		"ASP.NET", "beego", "Bottle", "CakePHP", "CherryPy", "Dancer", "Django", "Express", "Flask", "Flex",
		"FuelPHP", "Grails", "Grock", "GWT", "JSF", "Kohana", "Laravel", "Lift", "Merb", "Noir", "Padrino", "Play!",
		"Pyramid", "Ramaze", "Revel", "Ruby on Rails", "Scalatra", "seaside", "Sinatra", "Spring", "Sympfony",
		"Tapestry", "Tornado", "TurboGears", "Vaadin", "VRaptor", "web.py", "web2py", "Wicket", "Zend", "Zope",
	}
)
