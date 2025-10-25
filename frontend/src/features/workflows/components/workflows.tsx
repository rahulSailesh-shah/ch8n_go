import { EntityContainer, EntityHeader } from "@/components/entity-component";
import { useUpgradeModal } from "@/hooks/use-upgrade-modal";
import { useCreateWorkflow, useWorkflows } from "@/services/workflow/hooks";
import { useRouter } from "@tanstack/react-router";
import type { ReactNode } from "react";

export const WorkflowsList = () => {
  const { data: workflows, isLoading, isError, error } = useWorkflows();

  if (isLoading) {
    return <div className="bg-accent/20">Loading...</div>;
  }

  if (isError) {
    return (
      <div>
        <p className="text-sm font-normal text-muted-foreground">
          {error?.message}
        </p>
      </div>
    );
  }

  if (!workflows?.workflows?.length) {
    return <h1>No workflows found</h1>;
  }

  return (
    <div className="flex flex-col gap-y-4">
      {workflows.workflows.map((workflow) => (
        <div key={workflow.id} className="bg-accent/20">
          {workflow.name}
          <p>{workflow.userId}</p>
        </div>
      ))}
    </div>
  );
};

export const WorkflowsHeader = ({ disabled }: { disabled?: boolean }) => {
  const createWorkflow = useCreateWorkflow();
  const { handleError, modal } = useUpgradeModal();
  const router = useRouter();

  const handleCreateWorkflow = () => {
    createWorkflow.mutate(
      {
        name: "komromro",
        description: "mkowmow",
      },
      {
        onError: (error) => {
          handleError(error);
        },
        onSuccess: (workflow) => {
          router.navigate({
            to: `/workflows/${workflow?.id}`,
          });
        },
      }
    );
  };
  return (
    <>
      {modal}
      <EntityHeader
        title="Workflows"
        description="Create and manage your workflows"
        onNew={handleCreateWorkflow}
        newButtonLabel="New Workflow"
        disabled={disabled}
        isCreating={createWorkflow.isPending}
      />
    </>
  );
};

export const WorkflowsContainer = ({ children }: { children: ReactNode }) => {
  return (
    <EntityContainer
      header={<WorkflowsHeader />}
      search={<></>}
      pagination={<></>}
    >
      {children}
    </EntityContainer>
  );
};
