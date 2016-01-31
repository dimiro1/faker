package locales

/*
 * RULES
 * 1. The elements must be alphabetically ordered
 * 2. Maximun of 120 chars per lines
 */

func en_locale() Locale {
	return Locale{
		Code: "en",

		// This list was extracted from this github project https://github.com/stympy/faker/
		// See https://github.com/stympy/faker/blob/master/lib/locales/en.yml
		AppNames: []string{
			"Redhold", "Treeflex", "Trippledex", "Kanlam", "Bigtax", "Daltfresh", "Toughjoyfax", "Mat Lam Tam",
			"Otcom", "Tres-Zap", "Y-Solowarm", "Tresom", "Voltsillam", "Biodex", "Greenlam", "Viva", "Matsoft", "Temp",
			"Zoolab", "Subin", "Rank", "Job", "Stringtough", "Tin", "It", "Home Ing", "Zamit", "Sonsing", "Konklab",
			"Alpha", "Latlux", "Voyatouch", "Alphazap", "Holdlamis", "Zaam-Dox", "Sub-Ex", "Quo Lux", "Bamity",
			"Ventosanzap", "Lotstring", "Hatity", "Tempsoft", "Overhold", "Fixflex", "Konklux", "Zontrax", "Tampflex",
			"Span", "Namfix", "Transcof", "Stim", "Fix San", "Sonair", "Stronghold", "Fintone", "Y-find", "Opela",
			"Lotlux", "Ronstring", "Zathin", "Duobam", "Keylex", "Andalax", "Solarbreeze", "Cookley", "Vagram",
			"Aerified", "Pannier", "Asoka", "Regrant", "Wrapsafe", "Prodder", "Bytecard", "Bitchip", "Veribet",
			"Gembucket", "Cardguard", "Bitwolf", "Cardify", "Domainer", "Flowdesk", "Flexidy",
		},

		// Specific to USA
		AreaCodes: []string{
			"201", "202", "203", "205", "206", "207", "208", "209", "210", "212", "213", "214", "215", "216", "217",
			"218", "219", "224", "225", "227", "228", "229", "231", "234", "239", "240", "248", "251", "252", "253",
			"254", "256", "260", "262", "267", "269", "270", "276", "281", "283", "301", "302", "303", "304", "305",
			"307", "308", "309", "310", "312", "313", "314", "315", "316", "317", "318", "319", "320", "321", "323",
			"330", "331", "334", "336", "337", "339", "347", "351", "352", "360", "361", "386", "401", "402", "404",
			"405", "406", "407", "408", "409", "410", "412", "413", "414", "415", "417", "419", "423", "424", "425",
			"434", "435", "440", "443", "445", "464", "469", "470", "475", "478", "479", "480", "484", "501", "502",
			"503", "504", "505", "507", "508", "509", "510", "512", "513", "515", "516", "517", "518", "520", "530",
			"540", "541", "551", "557", "559", "561", "562", "563", "564", "567", "570", "571", "573", "574", "580",
			"585", "586", "601", "602", "603", "605", "606", "607", "608", "609", "610", "612", "614", "615", "616",
			"617", "618", "619", "620", "623", "626", "630", "631", "636", "641", "646", "650", "651", "660", "661",
			"662", "667", "678", "682", "701", "702", "703", "704", "706", "707", "708", "712", "713", "714", "715",
			"716", "717", "718", "719", "720", "724", "727", "731", "732", "734", "737", "740", "754", "757", "760",
			"763", "765", "770", "772", "773", "774", "775", "781", "785", "786", "801", "802", "803", "804", "805",
			"806", "808", "810", "812", "813", "814", "815", "816", "817", "818", "828", "830", "831", "832", "835",
			"843", "845", "847", "848", "850", "856", "857", "858", "859", "860", "862", "863", "864", "865", "870",
			"872", "878", "901", "903", "904", "906", "907", "908", "909", "910", "912", "913", "914", "915", "916",
			"917", "918", "919", "920", "925", "928", "931", "936", "937", "940", "941", "947", "949", "952", "954",
			"956", "959", "970", "971", "972", "973", "975", "978", "979", "980", "984", "985", "989",
		},
		CityPrefix: []string{"North", "East", "West", "South", "New", "Lake", "Port"},
		CitySuffix: []string{
			"town", "ton", "land", "ville", "berg", "burgh", "borough", "bury", "view", "port", "mouth", "stad", "furt",
			"chester", "mouth", "fort", "haven", "side", "shire"},
		CityNames: []string{
			"Albuquerque", "Arlington", "Atlanta", "Austin", "Baltimore", "Boston", "Charlotte", "Chicago", "Cleveland",
			"Colorado Springs", "Columbus", "Dallas", "Denver", "Detroit", "El Paso", "Fort Worth", "Fresno", "Honolulu",
			"Houston", "Indianapolis", "Jacksonville", "Kansas City", "Las Vegas", "Long Beach", "Los Angeles",
			"Louisville/Jefferson County", "Memphis", "Mesa", "Miami", "Milwaukee", "Minneapolis", "Nashville-Davidson",
			"New Orleans", "New York", "Oakland", "Oklahoma City", "Omaha", "Philadelphia", "Phoenix", "Portland",
			"Raleigh", "Sacramento", "St. Louis", "San Antonio", "San Diego", "San Francisco", "San Jose", "Seattle",
			"Tucson", "Tulsa", "Virginia Beach", "Washington", "Wichita",
		},
		CountryNames: []string{
			"Afghanistan", "Akrotiri", "Albania", "Algeria", "American Samoa", "Andorra", "Angola", "Anguilla",
			"Antarctica", "Antigua and Barbuda", "Argentina", "Armenia", "Aruba", "Ashmore and Cartier Islands",
			"Australia", "Austria", "Azerbaijan", "Bahamas, The", "Bahrain", "Bangladesh", "Barbados",
			"Bassas da India", "Belarus", "Belgium", "Belize", "Benin", "Bermuda", "Bhutan", "Bolivia",
			"Bosnia and Herzegovina", "Botswana", "Bouvet Island", "Brazil", "British Indian Ocean Territory",
			"British Virgin Islands", "Brunei", "Bulgaria", "Burkina Faso", "Burma", "Burundi", "Cambodia", "Cameroon",
			"Canada", "Cape Verde", "Cayman Islands", "Central African Republic", "Chad", "Chile", "China",
			"Christmas Island", "Clipperton Island", "Cocos (Keeling) Islands", "Colombia", "Comoros",
			"Congo, Democratic Republic of the", "Congo, Republic of the", "Cook Islands", "Coral Sea Islands",
			"Costa Rica", "Cote d'Ivoire", "Croatia", "Cuba", "Cyprus", "Czech Republic", "Denmark", "Dhekelia",
			"Djibouti", "Dominica", "Dominican Republic", "Ecuador", "Egypt", "El Salvador", "Equatorial Guinea",
			"Eritrea", "Estonia", "Ethiopia", "Europa Island", "Falkland Islands (Islas Malvinas)", "Faroe Islands",
			"Fiji", "Finland", "France", "French Guiana", "French Polynesia", "French Southern and Antarctic Lands",
			"Gabon", "Gambia, The", "Gaza Strip", "Georgia", "Germany", "Ghana", "Gibraltar", "Glorioso Islands",
			"Greece", "Greenland", "Grenada", "Guadeloupe", "Guam", "Guatemala", "Guernsey", "Guinea", "Guinea-Bissau",
			"Guyana", "Haiti", "Heard Island and McDonald Islands", "Holy See (Vatican City)", "Honduras", "Hong Kong",
			"Hungary", "Iceland", "India", "Indonesia", "Iran", "Iraq", "Ireland", "Isle of Man", "Israel", "Italy",
			"Jamaica", "Jan Mayen", "Japan", "Jersey", "Jordan", "Juan de Nova Island", "Kazakhstan", "Kenya",
			"Kiribati", "Korea, North", "Korea, South", "Kuwait", "Kyrgyzstan", "Laos", "Latvia", "Lebanon",
			"Lesotho", "Liberia", "Libya", "Liechtenstein", "Lithuania", "Luxembourg", "Macau", "Macedonia",
			"Madagascar", "Malawi", "Malaysia", "Maldives", "Mali", "Malta", "Marshall Islands", "Martinique",
			"Mauritania", "Mauritius", "Mayotte", "Mexico", "Micronesia, Federated States of", "Moldova", "Monaco",
			"Mongolia", "Montserrat", "Morocco", "Mozambique", "Namibia", "Nauru", "Navassa Island", "Nepal",
			"Netherlands", "Netherlands Antilles", "New Caledonia", "New Zealand", "Nicaragua", "Niger", "Nigeria",
			"Niue", "Norfolk Island", "Northern Mariana Islands", "Norway", "Oman", "Pakistan", "Palau", "Panama",
			"Papua New Guinea", "Paracel Islands", "Paraguay", "Peru", "Philippines", "Pitcairn Islands", "Poland",
			"Portugal", "Puerto Rico", "Qatar", "Reunion", "Romania", "Russia", "Rwanda", "Saint Helena",
			"Saint Kitts and Nevis", "Saint Lucia", "Saint Pierre and Miquelon", "Saint Vincent and the Grenadines",
			"Samoa", "San Marino", "Sao Tome and Principe", "Saudi Arabia", "Senegal", "Serbia and Montenegro",
			"Seychelles", "Sierra Leone", "Singapore", "Slovakia", "Slovenia", "Solomon Islands", "Somalia",
			"South Africa", "South Georgia and the South Sandwich Islands", "Spain", "Spratly Islands", "Sri Lanka",
			"Sudan", "Suriname", "Svalbard", "Swaziland", "Sweden", "Switzerland", "Syria", "Taiwan", "Tajikistan",
			"Tanzania", "Thailand", "Timor-Leste", "Togo", "Tokelau", "Tonga", "Trinidad and Tobago",
			"Tromelin Island", "Tunisia", "Turkey", "Turkmenistan", "Turks and Caicos Islands", "Tuvalu", "Uganda",
			"Ukraine", "United Arab Emirates", "United Kingdom", "United States", "Uruguay", "Uzbekistan", "Vanuatu",
			"Venezuela", "Vietnam", "Virgin Islands", "Wake Island", "Wallis and Futuna", "West Bank",
			"Western Sahara", "Yemen", "Zambia", "Zimbabwe"},

		// Specific to USA
		ExchangeCodes: []string{
			"201", "202", "203", "205", "206", "207", "208", "209", "210", "212", "213", "214", "215", "216", "217",
			"218", "219", "224", "225", "227", "228", "229", "231", "234", "239", "240", "248", "251", "252", "253",
			"254", "256", "260", "262", "267", "269", "270", "276", "281", "283", "301", "302", "303", "304", "305",
			"307", "308", "309", "310", "312", "313", "314", "315", "316", "317", "318", "319", "320", "321", "323",
			"330", "331", "334", "336", "337", "339", "347", "351", "352", "360", "361", "386", "401", "402", "404",
			"405", "406", "407", "408", "409", "410", "412", "413", "414", "415", "417", "419", "423", "424", "425",
			"434", "435", "440", "443", "445", "464", "469", "470", "475", "478", "479", "480", "484", "501", "502",
			"503", "504", "505", "507", "508", "509", "510", "512", "513", "515", "516", "517", "518", "520", "530",
			"540", "541", "551", "557", "559", "561", "562", "563", "564", "567", "570", "571", "573", "574", "580",
			"585", "586", "601", "602", "603", "605", "606", "607", "608", "609", "610", "612", "614", "615", "616",
			"617", "618", "619", "620", "623", "626", "630", "631", "636", "641", "646", "650", "651", "660", "661",
			"662", "667", "678", "682", "701", "702", "703", "704", "706", "707", "708", "712", "713", "714", "715",
			"716", "717", "718", "719", "720", "724", "727", "731", "732", "734", "737", "740", "754", "757", "760",
			"763", "765", "770", "772", "773", "774", "775", "781", "785", "786", "801", "802", "803", "804", "805",
			"806", "808", "810", "812", "813", "814", "815", "816", "817", "818", "828", "830", "831", "832", "835",
			"843", "845", "847", "848", "850", "856", "857", "858", "859", "860", "862", "863", "864", "865", "870",
			"872", "878", "901", "903", "904", "906", "907", "908", "909", "910", "912", "913", "914", "915", "916",
			"917", "918", "919", "920", "925", "928", "931", "936", "937", "940", "941", "947", "949", "952", "954",
			"956", "959", "970", "971", "972", "973", "975", "978", "979", "980", "984", "985", "989",
		},
		StateNames: []string{"Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado", "Connecticut",
			"Delaware", "Florida", "Georgia", "Hawaii", "Idaho", "Illinois", "Indiana", "Iowa", "Kansas", "Kentucky",
			"Louisiana", "Maine", "Maryland", "Massachusetts", "Michigan", "Minnesota", "Mississippi", "Missouri",
			"Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey", "New Mexico", "New York", "North Carolina",
			"North Dakota", "Ohio", "Oklahoma", "Oregon", "Pennsylvania", "Rhode Island", "South Carolina",
			"South Dakota", "Tennessee", "Texas", "Utah", "Vermont", "Virginia", "Washington", "West Virginia",
			"Wisconsin", "Wyoming"},
		StateAbbr: []string{"AL", "AK", "AZ", "AR", "CA", "CO", "CT", "DE", "FL", "GA", "HI", "ID", "IL", "IN", "IA",
			"KS", "KY", "LA", "ME", "MD", "MA", "MI", "MN", "MS", "MO", "MT", "NE", "NV", "NH", "NJ", "NM", "NY", "NC",
			"ND", "OH", "OK", "OR", "PA", "RI", "SC", "SD", "TN", "TX", "UT", "VT", "VA", "WA", "WV", "WI", "WY"},

		// See http://www.forbes.com/top-colleges/list/2/#tab:rank
		Universities: []string{
			"Amherst College", "Boston College", "Bowdoin College", "Brown University",
			"California Institute of Technology", "Carleton College", "Claremont McKenna College",
			"Colgate University", "College of William & Mary", "Columbia University", "Cornell University",
			"Dartmouth College", "Davidson College", "Duke University", "Georgetown University", "Harvard University",
			"Haverford College", "Massachusetts Institute of Technology", "Middlebury College",
			"Northwestern University", "Pomona College", "Princeton University", "Rice University",
			"Stanford University", "Swarthmore College", "Tufts University", "United States Air Force Academy",
			"United States Military Academy", "United States Naval Academy", "University of California, Berkeley",
			"University of Chicago", "University of Notre Dame", "University of Pennsylvania", "University of Virginia",
			"Vassar College", "Washington and Lee University", "Wellesley College", "Wesleyan University",
			"Williams College", "Yale University",
		},
	}
}
