package config

// active only if ccontains at lest 1 item
var LocationPrefixWhitelist = []string{
	// следват метроро червената линия от центъра (без него) до края
	"град София, Младост 4",
	"град София, Малинова долина",
	"град София, Младост 2",
	"град София, Младост 3",
	"град София, Младост 1",
	"град София, Младост 1А",
	"град София, Дианабад",
	"град София, Изток",
	"град София, Дървеница",
	"град София, Мусагеница",
	"град София, Изгрев",
	"град София, Лозенец",

	// "град София, Център",
	// "град София, Овча купел", // covers the base, 1 and 2
}

var LocationPrefixBlacklist = []string{
	// "град София, Люлин", // moje bi ne trqbva taka da ignorirav vsi4ki lulini
	// "град София, с. Мало Бучино",
	// "град София, Орландовци",
	// "град София, Студентски град",
	// "град София, с. Лозен",
	// "град София, Банишора",
	// "град София, Овча купел",
	// "град София, Карпузица",
	// "град София, Обеля",

	// // izvun sofiq
	// "град София, Бояна",
	// "град София, м-т Гърдова глава",
	// "град София, Суходол",

	// // "град София, Суходол",
	// // "град София, Люлин 9",
	// // "град София, Модерно предградие",
	// // "град София, Обеля",
	// // "град София, Илинден",
	// // "град София, Симеоново",
	// // "град София, Фондови жилища",
	// // "град София, Овча купел 1",
	// // "град София, Овча купел 2",
	// // "град София, Толстой",
	// // "град София, Филиповци",
}

var LinkBlacklist = []string{
	////// za remont
	// "https://www.imot.bg/obiava-1c173679076198229-prodava-tristaen-apartament-grad-sofiya-tsentar-bul-slivnitsa",
	// "https://www.imot.bg/obiava-1c175005241220903-prodava-tristaen-apartament-grad-sofiya-tsentar",
	// "https://www.imot.bg/obiava-1b177789507250802-prodava-dvustaen-apartament-grad-sofiya-tsentar-ul-sofroniy-vrachanski",
	////// postroen no bez nishto vutre
	// NOTE: imashe edin nqkoi koito IMA GOTOVA banq (kato o4evidno ne e v tozi spisuk)
	"https://www.imot.bg/obiava-1b141233159470515-prodava-dvustaen-apartament-grad-sofiya-malinova-dolina",
	"https://www.imot.bg/obiava-1b163653993995564-prodava-dvustaen-apartament-grad-sofiya-malinova-dolina",
	"https://www.imot.bg/obiava-1b147505965186133-prodava-dvustaen-apartament-grad-sofiya-malinova-dolina",
	"https://www.imot.bg/obiava-1b148006413338001-prodava-dvustaen-apartament-grad-sofiya-malinova-dolina",
	"https://www.imot.bg/obiava-1b178055651466634-prodava-dvustaen-apartament-grad-sofiya-gorna-banya",
	"https://www.imot.bg/obiava-1b177434557094178-prodava-dvustaen-apartament-grad-sofiya-gorna-banya",
	"https://www.imot.bg/obiava-1b177771647069753-prodava-dvustaen-apartament-grad-sofiya-druzhba-1",
	"https://www.imot.bg/obiava-1b177002376361273-prodava-dvustaen-apartament-grad-sofiya-malinova-dolina",
	"https://www.imot.bg/obiava-1b177124711187292-prodava-dvustaen-apartament-grad-sofiya-malinova-dolina",
	"https://www.imot.bg/obiava-1b177912721664093-prodava-dvustaen-apartament-grad-sofiya-slaviya",
	"https://www.imot.bg/obiava-1b173894211764965-prodava-dvustaen-apartament-grad-sofiya-malinova-dolina",
	"https://www.imot.bg/obiava-1b176881562103330-prodava-dvustaen-apartament-grad-sofiya-malinova-dolina",
	"https://www.imot.bg/obiava-1b175205951125310-prodava-dvustaen-apartament-grad-sofiya-malinova-dolina",
	"https://www.imot.bg/obiava-1b175467121497416-prodava-dvustaen-apartament-grad-sofiya-malinova-dolina",
	"https://www.imot.bg/obiava-1b177974020865607-prodava-dvustaen-apartament-grad-sofiya-malinova-dolina",
	"https://www.imot.bg/obiava-1b177160338356172-prodava-dvustaen-apartament-grad-sofiya-malinova-dolina",
	"https://www.imot.bg/obiava-1b177307100109691-prodava-dvustaen-apartament-grad-sofiya-malinova-dolina",
	"https://www.imot.bg/obiava-1b178056291621258-prodava-dvustaen-apartament-grad-sofiya-malinova-dolina",
	"https://www.imot.bg/obiava-1a148033075659734-prodava-ednostaen-apartament-grad-sofiya-malinova-dolina",
	"https://www.imot.bg/obiava-1a147506110948376-prodava-ednostaen-apartament-grad-sofiya-malinova-dolina",
	////// nqma snimki otvutre / cgi snimki / nqma snimki
	"https://www.imot.bg/obiava-1b177754380218825-prodava-dvustaen-apartament-grad-sofiya-suhata-reka",
	"https://www.imot.bg/obiava-1b178031461376813-prodava-dvustaen-apartament-grad-sofiya-malinova-dolina",
	"https://www.imot.bg/obiava-1b177970053215072-prodava-dvustaen-apartament-grad-sofiya-darvenitsa-bul-andrey-lyapchev",
	"https://www.imot.bg/obiava-1b177999892723954-prodava-dvustaen-apartament-grad-sofiya-mladost-3",
	"https://www.imot.bg/obiava-1b177305343027129-prodava-dvustaen-apartament-grad-sofiya-malinova-dolina",
	"https://www.imot.bg/obiava-1b176770861199951-prodava-dvustaen-apartament-grad-sofiya-malinova-dolina",
}
