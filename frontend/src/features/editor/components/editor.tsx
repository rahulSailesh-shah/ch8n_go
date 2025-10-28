import { useQueryWorkflow } from "@/features/workflows/hooks/use-workflows";
import { LoadingView } from "@/components/entity-component";
import { ErrorView } from "@/components/entity-component";

export const Editor = ({ workflowId }: { workflowId: string }) => {
  const {
    data: workflowData,
    isLoading,
    isError,
  } = useQueryWorkflow(workflowId);

  if (isLoading) {
    return <WorkflowLoadingView />;
  }

  if (isError) {
    return <WorkflowErrorView />;
  }

  return <div>{JSON.stringify(workflowData, null, 2)}</div>;
};

export const WorkflowLoadingView = () => {
  return <LoadingView message="Loading editor..." />;
};

export const WorkflowErrorView = () => {
  return <ErrorView message="An error occurred while loading the workflow" />;
};
