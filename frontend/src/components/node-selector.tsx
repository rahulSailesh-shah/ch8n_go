import { NodeType } from "@/config/node-types";
import { useReactFlow } from "@xyflow/react";
import { GlobeIcon, MousePointerIcon, WebhookIcon } from "lucide-react";
import { useCallback, type ReactNode } from "react";
import { toast } from "sonner";
import { v4 as uuidv4 } from "uuid";
import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "./ui/sheet";
import { Separator } from "./ui/separator";
import { INITIAL_NODE_VARIABLE_NAME } from "@/config/node-components";

export type NodeTypeOption = {
  type: NodeType;
  label: string;
  description: string;
  icon: React.ComponentType<{ className?: string }> | string;
};

const triggerNodes: NodeTypeOption[] = [
  {
    type: NodeType.MANUAL_TRIGGER,
    label: "Trigger Manually",
    description:
      "Runs the workflow on clicking a button. Good for manual testing.",
    icon: MousePointerIcon,
  },
  {
    type: NodeType.WEBHOOK_TRIGGER,
    label: "Trigger via Webhook",
    description: "Runs the workflow via a webhook.",
    icon: WebhookIcon,
  },
];

const executionNodes: NodeTypeOption[] = [
  {
    type: NodeType.HTTP_TRIGGER,
    label: "HTTP Request",
    description: "Trigger the workflow via HTTP",
    icon: GlobeIcon,
  },
];

interface NodeSelectorProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  children: ReactNode;
}

export const NodeSelector = ({
  open,
  onOpenChange,
  children,
}: NodeSelectorProps) => {
  const { setNodes, getNodes, screenToFlowPosition } = useReactFlow();

  const handleAddNode = useCallback(
    (nodeType: NodeTypeOption) => {
      if (nodeType.type === NodeType.MANUAL_TRIGGER) {
        const nodes = getNodes();
        const hasManualTrigger = nodes.some(
          (node) => node.type === NodeType.MANUAL_TRIGGER
        );
        if (hasManualTrigger) {
          toast.error("Only one manual trigger is allowed");
          return;
        }
      }

      setNodes((nodes) => {
        const hasInitialTrigger = nodes.some(
          (node) => node.type === NodeType.INITIAL
        );

        const centerX = window.innerWidth / 2;
        const centerY = window.innerHeight / 2;
        const flowPosition = screenToFlowPosition({
          x: centerX + (Math.random() - 0.5) * 200,
          y: centerY + (Math.random() - 0.5) * 200,
        });

        const newNode = {
          id: uuidv4(),
          type: nodeType.type,
          name: nodeType.type,
          position: flowPosition,
          data: { variableName: INITIAL_NODE_VARIABLE_NAME[nodeType.type] },
        };
        if (hasInitialTrigger) {
          return [newNode];
        }
        return [...nodes, newNode];
      });
      onOpenChange(false);
    },
    [getNodes, screenToFlowPosition, setNodes, onOpenChange]
  );

  return (
    <Sheet open={open} onOpenChange={onOpenChange}>
      <SheetTrigger asChild>{children}</SheetTrigger>
      <SheetContent side="right" className="w-full sm:max-w-md overflow-y-auto">
        <div className="flex flex-col gap-4">
          <SheetHeader>
            <SheetTitle>What triggers this workflow?</SheetTitle>
            <SheetDescription>
              A trigger is the first node in a workflow. It starts the workflow
              when a specific event occurs.
            </SheetDescription>
          </SheetHeader>
          <div>
            {triggerNodes.map((node) => {
              const Icon = node.icon;
              return (
                <div
                  key={node.type}
                  className="w-full justify-start h-auto py-5 px-4 rounded-none cursor-pointer border-l-2 border-transparent hover:border-l-primary"
                  onClick={() => handleAddNode(node)}
                >
                  <div className="flex items-center gap-6 w-full overflow-hidden">
                    {typeof Icon === "string" ? (
                      <img
                        src={Icon}
                        alt={node.label}
                        className="size-5 object-contain rounded-sm"
                      />
                    ) : (
                      <Icon className="size-5" />
                    )}
                    <div className="flex flex-col items-start text-left">
                      <span className="text-sm font-medium">{node.label}</span>
                      <span className="text-xs text-muted-foreground mt-1">
                        {node.description}
                      </span>
                    </div>
                  </div>
                </div>
              );
            })}
          </div>
          <Separator />
          <div>
            {executionNodes.map((node) => {
              const Icon = node.icon;
              return (
                <div
                  key={node.type}
                  className="w-full justify-start h-auto py-5 px-4 rounded-none cursor-pointer border-l-2 border-transparent hover:border-l-primary"
                  onClick={() => handleAddNode(node)}
                >
                  <div className="flex items-center gap-6 w-full overflow-hidden">
                    {typeof Icon === "string" ? (
                      <img
                        src={Icon}
                        alt={node.label}
                        className="size-5 object-contain rounded-sm"
                      />
                    ) : (
                      <Icon className="size-5" />
                    )}
                    <div className="flex flex-col items-start text-left">
                      <span className="text-sm font-medium">{node.label}</span>
                      <span className="text-xs text-muted-foreground mt-1">
                        {node.description}
                      </span>
                    </div>
                  </div>
                </div>
              );
            })}
          </div>
        </div>
      </SheetContent>
    </Sheet>
  );
};
