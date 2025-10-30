import { memo, useState } from "react";
import type { NodeProps } from "@xyflow/react";
import { MousePointer } from "lucide-react";
import { BaseTriggerNode } from "../base-trigger-node";
import { ManualTriggerDialog } from "./dialog";

export const ManualTriggerNode = memo((props: NodeProps) => {
  const [isDialogOpen, setIsDialogOpen] = useState(false);

  const handleOpenSettings = () => {
    setIsDialogOpen(true);
  };

  const nodeStatus = "initial";

  return (
    <>
      <ManualTriggerDialog open={isDialogOpen} onOpenChange={setIsDialogOpen} />
      <BaseTriggerNode
        {...props}
        icon={MousePointer}
        name="When Clicked 'Execute Workflow'"
        onSettingsClick={handleOpenSettings}
        onDoubleClick={handleOpenSettings}
        status={nodeStatus}
      />
    </>
  );
});

ManualTriggerNode.displayName = "ManualTriggerNode";
