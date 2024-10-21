const dag = {
  nodes: {
    "node1": { /* Node data */ },
    "node2": { /* Node data */ },
    "node3": { /* Node data */ },
    "node4": { /* Node data */ },
  },
  edges: [
    { source: "node1", target: "node2" },
    { source: "node1", target: "node3" },
    { source: "node2", target: "node4" },
  ]
};

function traverseDAG(startNode, visit) {
  const visited = new Set();
  const stack = [startNode];

  while (stack.length > 0) {
    const node = stack.pop();

    if (!visited.has(node)) {
      visit(node);
      visited.add(node);

      const outgoingEdges = dag.edges.filter(edge => edge.source === node);
      outgoingEdges.forEach(edge => stack.push(edge.target));
    }
  }
}

traverseDAG("node1", node => console.log(node));
