import { createFileRoute } from "@tanstack/react-router";
import { RegisterForm } from "../features/auth/components/RegisterForm";
import { requireNoAuth } from "@/lib/auth-utils";

export const Route = createFileRoute("/signup")({
  beforeLoad: requireNoAuth,
  component: RouteComponent,
});

function RouteComponent() {
  return <RegisterForm />;
}
