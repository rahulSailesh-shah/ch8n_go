import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";

interface ManualTriggerDialogProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
}

export const ManualTriggerDialog = ({
  open,
  onOpenChange,
}: ManualTriggerDialogProps) => {
  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Manual Trigger</DialogTitle>
          <DialogDescription>
            Configure settings for Manual Trigger
          </DialogDescription>
        </DialogHeader>
        <div className="py-4">
          <p className="text-sm text-muted-foreground">
            No configuration required, used to manually trigger the workflow.
          </p>
        </div>
      </DialogContent>
    </Dialog>
  );
};
