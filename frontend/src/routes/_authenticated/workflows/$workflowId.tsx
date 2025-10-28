import { Editor } from "@/features/editor/components/editor";
import EditorHeader from "@/features/editor/components/editor-header";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/_authenticated/workflows/$workflowId")({
  component: RouteComponent,
});

function RouteComponent() {
  const { workflowId } = Route.useParams();
  return (
    <>
      <EditorHeader workflowId={workflowId} />
      <main className="flex-1">
        <Editor workflowId={workflowId} />
      </main>
    </>
  );
}
