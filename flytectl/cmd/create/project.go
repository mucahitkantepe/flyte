package create

import (
	"context"
	"fmt"

	"github.com/flyteorg/flyte/flytectl/clierrors"
	"github.com/flyteorg/flyte/flytectl/cmd/config"
	"github.com/flyteorg/flyte/flytectl/cmd/config/subcommand/project"
	cmdCore "github.com/flyteorg/flyte/flytectl/cmd/core"
	"github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/admin"
	"github.com/flyteorg/flyte/flytestdlib/logger"
)

const (
	projectShort = "Creates project resources."
	projectLong  = `
Create a project given its name and id.

::

 flytectl create project --name flytesnacks --id flytesnacks --description "flytesnacks description" --labels app=flyte

.. note::
   The terms project/projects are interchangeable in these commands.

Create a project by definition file.

::

 flytectl create project --file project.yaml

.. code-block:: yaml

    id: "project-unique-id"
    name: "Name"
    labels:
       values:
         app: flyte
    description: "Some description for the project."

.. note::
    The project name shouldn't contain any whitespace characters.
`
)

func createProjectsCommand(ctx context.Context, args []string, cmdCtx cmdCore.CommandContext) error {
	projectSpec, err := project.DefaultProjectConfig.GetProjectSpec(config.GetConfig())
	if err != nil {
		return err
	}
	if projectSpec.GetId() == "" {
		return fmt.Errorf(clierrors.ErrProjectNotPassed) //nolint
	}
	if projectSpec.GetName() == "" {
		return fmt.Errorf(clierrors.ErrProjectNameNotPassed) //nolint
	}

	if project.DefaultProjectConfig.DryRun {
		logger.Debugf(ctx, "skipping RegisterProject request (DryRun)")
	} else {
		_, err := cmdCtx.AdminClient().RegisterProject(ctx, &admin.ProjectRegisterRequest{
			Project: &admin.Project{
				Id:          projectSpec.GetId(),
				Name:        projectSpec.GetName(),
				Description: projectSpec.GetDescription(),
				Labels:      projectSpec.GetLabels(),
			},
		})
		if err != nil {
			return err
		}
	}
	fmt.Println("project created successfully.")
	return nil
}
