import { memo, useState } from "react";
import { BaseExecutionNode } from "../base-execution-node";
import { useReactFlow, type Node, type NodeProps } from "@xyflow/react";
import { GlobeIcon } from "lucide-react";
import { HTTPRequestDialog, type HTTPFormType } from "./dialog";

type HttpRequestNodeData = {
  endpoint?: string;
  method?: "GET" | "POST" | "PUT" | "DELETE" | "PATCH";
  body?: string;
  [key: string]: unknown;
};

type HttpRequestNodeType = Node<HttpRequestNodeData>;

export const HttpRequestNode = memo((props: NodeProps<HttpRequestNodeType>) => {
  const [openDialog, setOpenDialog] = useState(false);
  const { setNodes } = useReactFlow();

  const handleOpenSettings = () => {
    setOpenDialog(true);
  };

  const handleSubmit = (data: HTTPFormType) => {
    setNodes((nodes) =>
      nodes.map((node) =>
        node.id === props.id
          ? { ...node, data: { ...node.data, ...data } }
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
        defaultEndpoint={nodeData.endpoint}
        defaultMethod={nodeData.method}
        defaultBody={nodeData.body}
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
