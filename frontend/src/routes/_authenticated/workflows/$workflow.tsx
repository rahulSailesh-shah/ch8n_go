import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/_authenticated/workflows/$workflow")({
  component: RouteComponent,
});

function RouteComponent() {
  return <div>Hello "/workflows/$workflow"!</div>;
}
