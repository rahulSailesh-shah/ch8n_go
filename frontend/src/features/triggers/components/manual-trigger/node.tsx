import { memo } from "react";
import type { NodeProps } from "@xyflow/react";
import { MousePointer } from "lucide-react";
import { BaseTriggerNode } from "../base-trigger-node";

export const ManualTriggerNode = memo((props: NodeProps) => {
  return (
    <>
      <BaseTriggerNode
        {...props}
        icon={MousePointer}
        name="When Clicked 'Execute Workflow'"
        onSettingsClick={() => {}}
        onDoubleClick={() => {}}
      />
    </>
  );
});

ManualTriggerNode.displayName = "ManualTriggerNode";
