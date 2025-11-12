import { InitialNode } from "@/components/initial-node";
import { NodeType } from "./node-types";
import type { NodeTypes } from "@xyflow/react";
import { HttpRequestNode } from "@/features/executions/components/http-request/node";
import { ManualTriggerNode } from "@/features/triggers/components/manual-trigger/node";
import { WebhookTriggerNode } from "@/features/triggers/components/webhook-trigger/node";

export const NODE_COMPONENTS = {
  [NodeType.INITIAL]: InitialNode,
  [NodeType.HTTP_TRIGGER]: HttpRequestNode,
  [NodeType.MANUAL_TRIGGER]: ManualTriggerNode,
  [NodeType.WEBHOOK_TRIGGER]: WebhookTriggerNode,
} as const satisfies NodeTypes;

export type RegisteredNodeTypes = keyof typeof NODE_COMPONENTS;

export const INITIAL_NODE_VARIABLE_NAME = {
  [NodeType.INITIAL]: "INITIAL",
  [NodeType.HTTP_TRIGGER]: "HTTP_TRIGGER",
  [NodeType.MANUAL_TRIGGER]: "MANUAL_TRIGGER",
  [NodeType.WEBHOOK_TRIGGER]: "WEBHOOK_TRIGGER",
} as const satisfies Record<RegisteredNodeTypes, string>;
