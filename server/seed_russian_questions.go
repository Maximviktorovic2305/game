package main

import (
	"fmt"
	"log"
	"server/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Database connection string - adjust these values to match your database
	dsn := "host=localhost user=postgres password=admin dbname=mill port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	// Connect to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Connected to database successfully")

	// Run migrations to ensure tables exist
	err = db.AutoMigrate(
		&models.Question{},
		&models.Option{},
		&models.GameSession{},
	)
	if err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	fmt.Println("Database migrations completed")

	// Clear existing questions and options
	db.Exec("DELETE FROM options")
	db.Exec("DELETE FROM questions")

	fmt.Println("Cleared existing questions and options")

	// Define prize amounts for each level
	prizes := []int{100, 200, 300, 500, 1000, 2000, 4000, 8000, 16000, 32000, 64000, 125000, 250000, 500000, 1000000}

	// Define questions data in Russian - 10 questions for each level
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
			Text:  "Какая самая большая планета Солнечной системы?",
			Level: 2,
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
			Text:  "Какой самый твердый минерал на Земле?",
			Level: 3,
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
			Text:  "Какая самая большая планета Солнечной системы?",
			Level: 4,
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
			Level: 4,
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
			Level: 4,
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
			Level: 4,
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
			Level: 4,
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
			Level: 4,
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
			Level: 4,
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
			Level: 4,
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
			Level: 4,
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
			Text:  "Какая птица является символом мудрости?",
			Level: 5,
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
			Level: 5,
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
			Level: 5,
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
			Level: 5,
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
			Text:  "Какая самая большая планета Солнечной системы?",
			Level: 6,
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
			Level: 6,
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
			Level: 6,
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
			Level: 6,
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
			Level: 6,
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
			Level: 6,
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
			Level: 6,
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
			Level: 6,
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
			Level: 6,
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
		{
			Text:  "Какая самая большая змея в мире?",
			Level: 7,
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
			Text:  "Какая самая большая планета Солнечной системы?",
			Level: 8,
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
			Level: 8,
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
			Level: 8,
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
			Level: 8,
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
			Level: 8,
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
			Level: 8,
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
			Level: 8,
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
			Level: 8,
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
			Level: 8,
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
		{
			Text:  "Какая самая большая змея в мире?",
			Level: 9,
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
		{
			Text:  "Какая самая большая змея в мире?",
			Level: 10,
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
		{
			Text:  "Какая самая большая змея в мире?",
			Level: 11,
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
		{
			Text:  "Какая самая большая змея в мире?",
			Level: 12,
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
		{
			Text:  "Какая самая большая змея в мире?",
			Level: 13,
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
		{
			Text:  "Какая самая большая змея в мире?",
			Level: 14,
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
		{
			Text:  "Какая самая большая змея в мире?",
			Level: 15,
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
	}

	// Create questions and options
	for i, qd := range questionsData {
		question := models.Question{
			Text:  qd.Text,
			Level: qd.Level,
			Prize: prizes[qd.Level-1], // Use the prize for the corresponding level
		}

		// Create the question first
		if err := db.Create(&question).Error; err != nil {
			log.Printf("Error creating question %d: %v", i+1, err)
			continue
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
				log.Printf("Error creating option for question %d: %v", i+1, err)
				continue
			}
		}

		fmt.Printf("Added question %d: %s\n", i+1, qd.Text)
	}

	fmt.Println("Successfully seeded database with Russian questions!")
}
