package config

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
	"https://www.imot.bg/obiava-1b178152527658636-prodava-dvustaen-apartament-grad-sofiya-moderno-predgradie",
	"https://www.imot.bg/obiava-1b175335919126899-prodava-dvustaen-apartament-grad-sofiya-vitosha",
	"https://www.imot.bg/obiava-1b177029734818751-prodava-dvustaen-apartament-grad-sofiya-knyazhevo",
	"https://www.imot.bg/obiava-1b178118565599729-prodava-dvustaen-apartament-grad-sofiya-serdika",
	"https://www.imot.bg/obiava-1b178005823784927-prodava-dvustaen-apartament-grad-sofiya-vrabnitsa-2-ul-asen-traykov",

	////// nqma/loshi/cgi snimki

	"https://www.imot.bg/obiava-1b177754380218825-prodava-dvustaen-apartament-grad-sofiya-suhata-reka",
	"https://www.imot.bg/obiava-1b178031461376813-prodava-dvustaen-apartament-grad-sofiya-malinova-dolina",
	"https://www.imot.bg/obiava-1b177970053215072-prodava-dvustaen-apartament-grad-sofiya-darvenitsa-bul-andrey-lyapchev",
	"https://www.imot.bg/obiava-1b177999892723954-prodava-dvustaen-apartament-grad-sofiya-mladost-3",
	"https://www.imot.bg/obiava-1b177305343027129-prodava-dvustaen-apartament-grad-sofiya-malinova-dolina",
	"https://www.imot.bg/obiava-1b176770861199951-prodava-dvustaen-apartament-grad-sofiya-malinova-dolina",
	"https://www.imot.bg/obiava-1b178047036137842-prodava-dvustaen-apartament-grad-sofiya-malinova-dolina",
	"https://www.imot.bg/obiava-1b177278400657202-prodava-dvustaen-apartament-grad-sofiya-nadezhda-2",
	"https://www.imot.bg/obiava-1c177089876512090-prodava-tristaen-apartament-grad-sofiya-tsentar-ul-tsar-samuil",
	"https://www.imot.bg/obiava-1b177392098560761-prodava-dvustaen-apartament-grad-sofiya-vrazhdebna",
	"https://www.imot.bg/obiava-1b177986885222097-prodava-dvustaen-apartament-grad-sofiya-benkovski-ul-bodra-smyana",
	"https://www.imot.bg/obiava-1b178151382152646-prodava-dvustaen-apartament-grad-sofiya-serdika",
	"https://www.imot.bg/obiava-1b177904866444144-prodava-dvustaen-apartament-grad-sofiya-tsentar-bul-knyaginya-mariya-luiza",
	"https://www.imot.bg/obiava-1b177694961544093-prodava-dvustaen-apartament-grad-sofiya-hadzhi-dimitar",
	"https://www.imot.bg/obiava-1b177202603689393-prodava-dvustaen-apartament-grad-sofiya-druzhba-1",

	//// super s4upen - beyond remont

	"https://www.imot.bg/obiava-1d175465115813246-prodava-chetiristaen-apartament-grad-sofiya-tsentar",
	"https://www.imot.bg/obiava-1b175005994592048-prodava-dvustaen-apartament-grad-sofiya-tsentar",
	"https://www.imot.bg/obiava-1b175005241220903-prodava-dvustaen-apartament-grad-sofiya-tsentar",
	"https://www.imot.bg/obiava-1c175465099419565-prodava-tristaen-apartament-grad-sofiya-tsentar",
	"https://www.imot.bg/obiava-1d177020934881050-prodava-chetiristaen-apartament-grad-sofiya-oborishte",

	//// lokaciqta e luja

	"https://www.imot.bg/obiava-1c177978613173915-prodava-tristaen-apartament-grad-sofiya-druzhba-2",
	"https://www.imot.bg/obiava-1c166434162354666-prodava-tristaen-apartament-grad-sofiya-7-mi-11-ti-kilometar",

	//// ako iska nqkoi moje da mi razbie prozoreca i da vleze ez

	"https://www.imot.bg/obiava-1b176880848839354-prodava-dvustaen-apartament-grad-sofiya-vitosha",
	"https://www.imot.bg/obiava-1b178031327895576-prodava-dvustaen-apartament-grad-sofiya-vitosha",

	//// other

	"https://www.imot.bg/obiava-1b177133940645532-prodava-dvustaen-apartament-grad-sofiya-lagera-ul-gebedzhe", // parter + za remont + skupo
}
