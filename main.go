package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	// 1.	Buat variabel dengan tipe data  sebagai berikut:
	// 		name (string)
	// 		age(int)
	// 		height(float)
	// 		IsMaried (boolean)
	// 		interestInCoding (Boolean)
	// 		tampilkan tipe datanya dan nilainya

	fmt.Println("======================= 1. Variables and Data Types of GoLang =======================")
	var name string = "Farhan Nur Hakim"
	var age int = 22
	var height float64 = 171.2
	var isMarried bool = false
	var interestInCoding bool = true

	fmt.Printf("Name: %s (%s)\n", name, reflect.TypeOf(name))
	fmt.Printf("Age: %d (%s)\n", age, reflect.TypeOf(age))
	fmt.Printf("Height: %.2f (%s)\n", height, reflect.TypeOf(height))
	fmt.Printf("Is Married: %t (%s)\n", isMarried, reflect.TypeOf(isMarried))
	fmt.Printf("Interest in coding: %t (%s)\n\n", interestInCoding, reflect.TypeOf(interestInCoding))

	// 2.	Dari soal no 1 ubahlah
	// 		dari int ke string
	// 		dari string ke int
	// 		dari float ke string
	// 		dari string ke float
	// 		tampilkan tipe datanya dan nilainya

	fmt.Println("======================= 2. Data Types Conversion =======================")

	fmt.Println("Convert from string to int")
	intName, errInt := strconv.Atoi(name)
	if errInt != nil {
		fmt.Println("Failed to convert string to int ", errInt)
	} else {
		fmt.Printf("Name: %d (%s)\n", intName, reflect.TypeOf(intName))
	}

	fmt.Println("Convert from string to float")
	floatName, errFloat := strconv.ParseFloat(name, 64)
	if errFloat != nil {
		fmt.Println("Failed to convert string to float ", errFloat)
	} else {
		fmt.Printf("Name: %f (%s)\n", floatName, reflect.TypeOf(floatName))
	}

	fmt.Println("Convert from int to string")
	stringAge := strconv.Itoa(age)
	fmt.Printf("Age: %s (%s)\n", stringAge, reflect.TypeOf(stringAge))

	fmt.Println("Convert from float to string")
	stringHeight := strconv.FormatFloat(float64(height), 'f', 4, 64)
	fmt.Printf("Height: %s (%s)\n\n", stringHeight, reflect.TypeOf(stringHeight))

	// 3.	Buat program yang menghitung rata-rata UN beserta gradenya, dengan mengisi 4 nilai yakni Bahasa indonesia, Bahasa Inggris Matematika dan IPA,
	// 		yang di dalam program tersebut memiliki validasi yaitu semua nilai tersebut harus diisi, dan juga untuk grade memiliki kondisi dengan ketentuan
	// 		sebagai berikut:
	// 		90 - 100 = A
	// 		80 - 89 = B
	// 		70 - 79 = C
	// 		60 - 69 = D
	// 		0 - 59 = E

	fmt.Println("======================= 3. Basic grading program =======================")
	var bahasaScoreStr, englishScoreStr, mathScoreStr, naturalScienceScoreStr string
	var bahasaScore, englishScore, mathScore, naturalScienceScore int

	for {
		fmt.Print("Input bahasa score: ")
		fmt.Scan(&bahasaScoreStr)
		tempBahasa, errIntBahasa := strconv.Atoi(bahasaScoreStr)
		bahasaScore = tempBahasa
		if bahasaScore > 100 || bahasaScore < 0 {
			fmt.Println("Range score should be between 0 and 100")
		} else if errIntBahasa != nil {
			fmt.Println("Value data type should be integer ", errIntBahasa)
		} else {
			break
		}
	}

	for {
		fmt.Print("Input english score: ")
		fmt.Scan(&englishScoreStr)
		tempEnglish, errIntEnglish := strconv.Atoi(englishScoreStr)
		englishScore = tempEnglish
		if englishScore > 100 || englishScore < 0 {
			fmt.Println("Range score should be between 0 and 100")
		} else if errIntEnglish != nil {
			fmt.Println("Value data type should be integer ", errIntEnglish)
		} else {
			break
		}
	}

	for {
		fmt.Print("Input math score: ")
		fmt.Scan(&mathScoreStr)
		tempMath, errIntMath := strconv.Atoi(mathScoreStr)
		mathScore = tempMath
		if mathScore > 100 || mathScore < 0 {
			fmt.Println("Range score should be between 0 and 100")
		} else if errIntMath != nil {
			fmt.Println("Value data type should be integer ", errIntMath)
		} else {
			break
		}
	}

	for {
		fmt.Print("Input natural science score: ")
		fmt.Scan(&naturalScienceScoreStr)
		tempNaturalScience, errIntNaturalScience := strconv.Atoi(naturalScienceScoreStr)
		naturalScienceScore = tempNaturalScience
		if naturalScienceScore > 100 || naturalScienceScore < 0 {
			fmt.Println("Range score should be between 0 and 100")
		} else if errIntNaturalScience != nil {
			fmt.Println("Value data type should be integer ", errIntNaturalScience)
		} else {
			break
		}
	}

	fmt.Printf("\nSubject score\nBahasa: %d (%s)\nEnglish: %d (%s)\nMath: %d (%s)\nNatural Science: %d (%s)\n\n", bahasaScore, reflect.TypeOf(bahasaScore), englishScore, reflect.TypeOf(englishScore), mathScore, reflect.TypeOf(mathScore), naturalScienceScore, reflect.TypeOf(naturalScienceScore))

	meanScore := float64((bahasaScore + englishScore + mathScore + naturalScienceScore) / 4)
	fmt.Printf("Mean score: %.2f\n\n", meanScore)

	if meanScore <= 100 && meanScore >= 90 {
		fmt.Println("Grade: A")
	} else if meanScore < 90 && meanScore >= 80 {
		fmt.Println("Grade: B")
	} else if meanScore < 80 && meanScore >= 70 {
		fmt.Println("Grade: C")
	} else if meanScore < 70 && meanScore >= 60 {
		fmt.Println("Grade: D")
	} else if meanScore < 60 && meanScore >= 0 {
		fmt.Println("Grade: E")
	} else {
		fmt.Println("Undefined")
	}
	fmt.Println()

	// 4.	Buatlah program yang memiliki satu variabel dengan nama “printSegitiga” yg berisi tipe data number yang menghasilkan output segitiga terbalik
	// 		yang berisi angka
	fmt.Println("======================= 4. Print Reverse Right Triangle =======================")
	var rows int
	var rowsStr string
	for {
		fmt.Print("Input rows: ")
		fmt.Scan(&rowsStr)
		tempRows, errIntRows := strconv.Atoi(rowsStr)
		rows = tempRows
		if errIntRows != nil {
			fmt.Println("Value data type should be integer")
		} else if rows <= 0 {
			fmt.Println("Rows should be more than 0")
		} else {
			break
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < rows-i; j++ {
			fmt.Printf("%d ", j+1)
		}
		fmt.Println()
	}
}
