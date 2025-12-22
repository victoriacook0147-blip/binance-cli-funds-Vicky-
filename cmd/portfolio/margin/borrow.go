package margin

import (
	"fmt"
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange/portfolio/margin"
	"github.com/UnipayFI/binance-cli/printer"
	"github.com/spf13/cobra"
)

var (
	borrowCmd = &cobra.Command{
		Use:   "borrow",
		Short: "Borrow",
		Long:  `Query margin max borrow and max withdraw.`,
	}

	marginMaxBorrowCmd = &cobra.Command{
		Use:   "max-borrow",
		Short: "Query margin max borrow",
		Long: `Query margin max borrow.
		
Docs Link: https://developers.binance.com/docs/derivatives/portfolio-margin/account/Margin-Max-Borrow`,
		Run: marginMaxBorrow,
	}

	marginMaxWithdrawCmd = &cobra.Command{
		Use:   "max-withdraw",
		Short: "Query Margin Max Withdraw",
		Long: `Query Margin Max Withdraw.
		
Docs Link: https://developers.binance.com/docs/derivatives/portfolio-margin/account/Query-Margin-Max-Withdraw`,
		Run: marginMaxWithdraw,
	}
)

func InitBorrowCmds() []*cobra.Command {
	marginMaxBorrowCmd.Flags().StringP("asset", "a", "", "asset")
	marginMaxBorrowCmd.MarkFlagRequired("asset")
	marginMaxWithdrawCmd.Flags().StringP("asset", "a", "", "asset")
	marginMaxWithdrawCmd.MarkFlagRequired("asset")
	borrowCmd.AddCommand(marginMaxBorrowCmd, marginMaxWithdrawCmd)
	return []*cobra.Command{borrowCmd}
}

func marginMaxBorrow(cmd *cobra.Command, _ []string) {
	asset, _ := cmd.Flags().GetString("asset")
	client := margin.NewClient(config.Config.APIKey, config.Config.APISecret)
	borrow, err := client.MarginMaxBorrow(asset)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(borrow)
}

func marginMaxWithdraw(cmd *cobra.Command, _ []string) {
	asset, _ := cmd.Flags().GetString("asset")
	client := margin.NewClient(config.Config.APIKey, config.Config.APISecret)
	withdraw, err := client.MarginMaxWithdraw(asset)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("margin max withdraw: %s %s\n", withdraw, asset)
}
