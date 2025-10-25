import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/_authenticated/credentials/$credentialId")({
  component: RouteComponent,
});

function RouteComponent() {
  const { credentialId } = Route.useParams();

  return <div>Credential {credentialId}</div>;
}
