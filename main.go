package main

import (
	"challenge-go/repository"
	"challenge-go/services"
	"fmt"
	"os"
)

type donate struct{
	services.DonateInfo
}

func main(){
	args := os.Args[1]
	
	repo := repository.NewRepository(args)
	serv := services.NewServices(repo)

	donates, err := serv.Sortdata()
	if err != nil {
		fmt.Println(err)
	}

	type Donate_info = services.DonateInfo
	
	var donate_info Donate_info
	err = serv.CalculateDonate(donates, &donate_info)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Performing donations...\n")
	fmt.Printf("done.\n\n")
	fmt.Printf("\t       Total Recieve: THB  %d\n",donate_info.Total_sum)
	fmt.Printf("\tSuccessfully donated: THB  %d\n",donate_info.Valid_sum)
	fmt.Printf("\t     Faulty donation: THB   %d\n\n",donate_info.Invalid_sum)
	fmt.Printf("\t  Average per person: THB     %d\n",(donate_info.Valid_sum/donate_info.Valid_card))
	fmt.Printf("\t          Top donest: THB     %d\n",donate_info.Top_donate)
	fmt.Printf("\t          Top donors: %s",donate_info.Top_donor)
}