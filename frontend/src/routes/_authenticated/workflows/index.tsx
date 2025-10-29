import AppHeader from "@/components/app-header";
import { QueryBoundary } from "@/components/query-boundary";
import {
  WorkflowErrorView,
  WorkflowLoadingView,
} from "@/features/editor/components/editor";
import {
  WorkflowsContainer,
  WorkflowsEmptyView,
  WorkflowsList,
} from "@/features/workflows/components/workflows";
import { useQueryWorkflows } from "@/features/workflows/hooks/use-workflows";
import { createFileRoute } from "@tanstack/react-router";
import { z } from "zod";

const workflowSearchSchema = z.object({
  page: z.number().int().positive().catch(1),
  limit: z.number().int().positive().catch(5),
  search: z.string().catch(""),
});

export type WorkflowSearchParams = z.infer<typeof workflowSearchSchema>;

export const Route = createFileRoute("/_authenticated/workflows/")({
  validateSearch: workflowSearchSchema,
  component: RouteComponent,
});

function RouteComponent() {
  const workflowsQuery = useQueryWorkflows();

  return (
    <>
      <QueryBoundary
        query={workflowsQuery}
        loadingFallback={<WorkflowLoadingView />}
        errorFallback={<WorkflowErrorView />}
        emptyFallback={<WorkflowsEmptyView />}
      >
        {(workflow) => (
          <>
            <AppHeader />
            <main className="flex-1">
              <WorkflowsContainer>
                <WorkflowsList data={workflow} />
              </WorkflowsContainer>
            </main>
          </>
        )}
      </QueryBoundary>
    </>
  );
}
