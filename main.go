package main

import (
	"challenge-go/repository"
	"challenge-go/services"
	"fmt"
	"os"
)

func main(){
	args := os.Args[1]
	
	repo := repository.NewRepository(args)
	serv := services.NewServices(repo)

	if donates, err := serv.Sortdata(); err == nil {
		if donateInfo, err := serv.CalculateDonate(donates); err == nil {
			fmt.Printf("Performing donations...\n")
			fmt.Printf("done.\n\n")
			fmt.Printf("\t       Total Recieve: THB  %d\n",donateInfo.TotalSum)
			fmt.Printf("\tSuccessfully donated: THB  %d\n",donateInfo.ValidSum)
			fmt.Printf("\t     Faulty donation: THB   %d\n\n",donateInfo.InvalidSum)
			fmt.Printf("\t  Average per person: THB     %d\n",(donateInfo.ValidSum/donateInfo.ValidCard))
			fmt.Printf("\t          Top donest: THB     %d\n",donateInfo.TopDonate)
			fmt.Printf("\t          Top donors: %s",donateInfo.TopDonor)
		}
	}
}