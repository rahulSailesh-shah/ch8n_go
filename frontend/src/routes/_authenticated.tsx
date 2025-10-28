import AppSideBar from "@/components/app-sidebar";
import { SidebarInset, SidebarProvider } from "@/components/ui/sidebar";
import { requireAuth } from "@/lib/auth-utils";
import { createFileRoute, Outlet } from "@tanstack/react-router";

export const Route = createFileRoute("/_authenticated")({
  beforeLoad: requireAuth,
  component: () => (
    <SidebarProvider>
      <AppSideBar />
      <SidebarInset>
        <Outlet />
      </SidebarInset>
    </SidebarProvider>
  ),
});
