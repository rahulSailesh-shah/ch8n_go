import { memo, useState } from "react";
import { useReactFlow, type Node, type NodeProps } from "@xyflow/react";
import { MousePointer } from "lucide-react";
import { BaseTriggerNode } from "../base-trigger-node";
import { ManualTriggerDialog } from "./dialog";

type ManualTriggerNodeData = {
  variableName: string;
  input?: string;
};

type ManualTriggerNodeType = Node<ManualTriggerNodeData>;

export const ManualTriggerNode = memo(
  (props: NodeProps<ManualTriggerNodeType>) => {
    const [isDialogOpen, setIsDialogOpen] = useState(false);
    const { setNodes } = useReactFlow();

    const handleOpenSettings = () => {
      setIsDialogOpen(true);
    };

    const nodeData = props.data;

    const nodeStatus = "initial";

    const handleSubmit = (values: ManualTriggerNodeData) => {
      setNodes((nodes) =>
        nodes.map((node) =>
          node.id === props.id
            ? { ...node, data: { ...node.data, ...values } }
            : node
        )
      );
    };

    return (
      <>
        <ManualTriggerDialog
          open={isDialogOpen}
          onOpenChange={setIsDialogOpen}
          onSubmit={handleSubmit}
          defaultValues={nodeData}
        />
        <BaseTriggerNode
          {...props}
          icon={MousePointer}
          name="Manual Trigger"
          onSettingsClick={handleOpenSettings}
          onDoubleClick={handleOpenSettings}
          status={nodeStatus}
        />
      </>
    );
  }
);

ManualTriggerNode.displayName = "ManualTriggerNode";
