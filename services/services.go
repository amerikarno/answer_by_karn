package services

import (
	"challenge-go/repository"
	"strconv"
	"strings"
	"time"
)

type DonateInfo struct{
	Invalid_sum int
	Valid_sum int
	Total_sum int
	Invalid_card int
	Valid_card int
	Total_card int
	Top_donate int
	Top_donor string
}

type Donate struct {
	amount int
	ccnumber string
	cvv string
	card Card
}

type Card struct {
	name string
	expiration_month int
	expiration_year	int
}

type Services interface{
	Sortdata() ([]Donate, error)
	CalculateDonate(donates []Donate,donate_info *DonateInfo) error
}

type services struct{repos repository.Repository}

func NewServices(repos repository.Repository) Services{
	return services{repos: repos}
}

func (s services) Sortdata() ([]Donate, error) {
	var donates []Donate
	csv, err := s.repos.Readfile()
	if err != nil {
		return nil, err
	}

	for i, lines := range strings.Split(*csv,"\n") {
		if i > 0 {
			line := make([]string,6,6)
			line = strings.Split(lines,",")
			amount, err := strconv.Atoi(line[1])
			if err != nil {
				return nil, err
			}
			month, err := strconv.Atoi(line[4])
			if err != nil {
				return nil, err
			}
			year, err := strconv.Atoi(line[5])
			if err != nil {
				return nil, err
			}

			card := Card{
				name: line[0],
				expiration_month: month,
				expiration_year: year,
			}

			donate := Donate{
				amount: amount,
				ccnumber: line[2],
				cvv: line[3],
				card: card,
			}
			
			donates = append(donates, donate)
		}
	}
	return donates, nil
}

func (s services) CalculateDonate(donates []Donate,donate_info *DonateInfo) error{

	thisyear := time.Now().Year()
	thismonth := time.Now().Month()
	
	// var donate_info DonateInfo

	for k, d := range donates {
		if d.amount > donate_info.Top_donate {
			donate_info.Top_donate = d.amount
			donate_info.Top_donor = d.card.name
		}
		donate_info.Total_card = k+1


		if d.card.expiration_month < int(thismonth) && d.card.expiration_year < thisyear {
			donate_info.Invalid_sum += d.amount
			donate_info.Invalid_card++
		} else {
			donate_info.Valid_sum += d.amount
			donate_info.Valid_card++
		}
		donate_info.Total_sum += d.amount
	}
	return nil
}