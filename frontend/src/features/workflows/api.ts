import { apiClient, ApiError } from "@/lib/api-client";
import type { WorkflowSearchParams } from "@/routes/_authenticated/workflows";
import type {
  CreateWorkflowRequest,
  PaginatedWorkflowResponse,
  WorkflowDetails,
  WorkflowDTO,
} from "./types";

const handleApiError = (errorMsg: string, status: number) => {
  throw new ApiError(errorMsg, status);
};

const EMPTY_PAGINATED_RESPONSE: PaginatedWorkflowResponse = {
  workflows: [],
  currentPage: 0,
  totalPages: 0,
  totalCount: 0,
  hasPreviousPage: false,
  hasNextPage: false,
};

export const createWorkflow = async (workflow: CreateWorkflowRequest) => {
  const { data, error, status } = await apiClient.post<WorkflowDTO>(
    "/workflows",
    workflow
  );
  if (error) handleApiError(error, status);
  return data;
};

export const getWorkflows = async (params: WorkflowSearchParams) => {
  const { data, error, status } =
    await apiClient.get<PaginatedWorkflowResponse>("/workflows", params);

  if (error) {
    if (status === 403) return EMPTY_PAGINATED_RESPONSE;
    handleApiError(error, status);
  }

  return data;
};

export const deleteWorkflow = async (id: string) => {
  const { error, status } = await apiClient.delete(`/workflows/${id}`);
  if (error) handleApiError(error, status);
};

export const updateWorkflow = async (
  id: string,
  workflow: CreateWorkflowRequest
) => {
  const { data, error, status } = await apiClient.put<WorkflowDTO>(
    `/workflows/${id}`,
    workflow
  );
  if (error) handleApiError(error, status);
  return data;
};

export const getWorkflow = async (
  id: string
): Promise<WorkflowDetails | null> => {
  const { data, error, status } = await apiClient.get<WorkflowDTO>(
    `/workflows/${id}`
  );
  if (error) handleApiError(error, status);
  if (!data) return null;

  return {
    workflowId: data.id,
    workflowName: data.name,
    nodes: (data.nodes ?? []).map((node) => ({
      id: node.id.toString(),
      type: node.type,
      position: { x: node.position.x, y: node.position.y },
      data: node.data,
    })),
    edges: (data.edges ?? []).map((edge) => ({
      id: edge.id.toString(),
      source: edge.sourceNodeId.toString(),
      target: edge.targetNodeId.toString(),
      fromOutput: edge.fromOutput,
      toInput: edge.toInput,
    })),
  };
};
