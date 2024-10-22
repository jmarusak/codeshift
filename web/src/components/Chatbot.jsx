// src/components/Chatbot.js
import React, { useState } from "react";
import "./Chatbot.css";  // You can add custom styles in this file

const Chatbot = () => {
  // State to hold the user input and chat history
  const [input, setInput] = useState("");
  const [messages, setMessages] = useState([
    { text: "Hello! How can I help you?", sender: "bot" }
  ]);

  // Handle input submission
  const handleSubmit = (e) => {
    e.preventDefault();
    if (input.trim() === "") return;

    // Add the user's message to the conversation
    const userMessage = { text: input, sender: "user" };
    setMessages([...messages, userMessage]);

    // Clear the input field
    setInput("");

    // Simulate bot response (you can replace this with actual API call)
    setTimeout(() => {
      const botResponse = { text: `You said: ${input}`, sender: "bot" };
      setMessages((prevMessages) => [...prevMessages, botResponse]);
    }, 500);
  };

  return (
    <div className="chatbot-container">
      <div className="chat-window">
        <div className="messages">
          {messages.map((message, index) => (
            <div
              key={index}
              className={`message ${message.sender === "bot" ? "bot" : "user"}`}
            >
              {message.text}
            </div>
          ))}
        </div>
      </div>

      <form onSubmit={handleSubmit} className="input-form">
        <input
          type="text"
          value={input}
          onChange={(e) => setInput(e.target.value)}
          placeholder="Type your message..."
        />
        <button type="submit">Send</button>
      </form>
    </div>
  );
};

export default Chatbot;
