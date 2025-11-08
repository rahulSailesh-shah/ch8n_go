import { memo, useState } from "react";
import { BaseExecutionNode } from "../base-execution-node";
import { useReactFlow, type Node, type NodeProps } from "@xyflow/react";
import { GlobeIcon } from "lucide-react";
import { HTTPRequestDialog, type HTTPRequestFormValues } from "./dialog";

type HttpRequestNodeData = {
  endpoint?: string;
  method?: "GET" | "POST" | "PUT" | "DELETE" | "PATCH";
  body?: string;
};

type HttpRequestNodeType = Node<HttpRequestNodeData>;

export const HttpRequestNode = memo((props: NodeProps<HttpRequestNodeType>) => {
  const [openDialog, setOpenDialog] = useState(false);
  const { setNodes } = useReactFlow();

  const handleOpenSettings = () => {
    setOpenDialog(true);
  };

  const handleSubmit = (values: HTTPRequestFormValues) => {
    setNodes((nodes) =>
      nodes.map((node) =>
        node.id === props.id
          ? { ...node, data: { ...node.data, ...values } }
          : node
      )
    );
  };

  const nodeData = props.data;
  const description = nodeData.endpoint
    ? `${nodeData.method || "GET"}: ${nodeData.endpoint}`
    : "Not Configured";

  const nodeStatus = "initial";
  return (
    <>
      <HTTPRequestDialog
        open={openDialog}
        onOpenChange={setOpenDialog}
        onSubmit={handleSubmit}
        defaultValues={nodeData}
      />
      <BaseExecutionNode
        {...props}
        description={description}
        icon={GlobeIcon}
        name="HTTP Request"
        onSettingsClick={handleOpenSettings}
        onDoubleClick={handleOpenSettings}
        status={nodeStatus}
      />
    </>
  );
});

HttpRequestNode.displayName = "HttpRequestNode";
