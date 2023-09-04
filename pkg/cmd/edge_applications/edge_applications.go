package edge_applications

import (
	"github.com/MakeNowJust/heredoc"
	msg "github.com/aziontech/azion-cli/messages/edge_applications"
	"github.com/aziontech/azion-cli/pkg/cmd/edge_applications/delete"
	"github.com/aziontech/azion-cli/pkg/cmd/edge_applications/describe"
	"github.com/aziontech/azion-cli/pkg/cmd/edge_applications/list"
	"github.com/aziontech/azion-cli/pkg/cmd/edge_applications/ls"
	"github.com/aziontech/azion-cli/pkg/cmd/edge_applications/update"
	"github.com/aziontech/azion-cli/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmd(f *cmdutil.Factory) *cobra.Command {
	edge_applicationsCmd := &cobra.Command{
		Use:   msg.EdgeApplicationsUsage,
		Short: msg.EdgeApplicationsShortDescription,
		Long:  msg.EdgeApplicationsLongDescription,
		Example: heredoc.Doc(`
		$ azioncli edge_applications --help
        `),
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	edge_applicationsCmd.AddCommand(describe.NewCmd(f))
	edge_applicationsCmd.AddCommand(delete.NewCmd(f))
	edge_applicationsCmd.AddCommand(list.NewCmd(f))
	edge_applicationsCmd.AddCommand(update.NewCmd(f))
	edge_applicationsCmd.AddCommand(ls.NewCmd(f))

	edge_applicationsCmd.Flags().BoolP("help", "h", false, msg.EdgeApplicationsFlagHelp)

	return edge_applicationsCmd
}
