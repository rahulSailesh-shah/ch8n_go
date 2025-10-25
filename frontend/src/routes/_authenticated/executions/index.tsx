import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/_authenticated/executions/")({
  component: RouteComponent,
});

function RouteComponent() {
  return (
    <main className="flex-1">
      <div className="bg-accent/20">Executions</div>
    </main>
  );
}
