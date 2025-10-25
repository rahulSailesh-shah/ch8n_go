import {
  WorkflowsContainer,
  WorkflowsList,
} from "@/features/workflows/components/workflows";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/_authenticated/workflows/")({
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
