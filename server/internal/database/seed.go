package database

import (
	"server/internal/models"

	"gorm.io/gorm"
)

func SeedQuestions(db *gorm.DB) error {
	return SeedQuestionsForce(db, false)
}

func SeedQuestionsForce(db *gorm.DB, force bool) error {
	// Check if questions already exist
	var count int64
	db.Model(&models.Question{}).Count(&count)
	if count > 0 && !force {
		return nil // Questions already seeded
	}

	// If force is true, delete all existing questions and options
	if force {
		db.Exec("DELETE FROM options")
		db.Exec("DELETE FROM questions")
	}

	// Define prize amounts for each level
	prizes := []int{100, 200, 300, 500, 1000, 2000, 4000, 8000, 16000, 32000, 64000, 125000, 250000, 500000, 1000000}

	// Define questions data in Russian - 10 unique questions for each level
	questionsData := []struct {
		Text    string
		Level   int
		Options []struct {
			Letter    string
			Text      string
			IsCorrect bool
		}
	}{
		// Level 1 Questions (10)
		{
			Text:  "Сколько дней в неделе?",
			Level: 1,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "5", IsCorrect: false},
				{Letter: "B", Text: "6", IsCorrect: false},
				{Letter: "C", Text: "7", IsCorrect: true},
				{Letter: "D", Text: "8", IsCorrect: false},
			},
		},
		{
			Text:  "Какой цвет неба в ясный день?",
			Level: 1,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Зеленый", IsCorrect: false},
				{Letter: "B", Text: "Красный", IsCorrect: false},
				{Letter: "C", Text: "Синий", IsCorrect: true},
				{Letter: "D", Text: "Желтый", IsCorrect: false},
			},
		},
		{
			Text:  "Сколько месяцев в году?",
			Level: 1,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "10", IsCorrect: false},
				{Letter: "B", Text: "11", IsCorrect: false},
				{Letter: "C", Text: "12", IsCorrect: true},
				{Letter: "D", Text: "13", IsCorrect: false},
			},
		},
		{
			Text:  "Сколько ног у кошки?",
			Level: 1,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "2", IsCorrect: false},
				{Letter: "B", Text: "3", IsCorrect: false},
				{Letter: "C", Text: "4", IsCorrect: true},
				{Letter: "D", Text: "5", IsCorrect: false},
			},
		},
		{
			Text:  "Какой самый большой океан на Земле?",
			Level: 1,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Атлантический", IsCorrect: false},
				{Letter: "B", Text: "Индийский", IsCorrect: false},
				{Letter: "C", Text: "Тихий", IsCorrect: true},
				{Letter: "D", Text: "Северный Ледовитый", IsCorrect: false},
			},
		},
		{
			Text:  "Сколько букв в русском алфавите?",
			Level: 1,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "30", IsCorrect: false},
				{Letter: "B", Text: "31", IsCorrect: false},
				{Letter: "C", Text: "33", IsCorrect: true},
				{Letter: "D", Text: "35", IsCorrect: false},
			},
		},
		{
			Text:  "Сколько часов в сутках?",
			Level: 1,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "12", IsCorrect: false},
				{Letter: "B", Text: "18", IsCorrect: false},
				{Letter: "C", Text: "24", IsCorrect: true},
				{Letter: "D", Text: "30", IsCorrect: false},
			},
		},
		{
			Text:  "Сколько минут в одном часе?",
			Level: 1,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "30", IsCorrect: false},
				{Letter: "B", Text: "45", IsCorrect: false},
				{Letter: "C", Text: "60", IsCorrect: true},
				{Letter: "D", Text: "90", IsCorrect: false},
			},
		},
		{
			Text:  "Сколько глаз у человека?",
			Level: 1,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "1", IsCorrect: false},
				{Letter: "B", Text: "2", IsCorrect: true},
				{Letter: "C", Text: "3", IsCorrect: false},
				{Letter: "D", Text: "4", IsCorrect: false},
			},
		},
		{
			Text:  "Сколько зубов у взрослого человека?",
			Level: 1,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "28", IsCorrect: false},
				{Letter: "B", Text: "30", IsCorrect: false},
				{Letter: "C", Text: "32", IsCorrect: true},
				{Letter: "D", Text: "34", IsCorrect: false},
			},
		},

		// Level 2 Questions (10)
		{
			Text:  "Какой химический элемент обозначается символом 'H'?",
			Level: 2,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Гелий", IsCorrect: false},
				{Letter: "B", Text: "Водород", IsCorrect: true},
				{Letter: "C", Text: "Ртуть", IsCorrect: false},
				{Letter: "D", Text: "Гафний", IsCorrect: false},
			},
		},
		{
			Text:  "Какой самый большой внутренний орган у человека?",
			Level: 2,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Сердце", IsCorrect: false},
				{Letter: "B", Text: "Печень", IsCorrect: true},
				{Letter: "C", Text: "Легкие", IsCorrect: false},
				{Letter: "D", Text: "Мозг", IsCorrect: false},
			},
		},
		{
			Text:  "Какая столица Франции?",
			Level: 2,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Лондон", IsCorrect: false},
				{Letter: "B", Text: "Берлин", IsCorrect: false},
				{Letter: "C", Text: "Париж", IsCorrect: true},
				{Letter: "D", Text: "Рим", IsCorrect: false},
			},
		},
		{
			Text:  "Какой металл самый тяжелый?",
			Level: 2,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Золото", IsCorrect: false},
				{Letter: "B", Text: "Свинец", IsCorrect: false},
				{Letter: "C", Text: "Осмий", IsCorrect: true},
				{Letter: "D", Text: "Платина", IsCorrect: false},
			},
		},
		{
			Text:  "Какой самый высокий водопад в мире?",
			Level: 2,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Ниагарский", IsCorrect: false},
				{Letter: "B", Text: "Виктория", IsCorrect: false},
				{Letter: "C", Text: "Анхель", IsCorrect: true},
				{Letter: "D", Text: "Игуасу", IsCorrect: false},
			},
		},
		{
			Text:  "Какой самый большой материк на Земле?",
			Level: 2,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Африка", IsCorrect: false},
				{Letter: "B", Text: "Северная Америка", IsCorrect: false},
				{Letter: "C", Text: "Евразия", IsCorrect: true},
				{Letter: "D", Text: "Австралия", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая длинная река в мире?",
			Level: 2,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Амазонка", IsCorrect: false},
				{Letter: "B", Text: "Миссисипи", IsCorrect: false},
				{Letter: "C", Text: "Нил", IsCorrect: true},
				{Letter: "D", Text: "Янцзы", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая большая пустыня в мире?",
			Level: 2,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Сахара", IsCorrect: false},
				{Letter: "B", Text: "Гоби", IsCorrect: false},
				{Letter: "C", Text: "Антарктическая", IsCorrect: true},
				{Letter: "D", Text: "Калахари", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая большая страна в мире по площади?",
			Level: 2,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Канада", IsCorrect: false},
				{Letter: "B", Text: "США", IsCorrect: false},
				{Letter: "C", Text: "Россия", IsCorrect: true},
				{Letter: "D", Text: "Китай", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая высокая гора в мире?",
			Level: 2,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Килиманджаро", IsCorrect: false},
				{Letter: "B", Text: "Эверест", IsCorrect: true},
				{Letter: "C", Text: "Эльбрус", IsCorrect: false},
				{Letter: "D", Text: "Аконкагуа", IsCorrect: false},
			},
		},

		// Level 3 Questions (10)
		{
			Text:  "Какой химический элемент обозначается символом 'O'?",
			Level: 3,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Золото", IsCorrect: false},
				{Letter: "B", Text: "Кислород", IsCorrect: true},
				{Letter: "C", Text: "Олово", IsCorrect: false},
				{Letter: "D", Text: "Осмий", IsCorrect: false},
			},
		},
		{
			Text:  "Какая столица Японии?",
			Level: 3,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Пекин", IsCorrect: false},
				{Letter: "B", Text: "Сеул", IsCorrect: false},
				{Letter: "C", Text: "Токио", IsCorrect: true},
				{Letter: "D", Text: "Бангкок", IsCorrect: false},
			},
		},
		{
			Text:  "Какая птица является символом мудрости?",
			Level: 3,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Сова", IsCorrect: true},
				{Letter: "B", Text: "Орел", IsCorrect: false},
				{Letter: "C", Text: "Сокол", IsCorrect: false},
				{Letter: "D", Text: "Ворона", IsCorrect: false},
			},
		},
		{
			Text:  "Какой металл жидкий при комнатной температуре?",
			Level: 3,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Железо", IsCorrect: false},
				{Letter: "B", Text: "Золото", IsCorrect: false},
				{Letter: "C", Text: "Ртуть", IsCorrect: true},
				{Letter: "D", Text: "Свинец", IsCorrect: false},
			},
		},
		{
			Text:  "Какой газ преобладает в атмосфере Земли?",
			Level: 3,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Кислород", IsCorrect: false},
				{Letter: "B", Text: "Углекислый газ", IsCorrect: false},
				{Letter: "C", Text: "Азот", IsCorrect: true},
				{Letter: "D", Text: "Водород", IsCorrect: false},
			},
		},
		{
			Text:  "Какой самый большой орган у человека?",
			Level: 3,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Сердце", IsCorrect: false},
				{Letter: "B", Text: "Мозг", IsCorrect: false},
				{Letter: "C", Text: "Кожа", IsCorrect: true},
				{Letter: "D", Text: "Печень", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая большая птица в мире?",
			Level: 3,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Кондор", IsCorrect: false},
				{Letter: "B", Text: "Страус", IsCorrect: true},
				{Letter: "C", Text: "Альбатрос", IsCorrect: false},
				{Letter: "D", Text: "Орел", IsCorrect: false},
			},
		},
		{
			Text:  "Какое самое большое млекопитающее в мире?",
			Level: 3,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Слон", IsCorrect: false},
				{Letter: "B", Text: "Синий кит", IsCorrect: true},
				{Letter: "C", Text: "Жираф", IsCorrect: false},
				{Letter: "D", Text: "Бегемот", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая большая змея в мире?",
			Level: 3,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Удав", IsCorrect: false},
				{Letter: "B", Text: "Анаконда", IsCorrect: true},
				{Letter: "C", Text: "Питон", IsCorrect: false},
				{Letter: "D", Text: "Кобра", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая большая кошка в мире?",
			Level: 3,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Лев", IsCorrect: false},
				{Letter: "B", Text: "Тигр", IsCorrect: true},
				{Letter: "C", Text: "Леопард", IsCorrect: false},
				{Letter: "D", Text: "Ягуар", IsCorrect: false},
			},
		},

		// Level 4 Questions (10)
		{
			Text:  "Какой химический элемент обозначается символом 'Fe'?",
			Level: 4,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Фтор", IsCorrect: false},
				{Letter: "B", Text: "Железо", IsCorrect: true},
				{Letter: "C", Text: "Фермий", IsCorrect: false},
				{Letter: "D", Text: "Фосфор", IsCorrect: false},
			},
		},
		{
			Text:  "Какая столица Германии?",
			Level: 4,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Париж", IsCorrect: false},
				{Letter: "B", Text: "Вена", IsCorrect: false},
				{Letter: "C", Text: "Берлин", IsCorrect: true},
				{Letter: "D", Text: "Цюрих", IsCorrect: false},
			},
		},
		{
			Text:  "Какой самый большой остров в мире?",
			Level: 4,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Мадагаскар", IsCorrect: false},
				{Letter: "B", Text: "Борнео", IsCorrect: false},
				{Letter: "C", Text: "Гренландия", IsCorrect: true},
				{Letter: "D", Text: "Новая Гвинея", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая длинная река в Европе?",
			Level: 4,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Дунай", IsCorrect: false},
				{Letter: "B", Text: "Рейн", IsCorrect: false},
				{Letter: "C", Text: "Волга", IsCorrect: true},
				{Letter: "D", Text: "Днепр", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая высокая гора в Европе?",
			Level: 4,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Монблан", IsCorrect: false},
				{Letter: "B", Text: "Эльбрус", IsCorrect: true},
				{Letter: "C", Text: "Маттерхорн", IsCorrect: false},
				{Letter: "D", Text: "Юнгфрау", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая большая пустыня в Азии?",
			Level: 4,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Гоби", IsCorrect: true},
				{Letter: "B", Text: "Каракумы", IsCorrect: false},
				{Letter: "C", Text: "Кызылкум", IsCorrect: false},
				{Letter: "D", Text: "Такла-Макан", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая большая страна в Африке?",
			Level: 4,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Эфиопия", IsCorrect: false},
				{Letter: "B", Text: "Египет", IsCorrect: false},
				{Letter: "C", Text: "Алжир", IsCorrect: true},
				{Letter: "D", Text: "Судан", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая длинная река в Африке?",
			Level: 4,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Конго", IsCorrect: false},
				{Letter: "B", Text: "Нил", IsCorrect: false},
				{Letter: "C", Text: "Нигер", IsCorrect: false},
				{Letter: "D", Text: "Замбези", IsCorrect: true},
			},
		},
		{
			Text:  "Какой самый большой океан на Земле?",
			Level: 4,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Атлантический", IsCorrect: false},
				{Letter: "B", Text: "Индийский", IsCorrect: false},
				{Letter: "C", Text: "Тихий", IsCorrect: true},
				{Letter: "D", Text: "Северный Ледовитый", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая большая пустыня в мире?",
			Level: 4,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Сахара", IsCorrect: false},
				{Letter: "B", Text: "Гоби", IsCorrect: false},
				{Letter: "C", Text: "Антарктическая", IsCorrect: true},
				{Letter: "D", Text: "Калахари", IsCorrect: false},
			},
		},

		// Level 5 Questions (10)
		{
			Text:  "Какой химический элемент обозначается символом 'Au'?",
			Level: 5,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Серебро", IsCorrect: false},
				{Letter: "B", Text: "Золото", IsCorrect: true},
				{Letter: "C", Text: "Алюминий", IsCorrect: false},
				{Letter: "D", Text: "Аргон", IsCorrect: false},
			},
		},
		{
			Text:  "Какая столица Италии?",
			Level: 5,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Милан", IsCorrect: false},
				{Letter: "B", Text: "Флоренция", IsCorrect: false},
				{Letter: "C", Text: "Рим", IsCorrect: true},
				{Letter: "D", Text: "Неаполь", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая большая планета Солнечной системы?",
			Level: 5,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Сатурн", IsCorrect: false},
				{Letter: "B", Text: "Юпитер", IsCorrect: true},
				{Letter: "C", Text: "Нептун", IsCorrect: false},
				{Letter: "D", Text: "Уран", IsCorrect: false},
			},
		},
		{
			Text:  "Какой самый твердый минерал на Земле?",
			Level: 5,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Золото", IsCorrect: false},
				{Letter: "B", Text: "Железо", IsCorrect: false},
				{Letter: "C", Text: "Алмаз", IsCorrect: true},
				{Letter: "D", Text: "Платина", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая большая клетка в человеческом теле?",
			Level: 5,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Красная кровяная клетка", IsCorrect: false},
				{Letter: "B", Text: "Белая кровяная клетка", IsCorrect: false},
				{Letter: "C", Text: "Нервная клетка", IsCorrect: false},
				{Letter: "D", Text: "Яйцеклетка", IsCorrect: true},
			},
		},
		{
			Text:  "Какой самый длинный мост в мире?",
			Level: 5,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Золотые ворота", IsCorrect: false},
				{Letter: "B", Text: "Босфор", IsCorrect: false},
				{Letter: "C", Text: "Дамба в Хайфаре", IsCorrect: true},
				{Letter: "D", Text: "Тамис", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая большая птица в мире?",
			Level: 5,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Кондор", IsCorrect: false},
				{Letter: "B", Text: "Страус", IsCorrect: true},
				{Letter: "C", Text: "Альбатрос", IsCorrect: false},
				{Letter: "D", Text: "Орел", IsCorrect: false},
			},
		},
		{
			Text:  "Какое самое большое млекопитающее в мире?",
			Level: 5,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Слон", IsCorrect: false},
				{Letter: "B", Text: "Синий кит", IsCorrect: true},
				{Letter: "C", Text: "Жираф", IsCorrect: false},
				{Letter: "D", Text: "Бегемот", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая большая змея в мире?",
			Level: 5,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Удав", IsCorrect: false},
				{Letter: "B", Text: "Анаконда", IsCorrect: true},
				{Letter: "C", Text: "Питон", IsCorrect: false},
				{Letter: "D", Text: "Кобра", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая большая кошка в мире?",
			Level: 5,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Лев", IsCorrect: false},
				{Letter: "B", Text: "Тигр", IsCorrect: true},
				{Letter: "C", Text: "Леопард", IsCorrect: false},
				{Letter: "D", Text: "Ягуар", IsCorrect: false},
			},
		},

		// Level 6 Questions (10)
		{
			Text:  "Какой химический элемент обозначается символом 'Ag'?",
			Level: 6,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Золото", IsCorrect: false},
				{Letter: "B", Text: "Серебро", IsCorrect: true},
				{Letter: "C", Text: "Аргон", IsCorrect: false},
				{Letter: "D", Text: "Алюминий", IsCorrect: false},
			},
		},
		{
			Text:  "Какая столица Австралии?",
			Level: 6,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Сидней", IsCorrect: false},
				{Letter: "B", Text: "Мельбурн", IsCorrect: false},
				{Letter: "C", Text: "Канберра", IsCorrect: true},
				{Letter: "D", Text: "Брисбен", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая длинная река в мире?",
			Level: 6,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Амазонка", IsCorrect: false},
				{Letter: "B", Text: "Миссисипи", IsCorrect: false},
				{Letter: "C", Text: "Нил", IsCorrect: true},
				{Letter: "D", Text: "Янцзы", IsCorrect: false},
			},
		},
		{
			Text:  "Какой самый высокий водопад в мире?",
			Level: 6,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Ниагарский", IsCorrect: false},
				{Letter: "B", Text: "Виктория", IsCorrect: false},
				{Letter: "C", Text: "Анхель", IsCorrect: true},
				{Letter: "D", Text: "Игуасу", IsCorrect: false},
			},
		},
		{
			Text:  "Какой самый большой материк на Земле?",
			Level: 6,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Африка", IsCorrect: false},
				{Letter: "B", Text: "Северная Америка", IsCorrect: false},
				{Letter: "C", Text: "Евразия", IsCorrect: true},
				{Letter: "D", Text: "Австралия", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая большая пустыня в мире?",
			Level: 6,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Сахара", IsCorrect: false},
				{Letter: "B", Text: "Гоби", IsCorrect: false},
				{Letter: "C", Text: "Антарктическая", IsCorrect: true},
				{Letter: "D", Text: "Калахари", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая большая страна в мире по площади?",
			Level: 6,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Канада", IsCorrect: false},
				{Letter: "B", Text: "США", IsCorrect: false},
				{Letter: "C", Text: "Россия", IsCorrect: true},
				{Letter: "D", Text: "Китай", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая высокая гора в мире?",
			Level: 6,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Килиманджаро", IsCorrect: false},
				{Letter: "B", Text: "Эверест", IsCorrect: true},
				{Letter: "C", Text: "Эльбрус", IsCorrect: false},
				{Letter: "D", Text: "Аконкагуа", IsCorrect: false},
			},
		},
		{
			Text:  "Какой самый большой океан на Земле?",
			Level: 6,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Атлантический", IsCorrect: false},
				{Letter: "B", Text: "Индийский", IsCorrect: false},
				{Letter: "C", Text: "Тихий", IsCorrect: true},
				{Letter: "D", Text: "Северный Ледовитый", IsCorrect: false},
			},
		},
		{
			Text:  "Сколько зубов у взрослого человека?",
			Level: 6,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "28", IsCorrect: false},
				{Letter: "B", Text: "30", IsCorrect: false},
				{Letter: "C", Text: "32", IsCorrect: true},
				{Letter: "D", Text: "34", IsCorrect: false},
			},
		},

		// Level 7 Questions (10)
		{
			Text:  "Какой химический элемент обозначается символом 'Na'?",
			Level: 7,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Натрий", IsCorrect: true},
				{Letter: "B", Text: "Никель", IsCorrect: false},
				{Letter: "C", Text: "Ниобий", IsCorrect: false},
				{Letter: "D", Text: "Неон", IsCorrect: false},
			},
		},
		{
			Text:  "Какая столица Канады?",
			Level: 7,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Торонто", IsCorrect: false},
				{Letter: "B", Text: "Ванкувер", IsCorrect: false},
				{Letter: "C", Text: "Оттава", IsCorrect: true},
				{Letter: "D", Text: "Монреаль", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая большая планета Солнечной системы?",
			Level: 7,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Сатурн", IsCorrect: false},
				{Letter: "B", Text: "Юпитер", IsCorrect: true},
				{Letter: "C", Text: "Нептун", IsCorrect: false},
				{Letter: "D", Text: "Уран", IsCorrect: false},
			},
		},
		{
			Text:  "Какой самый твердый минерал на Земле?",
			Level: 7,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Золото", IsCorrect: false},
				{Letter: "B", Text: "Железо", IsCorrect: false},
				{Letter: "C", Text: "Алмаз", IsCorrect: true},
				{Letter: "D", Text: "Платина", IsCorrect: false},
			},
		},
		{
			Text:  "Какая птица является символом мудрости?",
			Level: 7,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Сова", IsCorrect: true},
				{Letter: "B", Text: "Орел", IsCorrect: false},
				{Letter: "C", Text: "Сокол", IsCorrect: false},
				{Letter: "D", Text: "Ворона", IsCorrect: false},
			},
		},
		{
			Text:  "Какой металл жидкий при комнатной температуре?",
			Level: 7,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Железо", IsCorrect: false},
				{Letter: "B", Text: "Золото", IsCorrect: false},
				{Letter: "C", Text: "Ртуть", IsCorrect: true},
				{Letter: "D", Text: "Свинец", IsCorrect: false},
			},
		},
		{
			Text:  "Какой газ преобладает в атмосфере Земли?",
			Level: 7,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Кислород", IsCorrect: false},
				{Letter: "B", Text: "Углекислый газ", IsCorrect: false},
				{Letter: "C", Text: "Азот", IsCorrect: true},
				{Letter: "D", Text: "Водород", IsCorrect: false},
			},
		},
		{
			Text:  "Какой самый большой орган у человека?",
			Level: 7,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Сердце", IsCorrect: false},
				{Letter: "B", Text: "Мозг", IsCorrect: false},
				{Letter: "C", Text: "Кожа", IsCorrect: true},
				{Letter: "D", Text: "Печень", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая большая птица в мире?",
			Level: 7,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Кондор", IsCorrect: false},
				{Letter: "B", Text: "Страус", IsCorrect: true},
				{Letter: "C", Text: "Альбатрос", IsCorrect: false},
				{Letter: "D", Text: "Орел", IsCorrect: false},
			},
		},
		{
			Text:  "Какое самое большое млекопитающее в мире?",
			Level: 7,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Слон", IsCorrect: false},
				{Letter: "B", Text: "Синий кит", IsCorrect: true},
				{Letter: "C", Text: "Жираф", IsCorrect: false},
				{Letter: "D", Text: "Бегемот", IsCorrect: false},
			},
		},

		// Level 8 Questions (10)
		{
			Text:  "Какой химический элемент обозначается символом 'Cu'?",
			Level: 8,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Кальций", IsCorrect: false},
				{Letter: "B", Text: "Кадмий", IsCorrect: false},
				{Letter: "C", Text: "Медь", IsCorrect: true},
				{Letter: "D", Text: "Цезий", IsCorrect: false},
			},
		},
		{
			Text:  "Какая столица Бразилии?",
			Level: 8,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Сан-Паулу", IsCorrect: false},
				{Letter: "B", Text: "Рио-де-Жанейро", IsCorrect: false},
				{Letter: "C", Text: "Бразилиа", IsCorrect: true},
				{Letter: "D", Text: "Салвадор", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая длинная река в мире?",
			Level: 8,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Амазонка", IsCorrect: false},
				{Letter: "B", Text: "Миссисипи", IsCorrect: false},
				{Letter: "C", Text: "Нил", IsCorrect: true},
				{Letter: "D", Text: "Янцзы", IsCorrect: false},
			},
		},
		{
			Text:  "Какой самый высокий водопад в мире?",
			Level: 8,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Ниагарский", IsCorrect: false},
				{Letter: "B", Text: "Виктория", IsCorrect: false},
				{Letter: "C", Text: "Анхель", IsCorrect: true},
				{Letter: "D", Text: "Игуасу", IsCorrect: false},
			},
		},
		{
			Text:  "Какой самый большой материк на Земле?",
			Level: 8,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Африка", IsCorrect: false},
				{Letter: "B", Text: "Северная Америка", IsCorrect: false},
				{Letter: "C", Text: "Евразия", IsCorrect: true},
				{Letter: "D", Text: "Австралия", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая большая пустыня в мире?",
			Level: 8,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Сахара", IsCorrect: false},
				{Letter: "B", Text: "Гоби", IsCorrect: false},
				{Letter: "C", Text: "Антарктическая", IsCorrect: true},
				{Letter: "D", Text: "Калахари", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая большая страна в мире по площади?",
			Level: 8,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Канада", IsCorrect: false},
				{Letter: "B", Text: "США", IsCorrect: false},
				{Letter: "C", Text: "Россия", IsCorrect: true},
				{Letter: "D", Text: "Китай", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая высокая гора в мире?",
			Level: 8,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Килиманджаро", IsCorrect: false},
				{Letter: "B", Text: "Эверест", IsCorrect: true},
				{Letter: "C", Text: "Эльбрус", IsCorrect: false},
				{Letter: "D", Text: "Аконкагуа", IsCorrect: false},
			},
		},
		{
			Text:  "Какой самый большой океан на Земле?",
			Level: 8,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Атлантический", IsCorrect: false},
				{Letter: "B", Text: "Индийский", IsCorrect: false},
				{Letter: "C", Text: "Тихий", IsCorrect: true},
				{Letter: "D", Text: "Северный Ледовитый", IsCorrect: false},
			},
		},
		{
			Text:  "Сколько зубов у взрослого человека?",
			Level: 8,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "28", IsCorrect: false},
				{Letter: "B", Text: "30", IsCorrect: false},
				{Letter: "C", Text: "32", IsCorrect: true},
				{Letter: "D", Text: "34", IsCorrect: false},
			},
		},

		// Level 9 Questions (10)
		{
			Text:  "Какой химический элемент обозначается символом 'Pb'?",
			Level: 9,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Платина", IsCorrect: false},
				{Letter: "B", Text: "Полоний", IsCorrect: false},
				{Letter: "C", Text: "Свинец", IsCorrect: true},
				{Letter: "D", Text: "Протактиний", IsCorrect: false},
			},
		},
		{
			Text:  "Какая столица Индии?",
			Level: 9,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Мумбаи", IsCorrect: false},
				{Letter: "B", Text: "Бангалор", IsCorrect: false},
				{Letter: "C", Text: "Нью-Дели", IsCorrect: true},
				{Letter: "D", Text: "Калькутта", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая большая планета Солнечной системы?",
			Level: 9,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Сатурн", IsCorrect: false},
				{Letter: "B", Text: "Юпитер", IsCorrect: true},
				{Letter: "C", Text: "Нептун", IsCorrect: false},
				{Letter: "D", Text: "Уран", IsCorrect: false},
			},
		},
		{
			Text:  "Какой самый твердый минерал на Земле?",
			Level: 9,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Золото", IsCorrect: false},
				{Letter: "B", Text: "Железо", IsCorrect: false},
				{Letter: "C", Text: "Алмаз", IsCorrect: true},
				{Letter: "D", Text: "Платина", IsCorrect: false},
			},
		},
		{
			Text:  "Какая птица является символом мудрости?",
			Level: 9,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Сова", IsCorrect: true},
				{Letter: "B", Text: "Орел", IsCorrect: false},
				{Letter: "C", Text: "Сокол", IsCorrect: false},
				{Letter: "D", Text: "Ворона", IsCorrect: false},
			},
		},
		{
			Text:  "Какой металл жидкий при комнатной температуре?",
			Level: 9,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Железо", IsCorrect: false},
				{Letter: "B", Text: "Золото", IsCorrect: false},
				{Letter: "C", Text: "Ртуть", IsCorrect: true},
				{Letter: "D", Text: "Свинец", IsCorrect: false},
			},
		},
		{
			Text:  "Какой газ преобладает в атмосфере Земли?",
			Level: 9,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Кислород", IsCorrect: false},
				{Letter: "B", Text: "Углекислый газ", IsCorrect: false},
				{Letter: "C", Text: "Азот", IsCorrect: true},
				{Letter: "D", Text: "Водород", IsCorrect: false},
			},
		},
		{
			Text:  "Какой самый большой орган у человека?",
			Level: 9,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Сердце", IsCorrect: false},
				{Letter: "B", Text: "Мозг", IsCorrect: false},
				{Letter: "C", Text: "Кожа", IsCorrect: true},
				{Letter: "D", Text: "Печень", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая большая птица в мире?",
			Level: 9,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Кондор", IsCorrect: false},
				{Letter: "B", Text: "Страус", IsCorrect: true},
				{Letter: "C", Text: "Альбатрос", IsCorrect: false},
				{Letter: "D", Text: "Орел", IsCorrect: false},
			},
		},
		{
			Text:  "Какое самое большое млекопитающее в мире?",
			Level: 9,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Слон", IsCorrect: false},
				{Letter: "B", Text: "Синий кит", IsCorrect: true},
				{Letter: "C", Text: "Жираф", IsCorrect: false},
				{Letter: "D", Text: "Бегемот", IsCorrect: false},
			},
		},

		// Level 10 Questions (10)
		{
			Text:  "Какой химический элемент обозначается символом 'Hg'?",
			Level: 10,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Гафний", IsCorrect: false},
				{Letter: "B", Text: "Германий", IsCorrect: false},
				{Letter: "C", Text: "Ртуть", IsCorrect: true},
				{Letter: "D", Text: "Гольмий", IsCorrect: false},
			},
		},
		{
			Text:  "Какая столица Мексики?",
			Level: 10,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Гвадалахара", IsCorrect: false},
				{Letter: "B", Text: "Монтеррей", IsCorrect: false},
				{Letter: "C", Text: "Мехико", IsCorrect: true},
				{Letter: "D", Text: "Канкун", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая большая планета Солнечной системы?",
			Level: 10,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Сатурн", IsCorrect: false},
				{Letter: "B", Text: "Юпитер", IsCorrect: true},
				{Letter: "C", Text: "Нептун", IsCorrect: false},
				{Letter: "D", Text: "Уран", IsCorrect: false},
			},
		},
		{
			Text:  "Какой самый твердый минерал на Земле?",
			Level: 10,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Золото", IsCorrect: false},
				{Letter: "B", Text: "Железо", IsCorrect: false},
				{Letter: "C", Text: "Алмаз", IsCorrect: true},
				{Letter: "D", Text: "Платина", IsCorrect: false},
			},
		},
		{
			Text:  "Какая птица является символом мудрости?",
			Level: 10,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Сова", IsCorrect: true},
				{Letter: "B", Text: "Орел", IsCorrect: false},
				{Letter: "C", Text: "Сокол", IsCorrect: false},
				{Letter: "D", Text: "Ворона", IsCorrect: false},
			},
		},
		{
			Text:  "Какой металл жидкий при комнатной температуре?",
			Level: 10,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Железо", IsCorrect: false},
				{Letter: "B", Text: "Золото", IsCorrect: false},
				{Letter: "C", Text: "Ртуть", IsCorrect: true},
				{Letter: "D", Text: "Свинец", IsCorrect: false},
			},
		},
		{
			Text:  "Какой газ преобладает в атмосфере Земли?",
			Level: 10,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Кислород", IsCorrect: false},
				{Letter: "B", Text: "Углекислый газ", IsCorrect: false},
				{Letter: "C", Text: "Азот", IsCorrect: true},
				{Letter: "D", Text: "Водород", IsCorrect: false},
			},
		},
		{
			Text:  "Какой самый большой орган у человека?",
			Level: 10,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Сердце", IsCorrect: false},
				{Letter: "B", Text: "Мозг", IsCorrect: false},
				{Letter: "C", Text: "Кожа", IsCorrect: true},
				{Letter: "D", Text: "Печень", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая большая птица в мире?",
			Level: 10,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Кондор", IsCorrect: false},
				{Letter: "B", Text: "Страус", IsCorrect: true},
				{Letter: "C", Text: "Альбатрос", IsCorrect: false},
				{Letter: "D", Text: "Орел", IsCorrect: false},
			},
		},
		{
			Text:  "Какое самое большое млекопитающее в мире?",
			Level: 10,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Слон", IsCorrect: false},
				{Letter: "B", Text: "Синий кит", IsCorrect: true},
				{Letter: "C", Text: "Жираф", IsCorrect: false},
				{Letter: "D", Text: "Бегемот", IsCorrect: false},
			},
		},

		// Level 11 Questions (10)
		{
			Text:  "Какой химический элемент обозначается символом 'Sn'?",
			Level: 11,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Скандий", IsCorrect: false},
				{Letter: "B", Text: "Селен", IsCorrect: false},
				{Letter: "C", Text: "Олово", IsCorrect: true},
				{Letter: "D", Text: "Самарий", IsCorrect: false},
			},
		},
		{
			Text:  "Какая столица Аргентины?",
			Level: 11,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Кордова", IsCorrect: false},
				{Letter: "B", Text: "Буэнос-Айрес", IsCorrect: true},
				{Letter: "C", Text: "Росарио", IsCorrect: false},
				{Letter: "D", Text: "Мендоса", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая большая планета Солнечной системы?",
			Level: 11,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Сатурн", IsCorrect: false},
				{Letter: "B", Text: "Юпитер", IsCorrect: true},
				{Letter: "C", Text: "Нептун", IsCorrect: false},
				{Letter: "D", Text: "Уран", IsCorrect: false},
			},
		},
		{
			Text:  "Какой самый твердый минерал на Земле?",
			Level: 11,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Золото", IsCorrect: false},
				{Letter: "B", Text: "Железо", IsCorrect: false},
				{Letter: "C", Text: "Алмаз", IsCorrect: true},
				{Letter: "D", Text: "Платина", IsCorrect: false},
			},
		},
		{
			Text:  "Какая птица является символом мудрости?",
			Level: 11,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Сова", IsCorrect: true},
				{Letter: "B", Text: "Орел", IsCorrect: false},
				{Letter: "C", Text: "Сокол", IsCorrect: false},
				{Letter: "D", Text: "Ворона", IsCorrect: false},
			},
		},
		{
			Text:  "Какой металл жидкий при комнатной температуре?",
			Level: 11,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Железо", IsCorrect: false},
				{Letter: "B", Text: "Золото", IsCorrect: false},
				{Letter: "C", Text: "Ртуть", IsCorrect: true},
				{Letter: "D", Text: "Свинец", IsCorrect: false},
			},
		},
		{
			Text:  "Какой газ преобладает в атмосфере Земли?",
			Level: 11,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Кислород", IsCorrect: false},
				{Letter: "B", Text: "Углекислый газ", IsCorrect: false},
				{Letter: "C", Text: "Азот", IsCorrect: true},
				{Letter: "D", Text: "Водород", IsCorrect: false},
			},
		},
		{
			Text:  "Какой самый большой орган у человека?",
			Level: 11,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Сердце", IsCorrect: false},
				{Letter: "B", Text: "Мозг", IsCorrect: false},
				{Letter: "C", Text: "Кожа", IsCorrect: true},
				{Letter: "D", Text: "Печень", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая большая птица в мире?",
			Level: 11,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Кондор", IsCorrect: false},
				{Letter: "B", Text: "Страус", IsCorrect: true},
				{Letter: "C", Text: "Альбатрос", IsCorrect: false},
				{Letter: "D", Text: "Орел", IsCorrect: false},
			},
		},
		{
			Text:  "Какое самое большое млекопитающее в мире?",
			Level: 11,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Слон", IsCorrect: false},
				{Letter: "B", Text: "Синий кит", IsCorrect: true},
				{Letter: "C", Text: "Жираф", IsCorrect: false},
				{Letter: "D", Text: "Бегемот", IsCorrect: false},
			},
		},

		// Level 12 Questions (10)
		{
			Text:  "Какой химический элемент обозначается символом 'W'?",
			Level: 12,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Ванадий", IsCorrect: false},
				{Letter: "B", Text: "Вольфрам", IsCorrect: true},
				{Letter: "C", Text: "Радий", IsCorrect: false},
				{Letter: "D", Text: "Рубидий", IsCorrect: false},
			},
		},
		{
			Text:  "Какая столица Египта?",
			Level: 12,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Александрия", IsCorrect: false},
				{Letter: "B", Text: "Каир", IsCorrect: true},
				{Letter: "C", Text: "Луксор", IsCorrect: false},
				{Letter: "D", Text: "Асуан", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая большая планета Солнечной системы?",
			Level: 12,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Сатурн", IsCorrect: false},
				{Letter: "B", Text: "Юпитер", IsCorrect: true},
				{Letter: "C", Text: "Нептун", IsCorrect: false},
				{Letter: "D", Text: "Уран", IsCorrect: false},
			},
		},
		{
			Text:  "Какой самый твердый минерал на Земле?",
			Level: 12,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Золото", IsCorrect: false},
				{Letter: "B", Text: "Железо", IsCorrect: false},
				{Letter: "C", Text: "Алмаз", IsCorrect: true},
				{Letter: "D", Text: "Платина", IsCorrect: false},
			},
		},
		{
			Text:  "Какая птица является символом мудрости?",
			Level: 12,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Сова", IsCorrect: true},
				{Letter: "B", Text: "Орел", IsCorrect: false},
				{Letter: "C", Text: "Сокол", IsCorrect: false},
				{Letter: "D", Text: "Ворона", IsCorrect: false},
			},
		},
		{
			Text:  "Какой металл жидкий при комнатной температуре?",
			Level: 12,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Железо", IsCorrect: false},
				{Letter: "B", Text: "Золото", IsCorrect: false},
				{Letter: "C", Text: "Ртуть", IsCorrect: true},
				{Letter: "D", Text: "Свинец", IsCorrect: false},
			},
		},
		{
			Text:  "Какой газ преобладает в атмосфере Земли?",
			Level: 12,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Кислород", IsCorrect: false},
				{Letter: "B", Text: "Углекислый газ", IsCorrect: false},
				{Letter: "C", Text: "Азот", IsCorrect: true},
				{Letter: "D", Text: "Водород", IsCorrect: false},
			},
		},
		{
			Text:  "Какой самый большой орган у человека?",
			Level: 12,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Сердце", IsCorrect: false},
				{Letter: "B", Text: "Мозг", IsCorrect: false},
				{Letter: "C", Text: "Кожа", IsCorrect: true},
				{Letter: "D", Text: "Печень", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая большая птица в мире?",
			Level: 12,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Кондор", IsCorrect: false},
				{Letter: "B", Text: "Страус", IsCorrect: true},
				{Letter: "C", Text: "Альбатрос", IsCorrect: false},
				{Letter: "D", Text: "Орел", IsCorrect: false},
			},
		},
		{
			Text:  "Какое самое большое млекопитающее в мире?",
			Level: 12,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Слон", IsCorrect: false},
				{Letter: "B", Text: "Синий кит", IsCorrect: true},
				{Letter: "C", Text: "Жираф", IsCorrect: false},
				{Letter: "D", Text: "Бегемот", IsCorrect: false},
			},
		},

		// Level 13 Questions (10)
		{
			Text:  "Какой химический элемент обозначается символом 'Pt'?",
			Level: 13,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Полоний", IsCorrect: false},
				{Letter: "B", Text: "Протактиний", IsCorrect: false},
				{Letter: "C", Text: "Платина", IsCorrect: true},
				{Letter: "D", Text: "Плутоний", IsCorrect: false},
			},
		},
		{
			Text:  "Какая столица ЮАР?",
			Level: 13,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Кейптаун", IsCorrect: false},
				{Letter: "B", Text: "Йоханнесбург", IsCorrect: false},
				{Letter: "C", Text: "Претория", IsCorrect: true},
				{Letter: "D", Text: "Дурбан", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая большая планета Солнечной системы?",
			Level: 13,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Сатурн", IsCorrect: false},
				{Letter: "B", Text: "Юпитер", IsCorrect: true},
				{Letter: "C", Text: "Нептун", IsCorrect: false},
				{Letter: "D", Text: "Уран", IsCorrect: false},
			},
		},
		{
			Text:  "Какой самый твердый минерал на Земле?",
			Level: 13,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Золото", IsCorrect: false},
				{Letter: "B", Text: "Железо", IsCorrect: false},
				{Letter: "C", Text: "Алмаз", IsCorrect: true},
				{Letter: "D", Text: "Платина", IsCorrect: false},
			},
		},
		{
			Text:  "Какая птица является символом мудрости?",
			Level: 13,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Сова", IsCorrect: true},
				{Letter: "B", Text: "Орел", IsCorrect: false},
				{Letter: "C", Text: "Сокол", IsCorrect: false},
				{Letter: "D", Text: "Ворона", IsCorrect: false},
			},
		},
		{
			Text:  "Какой металл жидкий при комнатной температуре?",
			Level: 13,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Железо", IsCorrect: false},
				{Letter: "B", Text: "Золото", IsCorrect: false},
				{Letter: "C", Text: "Ртуть", IsCorrect: true},
				{Letter: "D", Text: "Свинец", IsCorrect: false},
			},
		},
		{
			Text:  "Какой газ преобладает в атмосфере Земли?",
			Level: 13,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Кислород", IsCorrect: false},
				{Letter: "B", Text: "Углекислый газ", IsCorrect: false},
				{Letter: "C", Text: "Азот", IsCorrect: true},
				{Letter: "D", Text: "Водород", IsCorrect: false},
			},
		},
		{
			Text:  "Какой самый большой орган у человека?",
			Level: 13,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Сердце", IsCorrect: false},
				{Letter: "B", Text: "Мозг", IsCorrect: false},
				{Letter: "C", Text: "Кожа", IsCorrect: true},
				{Letter: "D", Text: "Печень", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая большая птица в мире?",
			Level: 13,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Кондор", IsCorrect: false},
				{Letter: "B", Text: "Страус", IsCorrect: true},
				{Letter: "C", Text: "Альбатрос", IsCorrect: false},
				{Letter: "D", Text: "Орел", IsCorrect: false},
			},
		},
		{
			Text:  "Какое самое большое млекопитающее в мире?",
			Level: 13,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Слон", IsCorrect: false},
				{Letter: "B", Text: "Синий кит", IsCorrect: true},
				{Letter: "C", Text: "Жираф", IsCorrect: false},
				{Letter: "D", Text: "Бегемот", IsCorrect: false},
			},
		},

		// Level 14 Questions (10)
		{
			Text:  "Какой химический элемент обозначается символом 'U'?",
			Level: 14,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Углерод", IsCorrect: false},
				{Letter: "B", Text: "Уран", IsCorrect: true},
				{Letter: "C", Text: "Углерод", IsCorrect: false},
				{Letter: "D", Text: "Уран", IsCorrect: false},
			},
		},
		{
			Text:  "Какая столица Швеции?",
			Level: 14,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Копенгаген", IsCorrect: false},
				{Letter: "B", Text: "Осло", IsCorrect: false},
				{Letter: "C", Text: "Стокгольм", IsCorrect: true},
				{Letter: "D", Text: "Хельсинки", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая большая планета Солнечной системы?",
			Level: 14,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Сатурн", IsCorrect: false},
				{Letter: "B", Text: "Юпитер", IsCorrect: true},
				{Letter: "C", Text: "Нептун", IsCorrect: false},
				{Letter: "D", Text: "Уран", IsCorrect: false},
			},
		},
		{
			Text:  "Какой самый твердый минерал на Земле?",
			Level: 14,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Золото", IsCorrect: false},
				{Letter: "B", Text: "Железо", IsCorrect: false},
				{Letter: "C", Text: "Алмаз", IsCorrect: true},
				{Letter: "D", Text: "Платина", IsCorrect: false},
			},
		},
		{
			Text:  "Какая птица является символом мудрости?",
			Level: 14,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Сова", IsCorrect: true},
				{Letter: "B", Text: "Орел", IsCorrect: false},
				{Letter: "C", Text: "Сокол", IsCorrect: false},
				{Letter: "D", Text: "Ворона", IsCorrect: false},
			},
		},
		{
			Text:  "Какой металл жидкий при комнатной температуре?",
			Level: 14,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Железо", IsCorrect: false},
				{Letter: "B", Text: "Золото", IsCorrect: false},
				{Letter: "C", Text: "Ртуть", IsCorrect: true},
				{Letter: "D", Text: "Свинец", IsCorrect: false},
			},
		},
		{
			Text:  "Какой газ преобладает в атмосфере Земли?",
			Level: 14,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Кислород", IsCorrect: false},
				{Letter: "B", Text: "Углекислый газ", IsCorrect: false},
				{Letter: "C", Text: "Азот", IsCorrect: true},
				{Letter: "D", Text: "Водород", IsCorrect: false},
			},
		},
		{
			Text:  "Какой самый большой орган у человека?",
			Level: 14,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Сердце", IsCorrect: false},
				{Letter: "B", Text: "Мозг", IsCorrect: false},
				{Letter: "C", Text: "Кожа", IsCorrect: true},
				{Letter: "D", Text: "Печень", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая большая птица в мире?",
			Level: 14,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Кондор", IsCorrect: false},
				{Letter: "B", Text: "Страус", IsCorrect: true},
				{Letter: "C", Text: "Альбатрос", IsCorrect: false},
				{Letter: "D", Text: "Орел", IsCorrect: false},
			},
		},
		{
			Text:  "Какое самое большое млекопитающее в мире?",
			Level: 14,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Слон", IsCorrect: false},
				{Letter: "B", Text: "Синий кит", IsCorrect: true},
				{Letter: "C", Text: "Жираф", IsCorrect: false},
				{Letter: "D", Text: "Бегемот", IsCorrect: false},
			},
		},

		// Level 15 Questions (10)
		{
			Text:  "Какой химический элемент обозначается символом 'Au'?",
			Level: 15,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Серебро", IsCorrect: false},
				{Letter: "B", Text: "Золото", IsCorrect: true},
				{Letter: "C", Text: "Алюминий", IsCorrect: false},
				{Letter: "D", Text: "Аргон", IsCorrect: false},
			},
		},
		{
			Text:  "Какая столица Норвегии?",
			Level: 15,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Копенгаген", IsCorrect: false},
				{Letter: "B", Text: "Осло", IsCorrect: false},
				{Letter: "C", Text: "Стокгольм", IsCorrect: false},
				{Letter: "D", Text: "Осло", IsCorrect: true},
			},
		},
		{
			Text:  "Какая самая большая планета Солнечной системы?",
			Level: 15,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Сатурн", IsCorrect: false},
				{Letter: "B", Text: "Юпитер", IsCorrect: true},
				{Letter: "C", Text: "Нептун", IsCorrect: false},
				{Letter: "D", Text: "Уран", IsCorrect: false},
			},
		},
		{
			Text:  "Какой самый твердый минерал на Земле?",
			Level: 15,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Золото", IsCorrect: false},
				{Letter: "B", Text: "Железо", IsCorrect: false},
				{Letter: "C", Text: "Алмаз", IsCorrect: true},
				{Letter: "D", Text: "Платина", IsCorrect: false},
			},
		},
		{
			Text:  "Какая птица является символом мудрости?",
			Level: 15,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Сова", IsCorrect: true},
				{Letter: "B", Text: "Орел", IsCorrect: false},
				{Letter: "C", Text: "Сокол", IsCorrect: false},
				{Letter: "D", Text: "Ворона", IsCorrect: false},
			},
		},
		{
			Text:  "Какой металл жидкий при комнатной температуре?",
			Level: 15,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Железо", IsCorrect: false},
				{Letter: "B", Text: "Золото", IsCorrect: false},
				{Letter: "C", Text: "Ртуть", IsCorrect: true},
				{Letter: "D", Text: "Свинец", IsCorrect: false},
			},
		},
		{
			Text:  "Какой газ преобладает в атмосфере Земли?",
			Level: 15,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Кислород", IsCorrect: false},
				{Letter: "B", Text: "Углекислый газ", IsCorrect: false},
				{Letter: "C", Text: "Азот", IsCorrect: true},
				{Letter: "D", Text: "Водород", IsCorrect: false},
			},
		},
		{
			Text:  "Какой самый большой орган у человека?",
			Level: 15,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Сердце", IsCorrect: false},
				{Letter: "B", Text: "Мозг", IsCorrect: false},
				{Letter: "C", Text: "Кожа", IsCorrect: true},
				{Letter: "D", Text: "Печень", IsCorrect: false},
			},
		},
		{
			Text:  "Какая самая большая птица в мире?",
			Level: 15,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Кондор", IsCorrect: false},
				{Letter: "B", Text: "Страус", IsCorrect: true},
				{Letter: "C", Text: "Альбатрос", IsCorrect: false},
				{Letter: "D", Text: "Орел", IsCorrect: false},
			},
		},
		{
			Text:  "Какое самое большое млекопитающее в мире?",
			Level: 15,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Слон", IsCorrect: false},
				{Letter: "B", Text: "Синий кит", IsCorrect: true},
				{Letter: "C", Text: "Жираф", IsCorrect: false},
				{Letter: "D", Text: "Бегемот", IsCorrect: false},
			},
		},
	}

	// Create questions and options
	for _, qd := range questionsData {
		question := models.Question{
			Text:  qd.Text,
			Level: qd.Level,
			Prize: prizes[qd.Level-1], // Use the prize for the corresponding level
		}

		// Create the question first
		if err := db.Create(&question).Error; err != nil {
			return err
		}

		// Create options for this question
		for _, od := range qd.Options {
			option := models.Option{
				QuestionID: question.ID,
				Letter:     od.Letter,
				Text:       od.Text,
				IsCorrect:  od.IsCorrect,
			}
			if err := db.Create(&option).Error; err != nil {
				return err
			}
		}
	}

	return nil
}