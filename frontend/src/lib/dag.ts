type Node = {
  id: string;
  deps: string[];
};

type Edge = {
  from: string;
  to: string;
};

class DAG {
  nodes: Node[];
  edges: Edge[];

  constructor() {
    this.nodes = [];
    this.edges = [];
  }

  addNode(node: Node) {
    this.nodes.push(node);
  }

  addEdge(edge: Edge) {
    this.edges.push(edge);
  }
}

class ExecutionEngine {
  dag: DAG;
  nodeMap: Map<string, Node>;
  adjList: Map<string, string[]>;
  inDegree: Map<string, number>;

  constructor(dag: DAG) {
    this.dag = dag;
    this.nodeMap = new Map();
    this.adjList = new Map();
    this.inDegree = new Map();

    this.buildGraph();
  }

  buildGraph() {
    this.dag.nodes.forEach((node) => {
      this.nodeMap.set(node.id, node);
      this.adjList.set(node.id, []);
      this.inDegree.set(node.id, 0);
    });

    this.dag.edges.forEach((edge) => {
      this.adjList.get(edge.from)!.push(edge.to);
      this.inDegree.set(edge.to, this.inDegree.get(edge.to)! + 1);
    });

    this.dag.edges.forEach((edge) => {
      const to = edge.to;
      this.nodeMap.get(to)!.deps.push(edge.from);
    });
  }

  validateGraph() {
    const visitedStack = new Set<string>();
    const recStack = new Set<string>();
    const hasCycle = (node: string) => {
      if (recStack.has(node)) {
        return true;
      }
      if (visitedStack.has(node)) {
        return false;
      }
      recStack.add(node);
      visitedStack.add(node);
      const neighbors = this.adjList.get(node)!;
      for (const neighbor of neighbors) {
        if (hasCycle(neighbor)) {
          return true;
        }
      }
      recStack.delete(node);
      return false;
    };
    this.nodeMap.forEach((node) => {
      if (hasCycle(node.id)) {
        throw new Error(`Cycle detected in node ${node.id}`);
      }
    });
  }

  getLevelOrder() {
    const inDegCopy = new Map(this.inDegree);
    const levels: string[][] = [];

    while (inDegCopy.size > 0) {
      const currLevel: string[] = [];
      this.nodeMap.forEach((node) => {
        if (inDegCopy.get(node.id) === 0) {
          currLevel.push(node.id);
          inDegCopy.delete(node.id);
        }
      });
      levels.push(currLevel);

      for (const node of currLevel) {
        inDegCopy.delete(node);
        const neighbors = this.adjList.get(node)!;
        for (const neighbor of neighbors) {
          inDegCopy.set(neighbor, inDegCopy.get(neighbor)! - 1);
        }
      }
    }
    return levels;
  }

  delay(ms: number) {
    return new Promise((resolve) => setTimeout(resolve, ms));
  }

  async execute() {
    this.validateGraph();
    const levelOrder = this.getLevelOrder();

    console.log(levelOrder.flat());

    for (let i = 0; i < levelOrder.length; i++) {
      const level = levelOrder[i];
      console.log("--".repeat(20));
      console.log(`Executing Level ${i + 1}: ${level}`);
      console.log("--".repeat(20));
      await Promise.all(level.map((node) => this.executeNode(node)));
    }
  }

  async executeNode(nodeId: string) {
    const node = this.nodeMap.get(nodeId);
    if (!node) {
      throw new Error(`Node ${nodeId} not found`);
    }

    await this.delay(1000);
    console.log(`Node ${node.id} executed`);
  }
}

const dag = new DAG();
dag.addNode({ id: "1", deps: [] });
dag.addNode({ id: "2", deps: [] });
dag.addNode({ id: "3", deps: [] });
dag.addNode({ id: "4", deps: [] });

dag.addEdge({ from: "1", to: "2" });
dag.addEdge({ from: "1", to: "3" });
dag.addEdge({ from: "2", to: "4" });
dag.addEdge({ from: "3", to: "4" });

const engine = new ExecutionEngine(dag);
await engine.execute();
