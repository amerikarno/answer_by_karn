package services

import (
	"challenge-go/repository"
	"errors"
	"strconv"
	"strings"
	"time"
)

// A DonateInfo ... allow main.go use DonateInfo struct
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

type donate struct {
	amount int
	ccnumber string
	cvv string
	card card
}

type card struct {
	name string
	expirationMonth int
	expirationYear	int
}

// A Services interface ... aPlug
type Services interface{
	Sortdata() (donates []donate,err error)
	CalculateDonate(donates []donate) (DonateInfo DonateInfo, err error)
}

type services struct{repos repository.Repository}

// NewServices ... function: an Adapter
func NewServices(repos repository.Repository) Services{
	return services{repos}
}

func (s services) Sortdata() (donates []donate,err error) {
	csv, err := s.repos.Readfile()
	if err != nil {
		return nil, errors.New("cannot read file")
	}
	
	for i, lines := range strings.Split(*csv,"\n") {
		var donat donate
		if i > 0 {
			line := make([]string,6,6)
			line = strings.Split(lines,",")
			if amount, err := strconv.Atoi(line[1]); err == nil {
				if month, err := strconv.Atoi(line[4]); err == nil {
					if year, err := strconv.Atoi(line[5]); err == nil {

						card := card{
							name: line[0],
							expirationMonth: month,
							expirationYear: year,
						}

						donat = donate{
							amount: amount,
							ccnumber: line[2],
							cvv: line[3],
							card: card,
						}
					}
				}
			}
			
			donates = append(donates, donat)
		}
	}
	return donates, nil
}

func (s services) CalculateDonate(donates []donate) (DonateInfo DonateInfo, err error){

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