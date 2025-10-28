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
  useQueryWorkflow,
  useUpdateWorkflow,
} from "@/features/workflows/hooks/use-workflows";
import { Link } from "@tanstack/react-router";
import { SaveIcon } from "lucide-react";
import { useEffect, useRef, useState } from "react";

export const EditorNameInput = ({ workflowId }: { workflowId: string }) => {
  const { data: workflow } = useQueryWorkflow(workflowId);
  const updateWorkflow = useUpdateWorkflow();
  const [isEditing, setIsEditing] = useState(false);
  const [name, setName] = useState("");
  const inputRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    if (workflow?.name) {
      setName(workflow?.name || "");
    }
  }, [workflow?.name]);

  useEffect(() => {
    if (isEditing && inputRef.current) {
      inputRef.current.focus();
      inputRef.current.select();
    }
  }, [isEditing]);

  const handleSave = async () => {
    if (name === workflow?.name) {
      setIsEditing(false);
      return;
    }

    try {
      if (inputRef.current) {
        updateWorkflow.mutateAsync({
          id: workflowId,
          workflow: {
            name: inputRef.current?.value || "",
          },
        });
        setIsEditing(false);
      }
    } catch {
      setName(workflow?.name || "");
    } finally {
      setIsEditing(false);
    }
  };

  const handleKeyDown = (e: React.KeyboardEvent) => {
    if (e.key === "Enter") {
      handleSave();
    } else if (e.key === "Escape") {
      setName(workflow?.name || "");
      setIsEditing(false);
    }
  };

  if (!workflow) {
    return null;
  }

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
      {workflow.name}
    </BreadcrumbItem>
  );
};

export const EditorBreadcrumb = ({ workflowId }: { workflowId: string }) => {
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
        <EditorNameInput workflowId={workflowId} />
      </BreadcrumbList>
    </Breadcrumb>
  );
};

export const EditorSaveButton = ({ workflowId }: { workflowId: string }) => {
  return (
    <div className="ml-auto">
      <Button onClick={() => console.log(workflowId)} disabled={false}>
        <SaveIcon className="size-4" />
        Save
      </Button>
    </div>
  );
};

const EditorHeader = ({ workflowId }: { workflowId: string }) => {
  return (
    <div className="flex h-14 shrink-0 items-center gap-2 border-b px-4 bg-background">
      <SidebarTrigger />
      <div className="flex flex-row items-center justify-between gap-x-4 w-full">
        <EditorBreadcrumb workflowId={workflowId} />
        <EditorSaveButton workflowId={workflowId} />
      </div>
      <div className="ml-auto">
        <ThemeToggle />
      </div>
    </div>
  );
};

export default EditorHeader;
