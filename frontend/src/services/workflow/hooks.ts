import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import {
  createWorkflow,
  getWorkflows,
  type CreateWorkflowRequest,
} from "./api";
import { toast } from "sonner";

export const useWorkflows = () => {
  return useQuery({
    queryKey: ["workflows"],
    queryFn: () => getWorkflows(),
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
