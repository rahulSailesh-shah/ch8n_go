import {
  useQuery,
  useMutation,
  useQueryClient,
  keepPreviousData,
} from "@tanstack/react-query";
import {
  createWorkflow,
  getWorkflows,
  type CreateWorkflowRequest,
} from "../api";
import { toast } from "sonner";
import { useSearch } from "@tanstack/react-router";

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
