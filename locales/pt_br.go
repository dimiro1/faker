package locales

/*
 * RULES
 * 1. The elements must be alphabetically ordered
 * 2. Maximun of 120 chars per lines
 */

func pt_br_locale() Locale {
	return Locale{
		Code: "pt-br",

		AddressesFormats:       []string{},
		AppNames:               []string{},
		AreaCodes:              []string{},
		BookTitles:             []string{},
		BookGenres:             []string{},
		BookPublishersFormats:  []string{},
		BookPublishersSuffixes: []string{},
		BuildingNumbers: []string{
			"@", "@#", "@##",
		},
		CityPrefix: []string{"Nova", "Velha", "Grande", "Vila", "Município de"},
		CitySuffix: []string{"do Descoberto", "de Nossa Senhora", "do Norte", "do Sul"},

		CityNamesFormats: []string{
			"{{ .CityPrefix }} {{ .FirstName }} {{ .CitySuffix }}",
			"{{ .CityPrefix }} {{ .FirstName }}",
			"{{ .LastName }} {{ .CitySuffix }}",
		},

		CompanyNamesFormats: []string{
			"{{ .LastName }} {{ .CompanySuffix }}",
			"{{ .LastName }}-{{ .LastName }}",
			"{{ .LastName }} & {{ .LastName }}",
			"{{ .LastName }} {{ .LastName }} e {{ .LastName }}",
		},

		CompanySuffixes: []string{
			"LTDA", "S/A", "e filhos", "e netos",
		},

		CountryNames: []string{
			"Afeganistão", "África do Sul", "Akrotiri", "Albânia", "Alemanha", "Andorra", "Angola", "Anguila",
			"Antárctida", "Antígua e Barbuda", "Antilhas Neerlandesas", "Arábia Saudita", "Arctic Ocean", "Argélia",
			"Argentina", "Arménia", "Aruba", "Ashmore and Cartier Islands", "Atlantic Ocean", "Austrália", "Áustria",
			"Azerbaijão", "Baamas", "Bangladeche", "Barbados", "Barém", "Bélgica", "Belize", "Benim", "Bermudas",
			"Bielorrússia", "Birmânia", "Bolívia", "Bósnia e Herzegovina", "Botsuana", "Brasil", "Brunei", "Bulgária",
			"Burquina Faso", "Burúndi", "Butão", "Cabo Verde", "Camarões", "Camboja", "Canadá", "Catar", "Cazaquistão",
			"Chade", "Chile", "China", "Chipre", "Clipperton Island", "Colômbia", "Comores", "Congo-Brazzaville",
			"Congo-Kinshasa", "Coral Sea Islands", "Coreia do Norte", "Coreia do Sul", "Costa do Marfim", "Costa Rica",
			"Croácia", "Cuba", "Dhekelia", "Dinamarca", "Domínica", "Egipto", "Emiratos Árabes Unidos", "Equador",
			"Eritreia", "Eslováquia", "Eslovénia", "Espanha", "Estados Unidos", "Estónia", "Etiópia", "Faroé", "Fiji",
			"Filipinas", "Finlândia", "França", "Gabão", "Gâmbia", "Gana", "Gaza Strip", "Geórgia",
			"Geórgia do Sul e Sandwich do Sul", "Gibraltar", "Granada", "Grécia", "Gronelândia", "Guame", "Guatemala",
			"Guernsey", "Guiana", "Guiné", "Guiné Equatorial", "Guiné-Bissau", "Haiti", "Honduras", "Hong Kong",
			"Hungria", "Iémen", "Ilha Bouvet", "Ilha do Natal", "Ilha Norfolk", "Ilhas Caimão", "Ilhas Cook",
			"Ilhas dos Cocos", "Ilhas Falkland", "Ilhas Heard e McDonald", "Ilhas Marshall", "Ilhas Salomão",
			"Ilhas Turcas e Caicos", "Ilhas Virgens Americanas", "Ilhas Virgens Britânicas", "Índia", "Indian Ocean",
			"Indonésia", "Irão", "Iraque", "Irlanda", "Islândia", "Israel", "Itália", "Jamaica", "Jan Mayen", "Japão",
			"Jersey", "Jibuti", "Jordânia", "Kuwait", "Laos", "Lesoto", "Letónia", "Líbano", "Libéria", "Líbia",
			"Listenstaine", "Lituânia", "Luxemburgo", "Macau", "Macedónia", "Madagáscar", "Malásia", "Malávi",
			"Maldivas", "Mali", "Malta", "Man, Isle of", "Marianas do Norte", "Marrocos", "Maurícia", "Mauritânia",
			"Mayotte", "México", "Micronésia", "Moçambique", "Moldávia", "Mónaco", "Mongólia", "Monserrate",
			"Montenegro", "Mundo", "Namíbia", "Nauru", "Navassa Island", "Nepal", "Nicarágua", "Níger", "Nigéria",
			"Niue", "Noruega", "Nova Caledónia", "Nova Zelândia", "Omã", "Pacific Ocean", "Países Baixos", "Palau",
			"Panamá", "Papua-Nova Guiné", "Paquistão", "Paracel Islands", "Paraguai", "Peru", "Pitcairn",
			"Polinésia Francesa", "Polónia", "Porto Rico", "Portugal", "Quénia", "Quirguizistão", "Quiribáti",
			"Reino Unido", "República Centro-Africana", "República Checa", "República Dominicana", "Roménia", "Ruanda",
			"Rússia", "Salvador", "Samoa", "Samoa Americana", "Santa Helena", "Santa Lúcia", "São Cristóvão e Neves",
			"São Marinho", "São Pedro e Miquelon", "São Tomé e Príncipe", "São Vicente e Granadinas", "Sara Ocidental",
			"Seicheles", "Senegal", "Serra Leoa", "Sérvia", "Singapura", "Síria", "Somália", "Southern Ocean",
			"Spratly Islands", "Sri Lanca", "Suazilândia", "Sudão", "Suécia", "Suíça", "Suriname",
			"Svalbard e Jan Mayen", "Tailândia", "Taiwan", "Tajiquistão", "Tanzânia",
			"Território Britânico do Oceano Índico", "Territórios Austrais Franceses", "Timor Leste", "Togo",
			"Tokelau", "Tonga", "Trindade e Tobago", "Tunísia", "Turquemenistão", "Turquia", "Tuvalu", "Ucrânia",
			"Uganda", "União Europeia", "Uruguai", "Usbequistão", "Vanuatu", "Vaticano", "Venezuela", "Vietname",
			"Wake Island", "Wallis e Futuna", "West Bank", "Zâmbia", "Zimbabué"},
		ExchangeCodes: []string{},
		FirstNames:    []string{},
		Gender: []string{
			"Masculino", "Feminino",
		},
		LastName: []string{},
		NamesFormats: []string{
			"{{ .FirstName }}",
		},
		PhoneNumbersFormats: []string{
			"+55 (011) #### ####", "+55 (021) #### ####", "+55 (031) #### ####", "+55 (041) #### ####",
			"+55 (051) #### ####", "+55 (061) #### ####", "+55 (071) #### ####", "+55 (081) #### ####",
			"+55 11 #### ####", "+55 21 #### ####", "+55 31 #### ####", "+55 41 #### ####", "+55 51 ### ####",
			"+55 61 #### ####", "+55 71 #### ####", "+55 81 #### ####", "+55 (011) ####-####", "+55 (021) ####-####",
			"+55 (031) ####-####", "+55 (041) ####-####", "+55 (051) ####-####", "+55 (061) ####-####",
			"+55 (071) ####-####", "+55 (081) ####-####", "+55 11 ####-####", "+55 21 ####-####", "+55 31 ####-####",
			"+55 41 ####-####", "+55 51 ### ####", "+55 61 ####-####", "+55 71 ####-####", "+55 81 ####-####",
			"(011) #### ####", "(021) #### ####", "(031) #### ####", "(041) #### ####", "(051) #### ####",
			"(061) #### ####", "(071) #### ####", "(081) #### ####", "11 #### ####", "21 #### ####", "31 #### ####",
			"41 #### ####", "51 ### ####", "61 #### ####", "71 #### ####", "81 #### ####", "(011) ####-####",
			"(021) ####-####", "(031) ####-####", "(041) ####-####", "(051) ####-####", "(061) ####-####",
			"(071) ####-####", "(081) ####-####", "11 ####-####", "21 ####-####", "31 ####-####", "41 ####-####",
			"51 ### ####", "61 ####-####", "71 ####-####", "81 ####-####", "#### ####", "####-####",
		},
		SecondaryAddressesFormats: []string{},
		StateNames:                []string{},
		StateAbbr:                 []string{},

		StreetAddressesFormats: []string{},
		StreetNamesFormats:     []string{},
		StreetSuffixes:         []string{},

		StoreDepartments: []string{},

		// See https://pt.wikipedia.org/wiki/Lista_de_universidades_do_Brasil
		Universities: []string{
			"Centro Universitário da Grande Dourados (UNIGRAN)", "Centro Universitário de Belo Horizonte (UniBH)",
			"Centro Universitário de Campo Grande (UNAES)", "Centro Universitário de Lavras (UNILAVRAS)",
			"Centro Universitário de Lins (UNILINS)", "Centro Universitário Dinâmica das Cataratas (UDC)",
			"Centro Universitário Estadual da Zona Oeste (UEZO)", "Centro Universitário Eurípides de Marília (UNIVEM)",
			"Centro Universitário Metodista Izabela Hendrix", "Faculdades Equipe (FAE)",
			"Faculdades Integradas de Três Lagoas (AEMS)",
			"Pontifícia Universidade Católica de Campinas (PUC-Campinas)",
			"Pontifícia Universidade Católica de Goiás (PUC-GO)",
			"Pontifícia Universidade Católica de Minas Gerais (PUC-Minas)",
			"Pontifícia Universidade Católica de São Paulo (PUC-SP)",
			"Pontifícia Universidade Católica do Paraná (PUC-PR)",
			"Pontifícia Universidade Católica do Rio de Janeiro (PUC-Rio)",
			"Pontifícia Universidade Católica do Rio Grande do Sul (PUC-RS)", "Universidade Anhembi Morumbi (UAM)",
			"Universidade Bandeirante de São Paulo (UNIBAN)", "Universidade Barriga Verde (UNIBAVE)",
			"Universidade Braz Cubas (UBC)", "Universidade Camilo Castelo Branco (UNICASTELO)",
			"Universidade Castelo Branco (UCB)", "Universidade Católica de Pelotas (UCPEL)",
			"Universidade Católica de Pernambuco (UNICAP)", "Universidade Católica de Petrópolis (UCP)",
			"Universidade Católica de Santos (UNISANTOS)", "Universidade Católica do Salvador (UCSAL)",
			"Universidade Católica Dom Bosco (UCDB)", "Universidade Central Paulista (UNICEP)",
			"Universidade Cidade de São Paulo (UNICID)", "Universidade Comunitária Regional de Chapecó (UNOCHAPECÓ)",
			"Universidade Cruzeiro do Sul (UNICSUL)", "Universidade Cândido Mendes (UCAM)",
			"Universidade da Amazônia (UNAMA)", "Universidade da Região da Campanha (URCAMP)",
			"Universidade da Região de Joinville (UNIVILLE)", "Universidade de Brasília (UnB)",
			"Universidade de Caxias do Sul (UCS)", "Universidade de Cruz Alta (UNICRUZ)",
			"Universidade de Cuiabá (UNIC)", "Universidade de Fortaleza (UNIFOR)", "Universidade de Franca (UNIFRAN)",
			"Universidade de Itaúna (UI)", "Universidade de Marília (UNIMAR)", "Universidade de Mogi das Cruzes (UMC)",
			"Universidade de Passo Fundo (UPF)", "Universidade de Pernambuco (UPE)",
			"Universidade de Ribeirão Preto (UNAERP)", "Universidade de Rio Verde (FESURV)",
			"Universidade de Santa Cruz do Sul (UNISC)", "Universidade de Santo Amaro (UNISA)",
			"Universidade de Sorocaba (UNISO)", "Universidade de São Paulo (USP)", "Universidade de Taubaté (UNITAU)",
			"Universidade de Uberaba (UNIUBE)", "Universidade do Contestado (UNC)",
			"Universidade do Estado da Bahia (UNEB)", "Universidade do Estado de Mato Grosso (UNEMAT)",
			"Universidade do Estado de Minas Gerais (UEMG)", "Universidade do Estado de Santa Catarina (UDESC)",
			"Universidade do Estado do Amapá (UEAP)", "Universidade do Estado do Amazonas (UEA)",
			"Universidade do Estado do Pará (UEPA)", "Universidade do Estado do Rio de Janeiro (UERJ)",
			"Universidade do Estado do Rio Grande do Norte (UERN)", "Universidade do Extremo Sul Catarinense (UNESC)",
			"Universidade do Grande ABC (UNIABC)", "Universidade do Grande Rio Professor José de Souza Herdy (UNIGRANRIO)",
			"Universidade do Norte (UNINORTE)", "Universidade do Oeste de Santa Catarina (UNOESC)",
			"Universidade do Oeste Paulista (UNOESTE)", "Universidade do Planalto Catarinense (UNIPLAC)",
			"Universidade do Sagrado Coração (USC)", "Universidade do Sul de Santa Catarina (UNISUL)",
			"Universidade do Tocantins (UNITINS)", "Universidade do Vale do Itajaí (UNIVALI)",
			"Universidade do Vale do Paraíba (UNIVAP)", "Universidade do Vale do Rio dos Sinos (UNISINOS)",
			"Universidade do Vale do Sapucaí (UNIVÁS)", "Universidade Estadual da Paraíba (UEPB)",
			"Universidade Estadual de Alagoas (UNEAL)", "Universidade Estadual de Campinas (UNICAMP)",
			"Universidade Estadual de Ciências da Saúde de Alagoas (UNCISAL)",
			"Universidade Estadual de Feira de Santana (UEFS)", "Universidade Estadual de Goiás (UEG)",
			"Universidade Estadual de Londrina (UEL)", "Universidade Estadual de Maringá (UEM)",
			"Universidade Estadual de Mato Grosso do Sul (UEMS)", "Universidade Estadual de Montes Claros (UNIMONTES)",
			"Universidade Estadual de Ponta Grossa (UEPG)", "Universidade Estadual de Roraima (UERR)",
			"Universidade Estadual de Santa Cruz (UESC)", "Universidade Estadual do Ceará (UECE)",
			"Universidade Estadual do Centro-Oeste (UNICENTRO)", "Universidade Estadual do Maranhão (UEMA)",
			"Universidade Estadual do Norte do Paraná (UENP)",
			"Universidade Estadual do Norte Fluminense Darcy Ribeiro (UENF)",
			"Universidade Estadual do Oeste do Paraná (UNIOESTE)", "Universidade Estadual do Paraná (UNESPAR)",
			"Universidade Estadual do Piauí (UESPI)", "Universidade Estadual do Rio Grande do Sul (UERGS)",
			"Universidade Estadual do Saber Tradicional da Amazônia (UESTA)",
			"Universidade Estadual do Sudoeste da Bahia (UESB)",
			"Universidade Estadual Paulista Júlio de Mesquita Filho (UNESP)",
			"Universidade Estadual Vale do Acaraú (UVA)",
			"Universidade Estácio de Sá (UNESA)", "Universidade Federal da Bahia (UFBA)",
			"Universidade Federal da Fronteira do Sul (UFFS)", "Universidade Federal da Grande Dourados (UFGD)",
			"Universidade Federal da Integração Internacional da Lusofonia Afro-Brasileira (UNILAB)",
			"Universidade Federal da Integração Latino-Americana (UNILA)", "Universidade Federal da Paraíba (UFPB)",
			"Universidade Federal de Alagoas (UFAL)", "Universidade Federal de Alfenas (UNIFAL)",
			"Universidade Federal de Campina Grande (UFCG)",
			"Universidade Federal de Ciências da Saúde de Porto Alegre (UFCSPA)", "Universidade Federal de Goiás (UFG)",
			"Universidade Federal de Itajubá (UNIFEI)", "Universidade Federal de Juiz de Fora (UFJF)",
			"Universidade Federal de Lavras (UFLA)", "Universidade Federal de Mato Grosso (UFMT)",
			"Universidade Federal de Mato Grosso do Sul (UFMS)", "Universidade Federal de Minas Gerais (UFMG)",
			"Universidade Federal de Ouro Preto (UFOP)", "Universidade Federal de Pelotas (UFPEL)",
			"Universidade Federal de Pernambuco (UFPE)", "Universidade Federal de Rondônia (UNIR)",
			"Universidade Federal de Roraima (UFRR)", "Universidade Federal de Santa Catarina (UFSC)",
			"Universidade Federal de Santa Maria (UFSM)", "Universidade Federal de Sergipe (UFS)",
			"Universidade Federal de São Carlos (UFSCar)", "Universidade Federal de São João del-Rei (UFSJ)",
			"Universidade Federal de São Paulo (UNIFESP)", "Universidade Federal de Uberlândia (UFU)",
			"Universidade Federal de Viçosa (UFV)", "Universidade Federal do ABC (UFABC)",
			"Universidade Federal do Acre (UFAC)", "Universidade Federal do Amapá (UNIFAP)",
			"Universidade Federal do Amazonas (UFAM)", "Universidade Federal do Ceará (UFC)",
			"Universidade Federal do Espírito Santo (UFES)", "Universidade Federal do Estado do Rio de Janeiro (UNIRIO)",
			"Universidade Federal do Maranhão (UFMA)", "Universidade Federal do Oeste da Bahia (UFOB)",
			"Universidade Federal do Oeste do Pará (UFOPA)", "Universidade Federal do Pampa (UNIPAMPA)",
			"Universidade Federal do Paraná (UFPR)", "Universidade Federal do Pará (UFPA)",
			"Universidade Federal do Piauí (UFPI)", "Universidade Federal do Recôncavo Baiano (UFRB)",
			"Universidade Federal do Rio de Janeiro (UFRJ)", "Universidade Federal do Rio Grande (FURG)",
			"Universidade Federal do Rio Grande do Norte (UFRN)", "Universidade Federal do Rio Grande do Sul (UFRGS)",
			"Universidade Federal do Sul da Bahia (UFSB)", "Universidade Federal do Sul e Sudeste do Pará (UNIFESSPA)",
			"Universidade Federal do Tocantins (UFT)", "Universidade Federal do Triângulo Mineiro (UFTM)",
			"Universidade Federal do Vale do São Francisco (UNIVASF)",
			"Universidade Federal dos Vales do Jequitinhonha e Mucuri (UFVJM)", "Universidade Federal Fluminense (UFF)",
			"Universidade Federal Rural da Amazônia (UFRA)", "Universidade Federal Rural de Pernambuco (UFRPE)",
			"Universidade Federal Rural do Rio de Janeiro (UFRRJ)", "Universidade Federal Rural do Semi-Árido (UFERSA)",
			"Universidade Feevale (FEEVALE)", "Universidade FUMEC (FUMEC)", "Universidade Gama Filho (UGF)",
			"Universidade Guarulhos (UNG)", "Universidade Ibirapuera (UNIB)", "Universidade Iguaçu (UNIG)",
			"Universidade José do Rosário Vellano (UNIFENAS)", "Universidade Luterana do Brasil (ULBRA)",
			"Universidade Metodista de Piracicaba (UNIMEP)", "Universidade Metodista de São Paulo (UMESP)",
			"Universidade Metropolitana de Santos (UNIMES)", "Universidade Municipal de São Caetano do Sul (USCS)",
			"Universidade Norte do Paraná (UNOPAR)", "Universidade Nove de Julho (UNINOVE)",
			"Universidade para o Desenvolvimento do Alto Vale do Itajaí (UNIDAVI)",
			"Universidade para o Desenvolvimento do Estado e da Região do Pantanal (UNIDERP)",
			"Universidade Paranaense (UNIPAR)", "Universidade Paulista (UNIP)", "Universidade Positivo (UP)",
			"Universidade Potiguar (UNP)", "Universidade Presbiteriana Mackenzie (MACKENZIE)",
			"Universidade Presidente Antônio Carlos (UNIPAC)", "Universidade Regional de Blumenau (FURB)",
			"Universidade Regional do Cariri (URCA)",
			"Universidade Regional do Noroeste do Estado do Rio Grande do Sul (UNIJUI)",
			"Universidade Regional Integrada do Alto Uruguai e das Missões (URI)",
			"Universidade Salgado de Oliveira (UNIVERSO)", "Universidade Salvador (UNIFACS)",
			"Universidade Santa Cecília (UNISANTA)", "Universidade Santa Úrsula (USU)",
			"Universidade Severino Sombra (USS)", "Universidade São Francisco (USF)",
			"Universidade São Judas Tadeu (USJT)", "Universidade São Marcos (USM)",
			"Universidade Tecnológica Federal do Paraná (UTFPR)", "Universidade Tiradentes (UNIT)",
			"Universidade Tuiuti do Paraná (UTP)", "Universidade Vale do Rio Doce (UNIVALE)",
			"Universidade Vale do Rio Verde (UNINCOR)", "Universidade Veiga de Almeida (UVA)",
			"Universidade Vila Velha (UVV)", "Univesidade do Alto Vale do Rio do Pixe (UNIARP)",
		},

		ZipFormats: []string{
			"#####-###",
		},
	}
}
