import {
  WorkflowsContainer,
  WorkflowsList,
} from "@/features/workflows/components/workflows";
import { createFileRoute } from "@tanstack/react-router";
import { z } from "zod";

const workflowSearchSchema = z.object({
  page: z.number().int().positive().catch(1),
  limit: z.number().int().positive().catch(1),
  search: z.string().catch(""),
});

export type WorkflowSearchParams = z.infer<typeof workflowSearchSchema>;

export const Route = createFileRoute("/_authenticated/workflows/")({
  validateSearch: workflowSearchSchema,
  component: RouteComponent,
});

function RouteComponent() {
  return (
    <WorkflowsContainer>
      <main className="flex-1">
        <WorkflowsList />
      </main>
    </WorkflowsContainer>
  );
}
