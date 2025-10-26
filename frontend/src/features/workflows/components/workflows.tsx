import {
  EntityContainer,
  EntityHeader,
  EntityPagination,
  EntitySearch,
} from "@/components/entity-component";
import { useUpgradeModal } from "@/hooks/use-upgrade-modal";
import { useNavigate, useRouter, useSearch } from "@tanstack/react-router";
import type { ReactNode } from "react";
import { useCreateWorkflow, useQueryWorkflows } from "../hooks/use-workflows";
import { useEntitySearch } from "@/hooks/use-entity-search";
import { Route } from "@/routes/_authenticated/workflows";

export const WorkflowsList = () => {
  const { data: workflows, isLoading, isError, error } = useQueryWorkflows();

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
        name: "rahul",
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

export const WorkflowsSearch = () => {
  const search = useSearch({
    from: "/_authenticated/workflows/",
  });
  const navigate = useNavigate({
    from: Route.fullPath,
  });
  const { searchValue, onSearchChange } = useEntitySearch({
    params: search,
    setParams: (params) => navigate({ search: params }),
  });
  return (
    <EntitySearch
      value={searchValue}
      onSearch={onSearchChange}
      placeholder="Search workflows"
    />
  );
};

export const WorkflowsPagination = () => {
  const { data: workflows, isLoading } = useQueryWorkflows();
  const totalPages = workflows?.totalPages || 1;

  const search = useSearch({
    from: "/_authenticated/workflows/",
  });
  const navigate = useNavigate({
    from: Route.fullPath,
  });

  return (
    <EntityPagination
      page={search.page}
      totalPages={totalPages}
      onPageChange={(page: number) => {
        console.log(page);
        navigate({ search: { ...search, page } });
      }}
      disabled={isLoading}
    />
  );
};

export const WorkflowsContainer = ({ children }: { children: ReactNode }) => {
  return (
    <EntityContainer
      header={<WorkflowsHeader />}
      search={<WorkflowsSearch />}
      pagination={<WorkflowsPagination />}
    >
      {children}
    </EntityContainer>
  );
};
