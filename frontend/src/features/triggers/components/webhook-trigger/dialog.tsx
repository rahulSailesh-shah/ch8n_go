import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogDescription,
} from "@/components/ui/dialog";
import { Item, ItemContent, ItemTitle } from "@/components/ui/item";
import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";
import { CopyIcon } from "lucide-react";

interface WebhookTriggerDialogProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
}

// TODO: Add copy to clipboard functionality
// TODO: Fix Styling
export const WebhookTriggerDialog = ({
  open,
  onOpenChange,
}: WebhookTriggerDialogProps) => {
  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent>
        <DialogHeader className="mb-4">
          <DialogTitle>Webhook Trigger</DialogTitle>
          <DialogDescription>
            No configurations required for the webhook trigger
          </DialogDescription>
        </DialogHeader>

        <div className="flex w-full max-w-md gap-6 flex-col">
          <div className="flex w-full max-w-md gap-2 flex-col">
            <Label>Variable Name</Label>
            <Item variant="outline" size="sm">
              <div>
                <ItemContent>
                  <ItemTitle>WEBHOOK_TRIGGER</ItemTitle>
                </ItemContent>
              </div>
            </Item>
          </div>

          <div className="flex w-full max-w-md gap-2 flex-col ">
            <Label>Webhook URL</Label>
            <div className="w-full flex items-center gap-2">
              <Item variant="outline" size="sm" className="w-full">
                <ItemContent>
                  <ItemTitle>
                    https://ferret-eminent-cobra.ngrok-free.app/api/webhook/86d6e51b-beec-4401-b14e-7736e0d0b123
                  </ItemTitle>
                </ItemContent>
              </Item>
              <Button variant="outline" size="icon-xl" aria-label="Submit">
                <CopyIcon />
              </Button>
            </div>
          </div>
        </div>
      </DialogContent>
    </Dialog>
  );
};
