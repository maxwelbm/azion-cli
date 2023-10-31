package edge_applications

import (
	"context"
	"fmt"
	"strings"

	"github.com/aziontech/azion-cli/pkg/logger"

	"github.com/fatih/color"

	"github.com/MakeNowJust/heredoc"
	table "github.com/MaxwelMazur/tablecli"
	"github.com/aziontech/azion-cli/messages/general"
	msg "github.com/aziontech/azion-cli/messages/list/edge_applications"
	api "github.com/aziontech/azion-cli/pkg/api/edge_applications"
	"github.com/aziontech/azion-cli/pkg/cmdutil"
	"github.com/aziontech/azion-cli/pkg/contracts"
	"github.com/aziontech/azion-cli/utils"
	"github.com/spf13/cobra"
)

func NewCmd(f *cmdutil.Factory) *cobra.Command {
	opts := &contracts.ListOptions{}

	cmd := &cobra.Command{
		Use:           msg.Usage,
		Short:         msg.ShortDescription,
		Long:          msg.LongDescription,
		SilenceUsage:  true,
		SilenceErrors: true, Example: heredoc.Doc(`
		$ azion list edge-application
		$ azion list edge-application --details
		$ azion list edge-application --page 1 
		$ azion list edge-application --page-size 5
		`),
		RunE: func(cmd *cobra.Command, _ []string) error {
			client := api.NewClient(f.HttpClient, f.Config.GetString("api_url"), f.Config.GetString("token"))

			if err := PrintTable(cmd, client, f, opts); err != nil {
				return fmt.Errorf(msg.ErrorGetAll.Error(), err)
			}
			return nil
		},
	}

	flags := cmd.Flags()
	flags.Int64Var(&opts.Page, "page", 1, general.ApiListFlagPage)
	flags.Int64Var(&opts.PageSize, "page-size", 10, general.ApiListFlagPageSize)
	flags.BoolVar(&opts.Details, "details", false, general.ApiListFlagDetails)
	flags.BoolP("help", "h", false, msg.HelpFlag)
	return cmd
}

func PrintTable(cmd *cobra.Command, client *api.Client, f *cmdutil.Factory, opts *contracts.ListOptions) error {
	c := context.Background()

	for {
		resp, err := client.List(c, opts)
		if err != nil {
			return err
		}

		tbl := table.New("ID", "NAME", "ACTIVE")
		tbl.WithWriter(f.IOStreams.Out)

		if opts.Details {
			tbl = table.New("ID", "NAME", "ACTIVE", "LAST EDITOR", "LAST MODIFIED", "DEBUG RULES")
		}

		headerFmt := color.New(color.FgBlue, color.Underline).SprintfFunc()
		columnFmt := color.New(color.FgGreen).SprintfFunc()
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

		for _, v := range resp.Results {
			tbl.AddRow(
				v.Id,
				utils.TruncateString(v.Name),
				v.Active,
				v.LastEditor,
				v.LastModified,
				v.DebugRules,
			)
		}

		format := strings.Repeat("%s", len(tbl.GetHeader())) + "\n"
		tbl.CalculateWidths([]string{})

		// print the header only in the first flow
		if opts.Page == 1 {
			logger.PrintHeader(tbl, format)
		}

		for _, row := range tbl.GetRows() {
			logger.PrintRow(tbl, format, row)
		}

		if opts.Page >= resp.TotalPages {
			break
		}

		if cmd.Flags().Changed("page") || cmd.Flags().Changed("page-size") {
			break
		}

		opts.Page++
	}

	return nil
}
