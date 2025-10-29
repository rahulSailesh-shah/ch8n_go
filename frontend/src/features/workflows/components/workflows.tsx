import { generateSlug } from "random-word-slugs";
import {
  EmptyView,
  EntityContainer,
  EntityHeader,
  EntityItem,
  EntityList,
  EntityPagination,
  EntitySearch,
  ErrorView,
  LoadingView,
} from "@/components/entity-component";
import { useUpgradeModal } from "@/hooks/use-upgrade-modal";
import { useNavigate, useRouter, useSearch } from "@tanstack/react-router";
import type { ReactNode } from "react";
import {
  useCreateWorkflow,
  useDeleteWorkflow,
  useQueryWorkflows,
} from "../hooks/use-workflows";
import { useEntitySearch } from "@/hooks/use-entity-search";
import { Route } from "@/routes/_authenticated/workflows";
import { WorkflowIcon } from "lucide-react";
import { formatDistanceToNow } from "date-fns";
import type { PaginatedWorkflowResponse, WorkflowDTO } from "../types";

export const WorkflowsList = ({
  data,
}: {
  data: PaginatedWorkflowResponse;
}) => {
  return (
    <EntityList
      items={data.workflows || []}
      getKey={(workflow) => workflow.id}
      emptyView={<WorkflowsEmptyView />}
      renderItem={(workflow) => (
        <WorkflowItem key={workflow.id} data={workflow} />
      )}
    />
  );
};

export const WorkflowItem = ({ data }: { data: WorkflowDTO }) => {
  const deleteWorkflow = useDeleteWorkflow();
  return (
    <EntityItem
      href={`/workflows/${data.id}`}
      title={data.name}
      subtitle={
        <>
          Created{" "}
          {formatDistanceToNow(new Date(data.createdAt), { addSuffix: true })}
          {"  "} &bull; Updated{" "}
          {formatDistanceToNow(new Date(data.updatedAt), { addSuffix: true })}
        </>
      }
      image={
        <div className="size-8 flex items-center justify-center">
          <WorkflowIcon className="size-5 text-muted-foreground" />
        </div>
      }
      key={data.id}
      onRemove={() => {
        deleteWorkflow.mutate(String(data.id));
      }}
      isRemoving={deleteWorkflow.isPending}
    />
  );
};

export const WorkflowsHeader = ({ disabled }: { disabled?: boolean }) => {
  const createWorkflow = useCreateWorkflow();
  const { handleError, modal } = useUpgradeModal();
  const router = useRouter();

  const handleCreateWorkflow = () => {
    createWorkflow.mutate(
      {
        name: generateSlug(),
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

export const WorkflowsLoadingView = () => {
  return <LoadingView message="Loading workflows..." />;
};

export const WorkflowsErrorView = () => {
  return <ErrorView message="An error occurred while loading workflows" />;
};

export const WorkflowsEmptyView = () => {
  const createWorkflow = useCreateWorkflow();
  const { handleError, modal } = useUpgradeModal();
  const router = useRouter();

  const handleCreateWorkflow = () => {
    createWorkflow.mutate(
      {
        name: generateSlug(),
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
      <EmptyView
        message="You haven't created any workflows yet. Get started by creating a new workflow."
        onNew={handleCreateWorkflow}
      />
    </>
  );
};
