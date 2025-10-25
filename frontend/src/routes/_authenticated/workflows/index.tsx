import { Button } from "@/components/ui/button";
import { useCreateWorkflow } from "@/services/workflow/hooks";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/_authenticated/workflows/")({
  component: RouteComponent,
});

function RouteComponent() {
  const createWorkflowMutation = useCreateWorkflow();

  return (
    <main className="flex-1">
      <div className="bg-accent/20">Workflows</div>
      <Button
        onClick={() =>
          createWorkflowMutation.mutate({
            name: "New Workflow",
            description: "New Workflow",
          })
        }
      >
        Create Workflow
      </Button>
    </main>
  );
}
