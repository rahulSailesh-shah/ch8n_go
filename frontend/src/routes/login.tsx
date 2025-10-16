import { createFileRoute } from "@tanstack/react-router";
import LoginForm from "../features/auth/components/LoginForm";
import { requireNoAuth } from "@/lib/auth-utils";

export const Route = createFileRoute("/login")({
  beforeLoad: requireNoAuth,
  component: RouteComponent,
});

function RouteComponent() {
  return <LoginForm />;
}
