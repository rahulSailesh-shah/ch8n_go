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
  });
};

export const useCreateWorkflow = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (workflow: CreateWorkflowRequest) => createWorkflow(workflow),
    onError: ({ message }) => {
      toast.error(message);
    },
    onSettled: () => queryClient.invalidateQueries({ queryKey: ["workflows"] }),
  });
};
