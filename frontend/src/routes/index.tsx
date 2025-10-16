import { createFileRoute } from "@tanstack/react-router";
import { requireAuth } from "@/lib/auth-utils";
import { LogoutButton } from "@/features/auth/components/LogoutButton";

export const Route = createFileRoute("/")({
  beforeLoad: requireAuth,
  component: RouteComponent,
});

function RouteComponent() {
  return (
    <div>
      <LogoutButton />
    </div>
  );
}
