package main

import (
	context "context"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "fredericschmidt.fr/grpcdemos/GoDemo01/calculatrice"
)

const (
	uri = "localhost:666"
)

func main() {

	// Action
	var action string
	flag.StringVar(&action, "action", "add", "This is the calculator operations (add, sub, table)")

	// Term A
	termA := flag.Int("termA", 0, "This is the TermA for ADD or SUB actions.")
	// Term B
	termB := flag.Int("termB", 0, "This is the TermB for ADD or SUB actions.")
	// Multiplier
	multiplier := flag.Int("Multiplier", 0, "This is the Multiplier for TABLE actions.")
	// Multiplicand
	multiplicand := flag.Int("Multiplicand", 0, "This is the Multiplicand for TABLE actions.")

	flag.Parse()

	/* ***** */
	fmt.Println("Client started")

	cnx, err := grpc.Dial(uri, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("connection impossible: %v", err)
	}

	defer cnx.Close()
	c := pb.NewCalculatorServiceClient(cnx)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	/* Action */
	switch action {
	case "ADD":
		pAdd := &pb.ValuesCalculatorRequest{
			TermX: (int32)(*termA),
			TermY: (int32)(*termB),
		}
		rs, err := c.AddOperation(ctx, pAdd)
		if err != nil {
			log.Fatalf("could not do the addition: %v", err)
		}
		fmt.Printf("%d + %d = %d\n", pAdd.TermX, pAdd.TermY, rs.GetResult())

	case "SUB":
		pSub := &pb.ValuesCalculatorRequest{
			TermX: (int32)(*termA),
			TermY: (int32)(*termB),
		}
		rs, err := c.SubOperation(ctx, pSub)
		if err != nil {
			log.Fatalf("could not do the addition: %v", err)
		}
		fmt.Printf("%d - %d = %d\n", pSub.TermX, pSub.TermY, rs.GetResult())

	case "TABLE":
		pTable := &pb.TableCalculatorRequest{
			Multiplicand: (int32)(*multiplicand),
			Multiplier:   (int32)(*multiplier),
		}
		rs, err := c.TableOperation(ctx, pTable)
		if err != nil {
			log.Fatalf("could not do the addition: %v", err)
		}
		fmt.Printf("Table de multiplation de %d X %d.\n", pTable.Multiplicand, pTable.Multiplier)
		for _, oneline := range rs.LineOfTable {
			fmt.Printf("%d X %d = %d\n", oneline.GetMultiplicand(), oneline.GetMultiplier(), oneline.GetProduct())
		}

	default:
		fmt.Printf("The actions %s is unknown.\n", action)
	}

	fmt.Println("Client ended.")
}
