import { UpgradeModal } from "@/components/upgrade-modal";
import { ApiError } from "@/lib/api-client";
import { useState } from "react";

export const useUpgradeModal = () => {
  const [open, setOpen] = useState(false);

  const handleError = (error: Error) => {
    if (error instanceof ApiError && error.status === 403) {
      setOpen(true);
      return true;
    }
    return false;
  };

  const modal = <UpgradeModal open={open} onOpenChange={setOpen} />;

  return {
    handleError,
    modal,
  };
};
