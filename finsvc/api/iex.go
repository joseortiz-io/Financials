package api

import (
	iex "github.com/goinvest/iexcloud/v2"
)


	client := iex.NewClient(os.Getenv("IEX_KEY"), iex.WithBaseURL("https://sandbox.iexapis.com/stable"))

	bs, err := client.AnnualBalanceSheets(context.Background(), "aapl", 1)
	if err != nil {
		log.Fatalf("Error getting balance sheets: %s", err)
	}

	
	statement := bs.Statements[0]

	// print each member of statement
	fmt.Println("ReportDate:", statement.ReportDate)
	fmt.Println("FilingType:", statement.FilingType)
	fmt.Println("FiscalDate:", statement.FiscalDate)
	fmt.Println("FiscalQuarter:", statement.FiscalQuarter)
	fmt.Println("FiscalYear:", statement.FiscalYear)
	fmt.Println("Currency:", statement.Currency)
	fmt.Println("CurrentCash:", statement.CurrentCash)
	fmt.Println("ShortTermInvestments:", statement.ShortTermInvestments)
	fmt.Println("Receivables:", statement.Receivables)
	fmt.Println("Inventory:", statement.Inventory)
	fmt.Println("OtherCurrentAssets:", statement.OtherCurrentAssets)
	fmt.Println("CurrentAssets:", statement.CurrentAssets)
	fmt.Println("LongTermInvestments:", statement.LongTermInvestments)
	fmt.Println("PropertyPlantEquipment:", statement.PropertyPlantEquipment)
	fmt.Println("Goodwill:", statement.Goodwill)
	fmt.Println("IntangibleAssets:", statement.IntangibleAssets)
	fmt.Println("OtherAssets:", statement.OtherAssets)
	fmt.Println("TotalAssets:", statement.TotalAssets)
	fmt.Println("AccountsPayable:", statement.AccountsPayable)
	fmt.Println("CurrentLongTermDebt:", statement.CurrentLongTermDebt)
	fmt.Println("OtherCurrentLiabilities:", statement.OtherCurrentLiabilities)
	fmt.Println("TotalCurrentLiabilities:", statement.TotalCurrentLiabilities)
	fmt.Println("LongTermDebt:", statement.LongTermDebt)
	fmt.Println("OtherLiablities:", statement.OtherLiablities)
	fmt.Println("MinorityInterest:", statement.MinorityInterest)
	fmt.Println("TotalLiabilities:", statement.TotalLiabilities)
	fmt.Println("CommonStock:", statement.CommonStock)
	fmt.Println("RetainedEarnings:", statement.RetainedEarnings)
	fmt.Println("TreasuryStock:", statement.TreasuryStock)
	fmt.Println("CapitalSurplus:", statement.CapitalSurplus)
	fmt.Println("ShareholderEquity:", statement.ShareholderEquity)
	fmt.Println("NetTangibleAssets:", statement.NetTangibleAssets)
	

	fmt.Println("Number of Balance sheets: ", len(bs.Statements))
	fmt.Println("Balance sheets: ", bs.Statements[0].Currency)
	fmt.Println("Balance sheets Report Date: ", bs.Statements[0].ReportDate)

	q, err := client.Quote(context.Background(), "aapl")
	if err != nil {
		log.Fatalf("Error getting quote: %s", err)
	}
	fmt.Println("Company Name: ", q.CompanyName)
	fmt.Println("Quote: ", q.Close)
	aapl := pb.Ticker{Symbol: "aapl"}
	fmt.Println("Ticker: ", aapl.GetSymbol())