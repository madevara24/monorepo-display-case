import { BrowserRouter, Routes, Route } from 'react-router-dom'
import HomePage from './pages/HomePage'
import ExperiencePage from './pages/ExperiencePage'
import Navbar from './components/Navbar'
import content from './data/content.json'
import './App.css'
import { ThemeProvider } from './context/ThemeContext'

function App() {
  return (
    <ThemeProvider>
      <BrowserRouter>
        <Navbar />
        <Routes>
          <Route path="/" element={<HomePage content={content} />} />
          <Route path="/experience" element={<ExperiencePage content={content} />} />
        </Routes>
      </BrowserRouter>
    </ThemeProvider>
  )
}

export default App
