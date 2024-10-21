export default const dag = {
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
