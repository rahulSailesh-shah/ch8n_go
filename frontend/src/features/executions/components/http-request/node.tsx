import { memo } from "react";
import { BaseExecutionNode } from "../base-execution-node";
import type { Node, NodeProps } from "@xyflow/react";
import { GlobeIcon } from "lucide-react";

type HttpRequestNodeData = {
  endpoint?: string;
  method?: "GET" | "POST" | "PUT" | "DELETE" | "PATCH";
  body?: string;
  [key: string]: unknown;
};

type HttpRequestNodeType = Node<HttpRequestNodeData>;

export const HttpRequestNode = memo((props: NodeProps<HttpRequestNodeType>) => {
  const nodeData = props.data as HttpRequestNodeData;
  const description = nodeData.endpoint
    ? `${nodeData.method || "GET"}: ${nodeData.endpoint}`
    : "Not Configured";

  return (
    <>
      <BaseExecutionNode
        {...props}
        description={description}
        icon={GlobeIcon}
        name="HTTP Request"
        onSettingsClick={() => {}}
        onDoubleClick={() => {}}
      />
    </>
  );
});

HttpRequestNode.displayName = "HttpRequestNode";
