import { memo, useState } from "react";
import { type Node, type NodeProps } from "@xyflow/react";
import { WebhookIcon } from "lucide-react";
import { BaseTriggerNode } from "../base-trigger-node";
import { WebhookTriggerDialog } from "./dialog";

type WebhookTriggerNodeType = Node;

export const WebhookTriggerNode = memo(
  (props: NodeProps<WebhookTriggerNodeType>) => {
    const [isDialogOpen, setIsDialogOpen] = useState(false);

    const handleOpenSettings = () => {
      setIsDialogOpen(true);
    };

    const nodeStatus = "initial";

    return (
      <>
        <WebhookTriggerDialog
          open={isDialogOpen}
          onOpenChange={setIsDialogOpen}
        />
        <BaseTriggerNode
          {...props}
          icon={WebhookIcon}
          name="Webhook Trigger"
          onSettingsClick={handleOpenSettings}
          onDoubleClick={handleOpenSettings}
          status={nodeStatus}
        />
      </>
    );
  }
);

WebhookTriggerNode.displayName = "WebhookTriggerNode";
