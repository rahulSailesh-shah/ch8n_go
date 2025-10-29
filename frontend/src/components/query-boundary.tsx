import type { ReactNode } from "react";
import type { UseQueryResult } from "@tanstack/react-query";
import {
  DefaultLoadingFallback,
  DefaultErrorFallback,
  DefaultEmptyFallback,
} from "./entity-component";

interface QueryBoundaryProps<T> {
  query: UseQueryResult<T | null | undefined>;
  loadingFallback?: ReactNode;
  errorFallback?: ReactNode;
  emptyFallback?: ReactNode;
  children: (data: NonNullable<T>) => ReactNode;
}

export function QueryBoundary<T>({
  query,
  loadingFallback = <DefaultLoadingFallback />,
  errorFallback = <DefaultErrorFallback />,
  emptyFallback = <DefaultEmptyFallback />,
  children,
}: QueryBoundaryProps<T>) {
  if (query.isLoading) {
    return <>{loadingFallback}</>;
  }

  if (query.isError) {
    return <>{errorFallback}</>;
  }

  if (query.data === null || query.data === undefined) {
    return <>{emptyFallback}</>;
  }

  return <>{children(query.data)}</>;
}
