import { FC, useState, useRef, useEffect } from 'react'
import styles from './ChatBot.module.css'
import ReactMarkdown from 'react-markdown'

interface Message {
  id: string
  type: 'user' | 'bot'
  content: string
  timestamp: Date
}

const ChatBot: FC = () => {
  const [messages, setMessages] = useState<Message[]>([
    {
      id: '1',
      type: 'bot',
      content: "Hi! I'm Aditya's AI assistant. Feel free to ask me anything about his experience, projects, or skills!",
      timestamp: new Date()
    }
  ])
  const [input, setInput] = useState('')
  const [isTyping, setIsTyping] = useState(false)
  const messagesEndRef = useRef<HTMLDivElement>(null)

  const scrollToBottom = () => {
    messagesEndRef.current?.scrollIntoView({ behavior: 'smooth' })
  }

  useEffect(() => {
    scrollToBottom()
  }, [messages])

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    if (!input.trim()) return

    const userMessage: Message = {
      id: Date.now().toString(),
      type: 'user',
      content: input,
      timestamp: new Date()
    }

    setMessages(prev => [...prev, userMessage])
    setInput('')
    setIsTyping(true)

    // Placeholder for API call
    // For now, just show a mock response after a delay
    setTimeout(() => {
      const botMessage: Message = {
        id: (Date.now() + 1).toString(),
        type: 'bot',
        content: "I'm currently in development, but I'll be able to help you learn more about Aditya's experience soon!",
        timestamp: new Date()
      }
      setMessages(prev => [...prev, botMessage])
      setIsTyping(false)
    }, 1000)
  }

  return (
    <section className={styles.chatbot}>
      <div className={styles.header}>
        <h2>Chat with AI Assistant</h2>
        <p className={styles.subtitle}>Ask about my experience, projects, or skills</p>
      </div>

      <div className={styles.chatContainer}>
        <div className={styles.messages}>
          {messages.map(message => (
            <div
              key={message.id}
              className={`${styles.message} ${
                message.type === 'user' ? styles.userMessage : styles.botMessage
              }`}
            >
              <div className={styles.messageContent}>
                <ReactMarkdown
                  components={{
                    strong: ({node, ...props}) => <strong className={styles.markdown} {...props} />,
                    a: ({node, ...props}) => <a className={styles.markdown} {...props} />,
                    em: ({node, ...props}) => <em className={styles.markdown} {...props} />,
                    code: ({node, ...props}) => <code className={styles.markdown} {...props} />
                  }}
                >
                  {message.content}
                </ReactMarkdown>
              </div>
              <div className={styles.timestamp}>
                {message.timestamp.toLocaleTimeString([], { 
                  hour: '2-digit', 
                  minute: '2-digit'
                })}
              </div>
            </div>
          ))}
          {isTyping && (
            <div className={`${styles.message} ${styles.botMessage}`}>
              <div className={styles.typingIndicator}>
                <span></span>
                <span></span>
                <span></span>
              </div>
            </div>
          )}
          <div ref={messagesEndRef} />
        </div>

        <form onSubmit={handleSubmit} className={styles.inputContainer}>
          <input
            type="text"
            value={input}
            onChange={(e) => setInput(e.target.value)}
            placeholder="Type your message..."
            className={styles.input}
          />
          <button type="submit" className={styles.sendButton}>
            Send
          </button>
        </form>
      </div>
    </section>
  )
}

export default ChatBot 