// Replace the setTimeout part in handleSubmit with actual API call
const handleSubmit = async (e) => {
  e.preventDefault();
  if (input.trim() === "") return;

  const userMessage = { text: input, sender: "user" };
  setMessages([...messages, userMessage]);
  setInput("");

  // Simulate API call to a chatbot service
  const response = await fetch("/api/chatbot", {
    method: "POST",
    body: JSON.stringify({ message: input }),
    headers: { "Content-Type": "application/json" }
  });
  const data = await response.json();

  const botResponse = { text: data.reply, sender: "bot" };
  setMessages((prevMessages) => [...prevMessages, botResponse]);
};
