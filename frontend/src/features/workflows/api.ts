import { apiClient, ApiError } from "@/lib/api-client";
import type { WorkflowSearchParams } from "@/routes/_authenticated/workflows";

export interface CreateWorkflowRequest {
  name: string;
  description?: string;
}

export interface UpdateWorkflowRequest {
  name?: string;
  description?: string;
}

export interface Workflow {
  id: string;
  name: string;
  description: string;
  userId: string;
  createdAt: string;
  updatedAt: string;
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

export const getWorkflows = async (params: WorkflowSearchParams) => {
  const { data, error, status } =
    await apiClient.get<PaginatedWorkflowResponse>("/workflows", params);
  if (error) {
    if (status === 403) {
      return {
        workflows: [],
        currentPage: 0,
        totalPages: 0,
        totalCount: 0,
        hasPreviousPage: false,
        hasNextPage: false,
      };
    }
    throw new ApiError(error, status);
  }
  return data;
};

export const deleteWorkflow = async (id: string) => {
  const { error, status } = await apiClient.delete(`/workflows/${id}`);
  if (error) {
    throw new ApiError(error, status);
  }
};

export const getWorkflow = async (id: string) => {
  const { data, error, status } = await apiClient.get<Workflow>(
    `/workflows/${id}`
  );
  if (error) {
    throw new ApiError(error, status);
  }
  return data;
};

export const updateWorkflow = async (
  id: string,
  workflow: UpdateWorkflowRequest
) => {
  const { data, error, status } = await apiClient.put<Workflow>(
    `/workflows/${id}`,
    workflow
  );
  if (error) {
    throw new ApiError(error, status);
  }
  return data;
};
