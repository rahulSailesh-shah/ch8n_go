import { createFileRoute, Link } from "@tanstack/react-router";
import LoginForm from "../features/auth/components/LoginForm";
import { requireNoAuth } from "@/lib/auth-utils";

export const Route = createFileRoute("/login")({
  beforeLoad: requireNoAuth,
  component: RouteComponent,
});

function RouteComponent() {
  return (
    <div className="bg-muted flex min-h-svh flex-col items-center justify-center">
      <div className="w-full max-w-sm flex flex-col gap-6">
        <Link to="/" className="flex items-center self-center font-medium">
          <img src="/logos/logo.svg" alt="Logo" className="w-24 h-24" />
        </Link>
        <LoginForm />
      </div>
    </div>
  );
}
