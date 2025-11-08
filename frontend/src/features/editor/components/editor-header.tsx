import { ThemeToggle } from "@/components/theme-toggle";
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbSeparator,
} from "@/components/ui/breadcrumb";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { SidebarTrigger } from "@/components/ui/sidebar";
import {
  useUpdateWorkflow,
  useUpdateWorkflowName,
} from "@/features/workflows/hooks/use-workflows";
import type { WorkflowDetails } from "@/features/workflows/types";
import { Link } from "@tanstack/react-router";
import { useAtomValue } from "jotai";
import { SaveIcon } from "lucide-react";
import { useEffect, useRef, useState } from "react";
import { editorAtom } from "../store/atoms";

export const EditorNameInput = ({
  workflow,
}: {
  workflow: WorkflowDetails;
}) => {
  const updateWorkflow = useUpdateWorkflowName();
  const [isEditing, setIsEditing] = useState(false);
  const [name, setName] = useState("");
  const inputRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    if (workflow?.workflowName) {
      setName(workflow?.workflowName || "");
    }
  }, [workflow?.workflowName]);

  useEffect(() => {
    if (isEditing && inputRef.current) {
      inputRef.current.focus();
      inputRef.current.select();
    }
  }, [isEditing]);

  const handleSave = async () => {
    if (name === workflow?.workflowName) {
      setIsEditing(false);
      return;
    }

    try {
      if (inputRef.current) {
        updateWorkflow.mutateAsync({
          id: workflow.workflowId.toString(),
          name: inputRef.current?.value || "",
        });
        setIsEditing(false);
      }
    } catch {
      setName(workflow?.workflowName || "");
    } finally {
      setIsEditing(false);
    }
  };

  const handleKeyDown = (e: React.KeyboardEvent) => {
    if (e.key === "Enter") {
      handleSave();
    } else if (e.key === "Escape") {
      setName(workflow?.workflowName || "");
      setIsEditing(false);
    }
  };

  if (isEditing) {
    return (
      <BreadcrumbItem className="cursor-pointer hover:text-foreground transition-colors">
        <Input
          ref={inputRef}
          value={name}
          onChange={(e) => setName(e.target.value)}
          onBlur={handleSave}
          onKeyDown={handleKeyDown}
          className="w-full text-foreground bg-transparent border-none outline-none"
          disabled={updateWorkflow.isPending}
        />
      </BreadcrumbItem>
    );
  }

  return (
    <BreadcrumbItem
      className="cursor-pointer hover:text-foreground transition-colors"
      onClick={() => setIsEditing(true)}
    >
      {workflow.workflowName}
    </BreadcrumbItem>
  );
};

export const EditorBreadcrumb = ({
  workflow,
}: {
  workflow: WorkflowDetails;
}) => {
  return (
    <Breadcrumb>
      <BreadcrumbList>
        <BreadcrumbItem>
          <BreadcrumbLink asChild>
            <Link to={`/workflows`} search={{ page: 1, limit: 5, search: "" }}>
              Workflows
            </Link>
          </BreadcrumbLink>
        </BreadcrumbItem>
        <BreadcrumbSeparator />
        <EditorNameInput workflow={workflow} />
      </BreadcrumbList>
    </Breadcrumb>
  );
};

export const EditorSaveButton = ({ workflowId }: { workflowId: string }) => {
  const editor = useAtomValue(editorAtom);

  const saveWorkflow = useUpdateWorkflow();

  const handleSave = () => {
    if (!editor) {
      return;
    }
    saveWorkflow.mutateAsync({
      id: workflowId.toString(),
      nodes: editor.getNodes(),
      edges: editor.getEdges(),
    });
  };

  return (
    <div className="ml-auto">
      <Button onClick={handleSave} disabled={saveWorkflow.isPending}>
        <SaveIcon className="size-4" />
        Save
      </Button>
    </div>
  );
};

const EditorHeader = ({ workflow }: { workflow: WorkflowDetails }) => {
  return (
    <div className="flex h-14 shrink-0 items-center gap-2 border-b px-4 bg-background">
      <SidebarTrigger />
      <div className="flex flex-row items-center justify-between gap-x-4 w-full">
        <EditorBreadcrumb workflow={workflow} />
        <EditorSaveButton workflowId={workflow.workflowId} />
      </div>
      <div className="ml-auto">
        <ThemeToggle />
      </div>
    </div>
  );
};

export default EditorHeader;
