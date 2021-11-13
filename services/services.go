package services

import (
	"challenge-go/repository"
	"strconv"
	"strings"
	"time"
)

type DonateInfo struct{
	InvalidSum int
	ValidSum int
	TotalSum int
	InvalidCard int
	ValidCard int
	TotalCard int
	TopDonate int
	TopDonor string
}

type Donate struct {
	amount int
	ccnumber string
	cvv string
	card Card
}

type Card struct {
	name string
	expirationMonth int
	expirationYear	int
}

type Services interface{
	Sortdata() ([]Donate, error)
	CalculateDonate(donates []Donate) (DonateInfo DonateInfo, err error)
}

type services struct{repos repository.Repository}

func NewServices(repos repository.Repository) Services{
	return services{repos: repos}
}

func (s services) Sortdata() ([]Donate, error) {
	csv, err := s.repos.Readfile()
	if err != nil {
		return nil, err
	}
	
	var donates []Donate
	for i, lines := range strings.Split(*csv,"\n") {
		var donate Donate
		if i > 0 {
			line := make([]string,6,6)
			line = strings.Split(lines,",")
			if amount, err := strconv.Atoi(line[1]); err == nil {
				if month, err := strconv.Atoi(line[4]); err == nil {
					if year, err := strconv.Atoi(line[5]); err == nil {

						card := Card{
							name: line[0],
							expirationMonth: month,
							expirationYear: year,
						}

						donate = Donate{
							amount: amount,
							ccnumber: line[2],
							cvv: line[3],
							card: card,
						}
					}
				}
			}
			
			donates = append(donates, donate)
		}
	}
	return donates, err
}

func (s services) CalculateDonate(donates []Donate) (DonateInfo DonateInfo, err error){

	thisyear := time.Now().Year()
	thismonth := time.Now().Month()
	
	// var DonateInfo DonateInfo

	for k, d := range donates {
		if d.amount > DonateInfo.TopDonate {
			DonateInfo.TopDonate = d.amount
			DonateInfo.TopDonor = d.card.name
		}
		DonateInfo.TotalCard = k+1


		if d.card.expirationMonth < int(thismonth) && d.card.expirationYear < thisyear {
			DonateInfo.InvalidSum += d.amount
			DonateInfo.InvalidCard++
		} else {
			DonateInfo.ValidSum += d.amount
			DonateInfo.ValidCard++
		}
		DonateInfo.TotalSum += d.amount
	}
	return DonateInfo, nil
}