import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/executions/$executionId")({
  component: RouteComponent,
});

function RouteComponent() {
  const { executionId } = Route.useParams();
  return <div>Execution {executionId}</div>;
}
