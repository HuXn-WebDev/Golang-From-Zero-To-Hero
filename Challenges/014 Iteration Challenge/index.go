package main

import "fmt"

func main() {
	vlcAuthors := []string {
	"Rémi Denis-Courmont",
	"Jean-Baptiste Kempf",
	"Laurent Aimar",
	"Gildas Bazin",
	"Felix Paul Kühne",
	"Rafaël Carré",
	"Pierre d'Herbemont",
	"Rémi Duraffort",
	"Derk-Jan Hartman",
	"Antoine Cellerier",
	"Samuel Hocevar",
	"Jean-Paul Saman",
	"Christophe Mutricy",
	"Clément Stenac",
	"Christophe Massiot",
	"Ilkka Ollakka",
	"Pierre Ynard",
	"François Cartegnie",
	"Damien Fouilleul",
	"Sigmund Augdal Helberg",
	"Erwan Tulou",
	"David Fuhrmann",
	"Olivier Teulière",
	"Cyril Deguet",
	"Eric Petit",
	"Filippo Carone",
	"Rocky Bernstein",
	"Hugo Beauzée-Luyssen",
	"Olivier Aubert",
	"Pavlov Konstantin",
	"Jakob Leben",
	"Benjamin Pracht",
	"Jean-Philippe André",
	"Steve Lhomme",
	"Stéphane Borel",
	"JP Dinger",
	"Geoffroy Couprie",
	"Martin Storsjö",
	"Marian Ďurkovič",
	"Ludovic Fauvet",
	"Yoann Peronneau",
	"Sébastien Escudier",
	"Jon Lech Johansen",
	"KO Myung-Hun",
	"Edward Wang",
	"Dennis van Amerongen",
	"Faustino Osuna",
	"Mirsal Ennaime",
	"Denis Charmet",
	"Jérôme Decoodt",
	"Loïc Minier",
	"David Flynn",
	"Frédéric Yhuel",
	"Kaarlo Raiha",
	"Mark Moriarty",
	"Christopher Mueller",
	"Fabio Ritrovato",
	"Tony Castley",
	"Srikanth Raju",
	"Michel Kaempf",
	"Jean-Marc Dressler",
	"Johan Bilien",
	"Vincent Seguin",
	"Simon Latapie",
	"Bernie Purcell",
	"Henri Fallon",
	"Sebastien Zwickert",
	"Christoph Miebach",
	"Adrien Maglo",
	"Emmanuel Puig",
	"Renaud Dartus",
	"Alexis de Lattre",
	"Vincent Penquerch",
	"Arnaud de Bossoreille de Ribou",
	"Mohammed Adnène Trojette",
	"Boris Dorès",
	"Jai Menon",
	"Anil Daoud",
	"Daniel Mierswa",
	"Naohiro Koriyama",
	"Rob Jonson",
	"Pierre Baillet",
	"Dominique Leuenberger",
	"Andre Pang",
	"Zoran Turalija",
	"Akash Mehrotra",
	"André Weber",
	"Anthony Loiseau",
	"Lukas Durfina",
	"Xavier Marchesini",
	"Cyril Mathé",
	"Devin Heitmueller",
	"Juho Vähä-Herttua",
	"Ken Self",
	"Alexis Ballier",
	"Juha Jeronen",
	"Nicolas Chauvet",
	"Richard Hosking",
	"Éric Lassauge",
	"Marc Ariberti",
	"Sébastien Toque",
	"Tobias Güntner",
	"Benoit Steiner",
	"Michel Lespinasse",
	"Carlo Calabrò",
	"Cheng Sun",
	"Michał Trzebiatowski",
	"Brad Smith",
	"Brendon Justin",
	"Alexey Sokolov",
	"Basos G",
	"Philippe Morin",
	"Steinar H. Gunderson",
	"Vicente Jimenez Aguilar",
	"Yuval Tze",
	"Yves Duret",
	"Benjamin Drung",
	"Michael Hanselmann",
	"Alex Merry",
	"Damien Lucas",
	"Grigori Goronzy",
	"Richard Shepherd",
	"Gaël Hendryckx",
	"Michael Feurstein",
	"Stephan Assmus",
	"Adrien Grand",
	"Colin Guthrie",
	"David Menestrina",
	"Dominique Martinet",
	"Gleb Pinigin",
	"Jason Luka",
	"Luc Saillard",
	"Luca Barbato",
	"Mario Speiß",
	"Pankaj Yadav",
	"Ramiro Polla",
	"Ronald Wright",
	"Rui Zhang",
	"Can Wu",
	"Christophe Courtaut",
	"FUJISAWA Tooru",
	"Hannes Domani",
	"Manol Manolov",
	"Timothy B. Terriberry",
	"Antoine Lejeune",
	"Arnaud Schauly",
	"Branko Kokanovic",
	"Dylan Yudaken",
	"Florian G. Pflug",
	"François Revol",
	"G Finch",
	"Keary Griffin",
	"Konstanty Bialkowski",
	"Ming Hu",
	"Philippe Coent",
	"Przemyslaw Fiala",
	"Tanguy Krotoff",
	"Vianney BOYER",
	"Casian Andrei",
	"Chris Smowton",
	"David Kaplan",
	"Eugenio Jarosiewicz",
	"Fabian Keil",
	"Guillaume Poussel",
	"John Peterson",
	"Justus Piater",
	"Mark Lee",
	"Martin T. H. Sandsmark",
	"Rune Botten",
	"Søren Bøg",
	"Toralf Niebuhr",
	"Tristan Matthews",
	"Angelo Haller",
	"Aurélien Nephtali",
	"Austin Burrow",
	"Bill C. Riemers",
	"Colin Delacroix",
	"Cristian Maglie",
	"Elminster2031",
	"Jakub Wieczorek",
	"John Freed",
	"Mark Hassman",
	"Martin Briza",
	"Mike Houben",
	"Romain Goyet",
	"Adrian Yanes",
	"Alexander Lakhin",
	"Anatoliy Anischovich",
	"Barry Wardell",
	"Ben Hutchings",
	"Besnard Jean-Baptiste",
	"Brian Weaver",
	"Clement Chesnin",
	"David Geldreich",
	"Diego Elio Pettenò",
	"Diego Fernando Nieto",
	"Georgi Chorbadzhiyski",
	"Jon Stacey",
	"Jonathan Rosser",
	"Joris van Rooij",
	"Kaloyan Kovachev",
	"Katsushi Kobayashi",
	"Kelly Anderson",
	"Loren Merritt",
	"Maciej Blizinski",
	"Mark Bidewell",
	"Miguel Angel Cabrera Moya",
	"Niles Bindel",
	"Samuel Pitoiset",
	"Scott Caudle",
	"Sean Robinson",
	"Sergey Radionov",
	"Simon Hailes",
	"Stephen Parry",
	"Sukrit Sangwan",
	"Thierry Reding",
	"Xavier Martin",
	"Alex Converse",
	"Alexander Bethke",
	"Alexandre Ratchov",
	"Andres Krapf",
	"Andrey Utkin",
	"Andri Pálsson",
	"Andy Chenee",
	"Anuradha Suraparaju",
	"Benjamin Poulain",
	"Brieuc Jeunhomme",
	"Chris Clayton",
	"Clément Lecigne",
	"Cédric Cocquebert",
	"Daniel Peng",
	"Danny Wood",
	"David K",
	"Edouard Gomez",
	"Emmanuel de Roux",
	"Frode Tennebø",
	"GBX",
	"Gaurav Narula",
	"Geraud CONTINSOUZAS",
	"Hugues Fruchet",
	"Jan Winter",
	"Jean-François Massol",
	"Jean-Philippe Grimaldi",
	"Josh Watzman",
	"Kai Lauterbach",
	"Konstantin Bogdanov",
	"Kuan-Chung Chiu",
	"Kuang Rufan",
	"Matthias Dahl",
	"Michael McEll",
	"Michael Ploujnikov",
	"Mike Schrag",
	"Nickolai Zeldovich",
	"Nicolas Bertrand",
	"Niklas Hayer",
	"Olafs Vandāns",
	"Olivier Gambier",
	"Paul Corke",
	"Ron Frederick",
	"Rov Juvano",
	"Sabourin Gilles",
	"Sam Lade",
	"Sandeep Kumar",
	"Sasha Koruga",
	"Sreng Jean",
	"Sven Petai",
	"Tomas Krotil",
	"Tomer Barletz",
	"Tristan Leteurtre",
	"Vittorio Giovara",
	"Wang Bo",
	"maxime Ripard",
	"xxcv",
	"Adam Hoka",
	"Adrian Knoth",
	"Adrien Cunin",
	"Alan Fischer",
	"Alan McCosh",
	"Alex Helfet",
	"Alexander Terentyev",
	"Alexandre Ferreira",
	"Alina Friedrichsen",
	"An L. Ber",
	"Andreas Schlick",
	"Andrew Schubert",
	"Andrey Makhnutin",
	"Arnaud Vallat",
	"Asad Mehmood",
	"Ashok Bhat",
	"Austin English",
	"Baptiste Coudurier",
	"Benoit Calvez",
	"Björn Stenberg",
	"Blake Livingston",
	"Brandon Brooks",
	"Brian Johnson",
	"Brian Kurle",
	"Cezar Elnazli",
	"Chris White",
	"Christian Masus",
	"Christoph Pfister",
	"Christoph Seibert",
	"Christopher Key",
	"Christopher Rath",
	"Claudio Ortelli",
	"Cody Russell",
	"Cristian Morales Vega",
	"Dan Rosenberg",
	"Daniel Marth",
	"Daniel Tisza",
	"Detlef Schroeder",
	"Diego Biurrun",
	"Dominik Rathann Mierzejewski",
	"Duncan Salerno",
	"Edward Sheldrake",
	"Elliot Murphy",
	"Eren Inan Canpolat",
	"Ernest E. Teem III",
	"Etienne Membrives",
	"Fargier Sylvain",
	"Fathi Boudra",
	"Felix Geyer",
	"Filipe Azevedo",
	"Finn Hughes",
	"Florian Hubold",
	"Florian Roeske",
	"Frank Enderle",
	"Frédéric Crozat",
	"Georg Seifert",
	"Gertjan Van Droogenbroeck",
	"Gilles Chanteperdrix",
	"Greg Farrell",
	"Gregory Maxwell",
	"Gwenole Beauchesne",
	"Götz Waschk",
	"Hans-Kristian Arntzen",
	"Harry Sintonen",
	"Iain Wade",
	"Ibraheem Paredath",
	"Isamu Arimoto",
	"Ismael Luceno",
	"James Bates",
	"James Bond",
	"James Turner",
	"Janne Grunau",
	"Janne Kujanpää",
	"Jarmo Torvinen",
	"Jason Scheunemann",
	"Jeff Lu",
	"Jeroen Ost",
	"Joe Taber",
	"Johann Ransay",
	"Johannes Weißl",
	"John Hendrikx",
	"John Stebbins",
	"Jonas Gehring",
	"Joseph S. Atkinson",
	"Juergen Lock",
	"Julien Lta BALLET",
	"Julien / Gellule",
	"Julien Humbert",
	"Kamil Baldyga",
	"Kamil Klimek",
	"Karlheinz Wohlmuth",
	"Kevin Anthony",
	"Kevin DuBois",
	"Lari Natri",
	"Lorenzo Pistone",
	"Lucas C. Villa Real",
	"Lukáš Lalinský",
	"Mal Graty",
	"Malte Tancred",
	"Martin Pöhlmann",
	"Martin Zeman",
	"Marton Balint",
	"Mathew King",
	"Mathieu Sonet",
	"Matthew A. Townsend",
	"Matthias Bauer",
	"Mika Tiainen",
	"Mike Cardillo",
	"Mounir Lamouri (volkmar)",
	"Natanael Copa",
	"Nathan Phillip Brink",
	"Nick Briggs",
	"Nick Pope",
	"Nil Geiswiller",
	"Pascal Thomet",
	"Pere Orga",
	"Peter Bak Nielsen",
	"Phil Roffe and David Grellscheid",
	"Philip Sequeira",
	"Pierre Souchay",
	"Piotr Fusik",
	"Pádraig Brady",
	"R.M",
	"Ralph Giles",
	"Ramon Gabarró",
	"Robert Forsman",
	"Robert Jedrzejczyk",
	"Robert Paciorek",
	"Rolf Ahrenberg",
	"Roman Pen",
	"Ruud Althuizen",
	"Samuli Suominen",
	"Scott Lyons",
	"Sebastian Birk",
	"Sergey Puzanov",
	"Sergio Ammirata",
	"Sharad Dixit",
	"Song Ye Wen",
	"Stephan Krempel",
	"Steven Kramer",
	"Steven Sheehy",
	"Sveinung Kvilhaugsvik",
	"Sylvain Cadhillac",
	"Sylver Bruneau",
	"Takahito HIRANO",
	"Theron Lewis",
	"Thijs Alkemade",
	"Tillmann Karras",
	"Timo Paulssen",
	"Timo Rothenpieler",
	"Tobias Rapp",
	"Tomasen",
	"Tony Vankrunkelsven",
	"Tristan Heaven",
	"Varphone Wong",
	"Vasily Fomin",
	"Vikram Narayanan",
	"Yannick Bréhon",
	"Yavor Doganov",
	"Yohann Martineau",
	"dharani.prabhu.s",
	"suheaven",
	"wucan",
	"김정은",
	"Adam Sampson",
	"Alexander Gall",
	"Alex Antropoff",
	"Alexis Guillard",
	"Alex Izvorski",
	"Amir Gouini",
	"Andrea Guzzo",
	"Andrew Flintham",
	"Andrew Zaikin",
	"Andy Lindsay",
	"Arai/Fujisawa Tooru",
	"Arkadiusz Miskiewicz",
	"Arnaud Gomes-do-Vale",
	"Arwed v. Merkatz",
	"Barak Ori",
	"Basil Achermann",
	"Benjamin Mironer",
	"Bill",
	"Bob Maguire",
	"Brian C. Wiles",
	"Brian Raymond",
	"Brian Robb",
	"Carsten Gottbehüt",
	"Carsten Haitzler",
	"Charles Hordis",
	"Chris Clepper",
	"Christian Henz",
	"Christof Baumgaertner",
	"Christophe Burgalat",
	"Christopher Johnson",
	"Cian Duffy",
	"Colin Simmonds",
	"Damian Ivereigh",
	"Daniel Fischer",
	"Daniel Stränger",
	"Danko Dolch",
	"Dennis Lou",
	"Dermot McGahon",
	"Douglas West",
	"Dugal Harris",
	"Emmanuel Blindauer",
	"Enrico Gueli",
	"Enrique Osuna",
	"Eren Türkay",
	"Eric Dudiak",
	"Espen Skoglund",
	"Ethan C. Baldridge",
	"François Seingier",
	"Frans van Veen",
	"Frédéric Ruget",
	"Gerald Hansink",
	"Gisle Vanem",
	"Glen Gray",
	"Goetz Waschk",
	"Gregory Hazel",
	"Gustaf Neumann",
	"Hang Su",
	"Hans Lambermont",
	"Hans-Peter Jansen",
	"Harris Dugal",
	"Heiko Panther",
	"Igor Helman",
	"Isaac Osunkunle",
	"Jan David Mol",
	"Jan Gerber",
	"Jan Van Boghout",
	"Jasper Alias",
	"Jean-Alexis Montignies",
	"Jean-Baptiste Le Stang",
	"Jeffrey Baker",
	"Jeroen Massar",
	"Jérôme Guilbaud",
	"Johannes Buchner",
	"Johen Michael Zorko",
	"Johnathan Rosser",
	"John Dalgliesh",
	"John Paul Lorenti",
	"Jörg",
	"Joseph Tulou",
	"Julien Blache",
	"Julien Plissonneau Duquène",
	"Julien Robert",
	"Kenneth Ostby",
	"Kenneth Self",
	"Kevin H. Patterson",
	"Koehler, Vitally",
	"K. Staring",
	"Lahiru Lakmal Priyadarshana",
	"Laurent Mutricy",
	"Leo Spalteholz",
	"Loox Thefuture",
	"Marc Nolette",
	"Marco Munderloh",
	"Mark Gritter",
	"Markus Kern",
	"Markus Kuespert",
	"Martin Hamrle",
	"Martin Kahr",
	"Mateus Krepsky Ludwich",
	"Mathias Kretschmer",
	"Mats Rojestal",
	"Matthias P. Nowak",
	"Matthieu Lochegnies",
	"Michael Mondragon",
	"Michael S. Feurstein",
	"Michel Lanners",
	"Mickael Hoerdt",
	"Miguel Angel Cabrera",
	"Mikko Hirvonen",
	"Moritz Bunkus",
	"Nilmoni Deb",
	"Olivier Houchard",
	"Olivier Pomel",
	"Ondrej Kuda aka Albert",
	"Øyvind Kolbu",
	"Pascal Levesque",
	"Patrick Horn",
	"Patrick McLean",
	"Pauline Castets",
	"Paul Mackerras",
	"Peter Surda",
	"Petr Vacek",
	"Philippe Van Hecke",
	"Pierre-Luc Beaudoin",
	"Pierre Marc Dumuid",
	"Régis Duchesne",
	"Remco Poortinga",
	"Rene Gollent",
	"Rob Casey",
	"Robson Braga Araujo",
	"Roine Gustafsson",
	"Roman Bednarek",
	"Rudolf Cornelissen",
	"Sašo Kiselkov",
	"Sebastian Jenny",
	"Shane Harper",
	"Stefán Freyr Stefánsson",
	"Steve Brown",
	"Steven M. Schultz",
	"Tapio Hiltunen",
	"Thomas L. Wood",
	"Thomas Mühlgrabner",
	"Thomas Parmelan",
	"Tim O Callagha",
	"Tim Schuerewegen",
	"Tong Ka Man",
	"Torsten Spindler",
	"Udo Richter",
	"Vincent Dimar",
	"Vincent Penne",
	"Vitalijus Slavinskas",
	"Vitaly V. Bursov",
	"Vladimir Chernyshov",
	"Wade Majors",
	"Wallace Wadge",
	"Watanabe Go",
	"William Hawkins",
	"Xavier Maillard",
	"Ye zhang",
	"Yuehua Zhao",
	}
	for index, value := range vlcAuthors {
		fmt.Println(index, "-", value)
	}
}