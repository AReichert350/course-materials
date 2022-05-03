package miner

import (
	"fmt"
	"regexp"
	"strings"
)

type DatabaseMiner interface {
	GetSchema() (*Schema, error)
}

type Schema struct {
	Databases []Database
}

type Database struct {
	Name   string
	Tables []Table
}

type Table struct {
	Name    string
	Columns []string
}

func Search(m DatabaseMiner) ([]string, error) {
	s, err := m.GetSchema()
	if err != nil {
		return nil, err
	}

	var (
		mineResults []string
	)

	re := getRegex()
	for _, database := range s.Databases {
		for _, table := range database.Tables {
			for _, field := range table.Columns {
				for _, r := range re {
					if r.MatchString(field) {
						for _, line := range strings.Split(database.String(), "\n") {
							mineResults = append(mineResults, line)
						}
						// log.Println(database)
						mineResults = append(mineResults, fmt.Sprintf("[+] HIT: %s\n", field))
						// log.Printf("[+] HIT: %s\n", field)
					}
				}
			}
		}
	}
	return mineResults, nil
}

func getRegex() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`(?i)social`),
		regexp.MustCompile(`(?i)ssn`),
		regexp.MustCompile(`(?i)pass(word)?`),
		regexp.MustCompile(`(?i)hash`),
		regexp.MustCompile(`(?i)ccnum`),
		regexp.MustCompile(`(?i)card`),
		regexp.MustCompile(`(?i)security`),
		regexp.MustCompile(`(?i)key`),
	}
}

func (s Schema) String() string {
	var ret string
	for _, database := range s.Databases {
		ret += fmt.Sprint(database.String() + "\n")
	}
	return ret
}

func (d Database) String() string {
	ret := fmt.Sprintf("[DB] = %+s\n", d.Name)
	for _, table := range d.Tables {
		ret += table.String()
	}
	return ret
}

func (t Table) String() string {
	ret := fmt.Sprintf("    [TABLE] = %+s\n", t.Name)
	for _, field := range t.Columns {
		ret += fmt.Sprintf("       [COL] = %+s\n", field)
	}
	return ret
}
