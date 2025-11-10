import {
  useQuery,
  useMutation,
  useQueryClient,
  keepPreviousData,
} from "@tanstack/react-query";
import {
  createWorkflow,
  deleteWorkflow,
  executeWorkflow,
  getWorkflow,
  getWorkflows,
  updateWorkflow,
  updateWorkflowName,
} from "../api";
import { toast } from "sonner";
import { useSearch } from "@tanstack/react-router";
import type { CreateWorkflowRequest, UpdateWorkflowRequest } from "../types";

export const useQueryWorkflows = () => {
  const search = useSearch({
    from: "/_authenticated/workflows/",
  });
  return useQuery({
    queryKey: ["workflows", search],
    queryFn: () => getWorkflows(search),
    retry: 0,
    placeholderData: keepPreviousData,
  });
};

export const useQueryWorkflow = (id: string) => {
  return useQuery({
    queryKey: ["workflow", id],
    queryFn: () => getWorkflow(id),
    retry: 0,
  });
};

export const useCreateWorkflow = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (workflow: CreateWorkflowRequest) => createWorkflow(workflow),
    onSuccess: (workflow) => {
      toast.success(`Workflow ${workflow?.name} created`);
      queryClient.invalidateQueries({ queryKey: ["workflows"] });
    },
    onError: ({ message }) => {
      toast.error(message);
    },
  });
};

export const useDeleteWorkflow = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (id: string) => deleteWorkflow(id),
    onSuccess: () => {
      toast.success(`Workflow deleted`);
      queryClient.invalidateQueries({ queryKey: ["workflows"] });
    },
    onError: ({ message }) => {
      toast.error(message);
    },
  });
};

export const useUpdateWorkflowName = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (payload: CreateWorkflowRequest) => updateWorkflowName(payload),
    onSuccess: (workflow, variables) => {
      toast.success(`Workflow ${workflow?.name} updated`);
      queryClient.invalidateQueries({
        queryKey: ["workflow", variables.id],
      });
      queryClient.invalidateQueries({
        queryKey: ["workflows"],
      });
    },
    onError: ({ message }) => {
      toast.error(message);
    },
  });
};

export const useUpdateWorkflow = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (payload: UpdateWorkflowRequest) => updateWorkflow(payload),
    onSuccess: (workflow, variables) => {
      toast.success(`Workflow ${workflow?.name} updated`);
      queryClient.invalidateQueries({
        queryKey: ["workflow", variables.id],
      });
      queryClient.invalidateQueries({
        queryKey: ["workflows"],
      });
    },
    onError: ({ message }) => {
      toast.error(message);
    },
  });
};

export const useExecuteWorkflow = () => {
  return useMutation({
    mutationFn: (id: string) => executeWorkflow(id),
    onSuccess: (data) => {
      toast.success(`Workflow: ${data?.name} executed`);
    },
    onError: ({ message }) => {
      toast.error(message);
    },
  });
};
