import AppHeader from "@/components/app-header";
import AppSideBar from "@/components/app-sidebar";
import { SidebarInset, SidebarProvider } from "@/components/ui/sidebar";
import { requireAuth } from "@/lib/auth-utils";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/credentials/")({
  beforeLoad: requireAuth,
  component: RouteComponent,
});

function RouteComponent() {
  return (
    <SidebarProvider>
      <AppSideBar />
      <SidebarInset>
        <AppHeader />
        <main className="flex-1">
          <div className="bg-accent/20">Credentials</div>
        </main>
      </SidebarInset>
    </SidebarProvider>
  );
}
