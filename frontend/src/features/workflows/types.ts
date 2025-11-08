import type { Edge, Node } from "@xyflow/react";

export interface CreateWorkflowRequest {
  id: string;
  name: string;
}

export interface UpdateWorkflowRequest {
  id: string;
  nodes: Node[];
  edges: Edge[];
}

export interface NodePosition {
  x: number;
  y: number;
}

export interface WorkflowNode {
  id: number;
  workflowId: number;
  name: string;
  type: string;
  position: NodePosition;
  data: Record<string, unknown>;
  createdAt: string;
  updatedAt: string;
}

export interface WorkflowEdge {
  id: number;
  workflowId: number;
  sourceNodeId: number;
  targetNodeId: number;
  fromOutput: string;
  toInput: string;
  createdAt: string;
  updatedAt: string;
}

export interface WorkflowDTO {
  id: string;
  name: string;
  description: string;
  userId: string;
  createdAt: string;
  updatedAt: string;
  nodes: WorkflowNode[];
  edges: WorkflowEdge[];
}

export interface PaginatedWorkflowResponse {
  workflows: WorkflowDTO[];
  currentPage: number;
  totalPages: number;
  totalCount: number;
  hasPreviousPage: boolean;
  hasNextPage: boolean;
}

export interface GetWorkflowsRequest {
  search?: string;
  limit?: number;
  offset?: number;
}

export interface WorkflowDetails {
  workflowId: string;
  workflowName: string;
  nodes: Node[];
  edges: Edge[];
}
