import { apiClient } from "@/lib/api-client";

export interface CreateWorkflowRequest {
  name: string;
  description: string;
}

export interface Workflow {
  id: string;
  name: string;
  description: string;
  userId: string;
  createdAt: string;
}

export const createWorkflow = async (workflow: CreateWorkflowRequest) => {
  const response = await apiClient.post<Workflow>("/workflows", workflow);
  return response.data;
};

export const getWorkflows = async () => {
  const response = await apiClient.get<Workflow[]>("/workflows");
  return response.data;
};
