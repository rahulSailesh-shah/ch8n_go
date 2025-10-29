import type { NodeTypes } from "@xyflow/react";
import { InitialNode } from "@/components/initial-node";

export const NODE_COMPONENTS = {
  INITIAL: InitialNode,
} as const satisfies NodeTypes;

export type RegisteredNodeTypes = keyof typeof NODE_COMPONENTS;
