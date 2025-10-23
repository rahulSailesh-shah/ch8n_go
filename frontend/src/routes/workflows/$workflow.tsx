import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/workflows/$workflow')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/workflows/$workflow"!</div>
}
