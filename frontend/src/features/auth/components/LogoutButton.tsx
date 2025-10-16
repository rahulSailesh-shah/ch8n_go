import { Button } from "@/components/ui/button";
import { authClient } from "@/lib/auth-client";
import { useRouter } from "@tanstack/react-router";
import { toast } from "sonner";
import { useState } from "react";

export function LogoutButton() {
  const router = useRouter();
  const [isLoading, setIsLoading] = useState(false);

  const onSubmit = async () => {
    setIsLoading(true);
    try {
      const { error } = await authClient.signOut(
        {},
        {
          onSuccess: () => {
            router.navigate({ to: "/login" });
          },
          onError: (ctx) => {
            toast("Error: " + ctx.error.message);
          },
        }
      );
      if (error) {
        toast("Error: " + error.message);
      }
    } catch (error) {
      toast("Something went wrong");
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <Button onClick={onSubmit} disabled={isLoading}>
      {isLoading ? "Signing out..." : "Sign out"}
    </Button>
  );
}
