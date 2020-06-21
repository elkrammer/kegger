// helper functions
package helper

import (
	"bytes"
	"fmt"
	"strconv"
	"text/template"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func FindDifference(a, b []int) (diff []int) {
	m := make(map[int]bool)

	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return
}

func ProcessTemplateFile(fileName string, vars interface{}) string {
	tmpl, err := template.ParseFiles(fileName)

	if err != nil {
		panic(err)
	}
	return ParseTemplate(tmpl, vars)
}

func ParseTemplate(t *template.Template, vars interface{}) string {
	var tmplBytes bytes.Buffer

	err := t.Execute(&tmplBytes, vars)
	if err != nil {
		panic(err)
	}
	return tmplBytes.String()
}

func FormatSpanishDate(t time.Time) string {
	// spanish translations

	var esDays = [...]string{
		"Domingo", "Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado",
	}
	var esMonths = [...]string{
		"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio",
		"Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre",
	}

	date := fmt.Sprintf("%s %02d de %s %s a las %02d:%02d",
		esDays[t.Weekday()],
		t.Day(),
		esMonths[t.Month()-1],
		strconv.Itoa(t.Year()),
		t.Hour(), t.Minute())

	return date
}
