import AppHeader from "@/components/app-header";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/_authenticated/credentials/")({
  component: RouteComponent,
});

function RouteComponent() {
  return (
    <main className="flex-1">
      <AppHeader />
      <div className="bg-accent/20">Credentials</div>
    </main>
  );
}
