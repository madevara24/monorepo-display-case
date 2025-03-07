import { FC, useState, useRef, useEffect } from 'react'
import ReactMarkdown from 'react-markdown'
import styles from './ChatBot.module.css'
import content from '../../data/content.json'

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
      content: "Hi! I'm Devara's AI assistant. Feel free to ask me anything about his experience, projects, or skills!",
      timestamp: new Date()
    }
  ])
  const [input, setInput] = useState('')
  const [isTyping, setIsTyping] = useState(false)
  const [isBackendAvailable, setIsBackendAvailable] = useState(true)
  const messagesEndRef = useRef<HTMLDivElement>(null)
  
  // Get API URL and result limit from environment variables with fallbacks
  const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080'
  const RESULT_LIMIT = import.meta.env.VITE_RESULT_LIMIT || 10

  // Check backend health when component mounts
  useEffect(() => {
    checkBackendHealth()
    
    // Set up periodic health check every 30 seconds
    const healthCheckInterval = setInterval(() => {
      checkBackendHealth()
    }, 30000)
    
    return () => clearInterval(healthCheckInterval)
  }, [])

  const checkBackendHealth = async () => {
    try {
      console.log("Checking health at:", API_URL)
      const response = await fetch(`${API_URL}/health`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
        credentials: 'include',
        signal: AbortSignal.timeout(3000)
      })
      
      const isAvailable = response.ok
      console.log("Backend available:", isAvailable)
      
      // If backend just came back online, add a message
      if (isAvailable && !isBackendAvailable) {
        const reconnectedMessage: Message = {
          id: Date.now().toString(),
          type: 'bot',
          content: "I'm back online and ready to answer your questions!",
          timestamp: new Date()
        }
        setMessages(prev => [...prev, reconnectedMessage])
      }
      
      setIsBackendAvailable(isAvailable)
    } catch (error) {
      console.error('Health check failed:', error)
      if (error instanceof TypeError && error.message.includes('Failed to fetch')) {
        console.error('Possible CORS error. Make sure backend CORS settings are correct')
      }
      setIsBackendAvailable(false)
    }
  }

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

    // Check if backend is available before making API call
    if (!isBackendAvailable) {
      // Add delay to simulate thinking
      setTimeout(() => {
        const offlineMessage: Message = {
          id: (Date.now() + 1).toString(),
          type: 'bot',
          content: "I'm currently unable to connect to my knowledge database. Please try again later or contact Devara directly if you have any urgent questions.",
          timestamp: new Date()
        }
        setMessages(prev => [...prev, offlineMessage])
        setIsTyping(false)
      }, 1000)
      return
    }

    try {
      const response = await fetch(`${API_URL}/ask`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        credentials: 'include',
        body: JSON.stringify({ 
          question: input,
          limit: Number(RESULT_LIMIT)
        }),
      })

      if (!response.ok) {
        throw new Error('Backend request failed')
      }
      
      const data = await response.json()
      
      const botMessage: Message = {
        id: (Date.now() + 1).toString(),
        type: 'bot',
        content: data.data.answer || "I'm sorry, I couldn't process your request at this time.",
        timestamp: new Date()
      }
      
      setMessages(prev => [...prev, botMessage])
    } catch (error) {
      console.error('Error:', error)
      setIsBackendAvailable(false)
      
      const errorMessage: Message = {
        id: (Date.now() + 1).toString(),
        type: 'bot',
        content: "I'm having trouble connecting to my knowledge database. Please try again later.",
        timestamp: new Date()
      }
      
      setMessages(prev => [...prev, errorMessage])
    } finally {
      setIsTyping(false)
    }
  }

  return (
    <section className={styles.chatbot}>
      <div className={styles.header}>
        <h2>AI Chatbot</h2>
        <p className={styles.subtitle}>Want to know more about me but still unsure wether to send me an email? Just ask this chatbot below! I'm still tinkering with the model and the data structure but it should do the trick for short simple questions.</p>
        <p className={styles.disclaimer}>
          This AI may occasionally provide inaccurate information. If you notice any errors, 
          please <a href={content.profile.links.email.url}>email me</a> directly.
        </p>
        {!isBackendAvailable && (
          <div className={styles.statusIndicator}>
            <span className={styles.offlineIcon}>⚠️</span> 
            <span className={styles.offlineText}>Service currently unavailable</span>
          </div>
        )}
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
            disabled={!isBackendAvailable && isTyping}
          />
          <button 
            type="submit" 
            className={styles.sendButton}
            disabled={!isBackendAvailable && isTyping}
          >
            Send
          </button>
        </form>
      </div>
    </section>
  )
}

export default ChatBot 