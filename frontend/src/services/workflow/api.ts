import { apiClient, ApiError } from "@/lib/api-client";

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

export interface PaginatedWorkflowResponse {
  workflows: Workflow[];
  currentPage: number;
  totalPages: number;
  totalCount: number;
  hasPreviousPage: boolean;
  hasNextPage: boolean;
}

export interface GetWorkflowsRequest {
  search?: string;
  limit?: number;
  offset?: number;
}

export const createWorkflow = async (workflow: CreateWorkflowRequest) => {
  const { data, error, status } = await apiClient.post<Workflow>(
    "/workflows",
    workflow
  );
  if (error) {
    throw new ApiError(error, status);
  }
  return data;
};

export const getWorkflows = async () => {
  const { data, error, status } =
    await apiClient.get<PaginatedWorkflowResponse>("/workflows");
  if (error) {
    throw new ApiError(error, status);
  }
  return data;
};
