import { LoadingView } from "@/components/entity-component";
import { ErrorView } from "@/components/entity-component";
import { useState, useCallback, useMemo } from "react";
import {
  ReactFlow,
  applyNodeChanges,
  applyEdgeChanges,
  addEdge,
  type Node,
  type Edge,
  type NodeChange,
  type EdgeChange,
  type Connection,
  Background,
  Controls,
  MiniMap,
  Panel,
} from "@xyflow/react";
import "@xyflow/react/dist/style.css";
import { NODE_COMPONENTS } from "@/config/node-components";
import type { WorkflowDetails } from "@/features/workflows/types";
import { AddNodeButton } from "./add-node-button";
import { editorAtom } from "../store/atoms";
import { useSetAtom } from "jotai";
import { NodeType } from "@/config/node-types";
import { ExecuteWorkflowButton } from "./execute-workflow-button";

interface EditorProps {
  workflow: WorkflowDetails;
}

export const Editor = ({ workflow }: EditorProps) => {
  const [nodes, setNodes] = useState<Node[]>(workflow.nodes);
  const [edges, setEdges] = useState<Edge[]>(workflow.edges);

  const setEditor = useSetAtom(editorAtom);

  const onNodesChange = useCallback(
    (changes: NodeChange[]) =>
      setNodes((nodesSnapshot) => applyNodeChanges(changes, nodesSnapshot)),
    []
  );

  const onEdgesChange = useCallback(
    (changes: EdgeChange[]) =>
      setEdges((edgesSnapshot) => applyEdgeChanges(changes, edgesSnapshot)),
    []
  );

  const onConnect = useCallback(
    (params: Connection) =>
      setEdges((edgesSnapshot) => addEdge(params, edgesSnapshot)),
    []
  );

  const hasManualTrigger = useMemo(
    () => nodes.some((node) => node.type === NodeType.MANUAL_TRIGGER),
    [nodes]
  );

  return (
    <div className="size-full">
      <ReactFlow
        nodes={nodes}
        edges={edges}
        onNodesChange={onNodesChange}
        onEdgesChange={onEdgesChange}
        onConnect={onConnect}
        nodeTypes={NODE_COMPONENTS}
        fitView
        onInit={setEditor}
        proOptions={{
          hideAttribution: true,
        }}
        snapGrid={[10, 10]}
        snapToGrid
        panOnScroll
        panOnDrag={false}
        selectionOnDrag
      >
        <Background />
        <Controls />
        <MiniMap />
        <Panel position="top-right">
          <AddNodeButton />
        </Panel>
        {hasManualTrigger && (
          <Panel position="bottom-center">
            <ExecuteWorkflowButton workflowId={workflow.workflowId} />
          </Panel>
        )}
      </ReactFlow>
    </div>
  );
};

export const WorkflowLoadingView = () => {
  return <LoadingView message="Loading editor..." />;
};

export const WorkflowErrorView = () => {
  return <ErrorView message="An error occurred while loading the workflow" />;
};
