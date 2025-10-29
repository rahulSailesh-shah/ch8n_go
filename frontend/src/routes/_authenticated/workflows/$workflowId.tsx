import { QueryBoundary } from "@/components/query-boundary";
import {
  Editor,
  WorkflowErrorView,
  WorkflowLoadingView,
} from "@/features/editor/components/editor";
import EditorHeader from "@/features/editor/components/editor-header";
import { useQueryWorkflow } from "@/features/workflows/hooks/use-workflows";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/_authenticated/workflows/$workflowId")({
  component: RouteComponent,
});

function RouteComponent() {
  const { workflowId } = Route.useParams();
  const workflowQuery = useQueryWorkflow(workflowId);
  return (
    <>
      <QueryBoundary
        query={workflowQuery}
        loadingFallback={<WorkflowLoadingView />}
        errorFallback={<WorkflowErrorView />}
      >
        {(workflow) => (
          <>
            <EditorHeader workflow={workflow} />
            <main className="flex-1">
              <Editor workflow={workflow} />
            </main>
          </>
        )}
      </QueryBoundary>
    </>
  );
}
