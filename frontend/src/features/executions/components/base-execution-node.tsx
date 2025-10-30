import React, { memo } from "react";
import {
  BaseNode,
  BaseNodeContent,
} from "../../../components/react-flow/base-node";
import { BaseHandle } from "../../../components/react-flow/base-handle";
import { WorkflowNode } from "../../../components/workflow-node";
import { Position, useReactFlow, type NodeProps } from "@xyflow/react";
import type { LucideIcon } from "lucide-react";

interface BaseExecutionNodeProps extends NodeProps {
  icon: LucideIcon | string;
  name: string;
  description?: string;
  children?: React.ReactNode;
  // status?: NodeStatus;
  onSettingsClick?: () => void;
  onDoubleClick?: () => void;
}

export const BaseExecutionNode = memo(
  ({
    id,
    icon: Icon,
    name,
    description,
    children,
    onSettingsClick,
    onDoubleClick,
  }: BaseExecutionNodeProps) => {
    const { setNodes, setEdges } = useReactFlow();

    const handleDelete = () => {
      setNodes((nodes) => nodes.filter((node) => node.id !== id));
      setEdges((edges) =>
        edges.filter((edge) => edge.source !== id && edge.target !== id)
      );
    };

    return (
      <div>
        <WorkflowNode
          name={name}
          description={description}
          onSettings={onSettingsClick}
          onDelete={handleDelete}
        >
          <BaseNode onDoubleClick={onDoubleClick}>
            <BaseNodeContent>
              {typeof Icon === "string" ? (
                <img
                  src={Icon}
                  alt={name}
                  width={16}
                  height={16}
                  className="size-5 object-contain rounded-sm"
                />
              ) : (
                <Icon className="size-5 text-muted-foreground" />
              )}
              {children}
              <BaseHandle
                id="target-1"
                type="target"
                position={Position.Left}
              />
              <BaseHandle
                id="source-1"
                type="source"
                position={Position.Right}
              />
            </BaseNodeContent>
          </BaseNode>
        </WorkflowNode>
      </div>
    );
  }
);

BaseExecutionNode.displayName = "BaseExecutionNode";
