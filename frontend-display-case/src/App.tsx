import { BrowserRouter, Routes, Route } from 'react-router-dom'
import HomePage from './pages/HomePage'
import ExperiencePage from './pages/ExperiencePage'
import Navbar from './components/Navbar'
import Footer from './components/Footer/Footer'
import content from './data/content.json'
import './App.css'
import { ThemeProvider } from './context/ThemeContext'

function App() {
  return (
    <ThemeProvider>
      <BrowserRouter>
        <div className="app-container">
          <Navbar />
          <main>
            <Routes>
              <Route path="/" element={<HomePage content={content} />} />
              <Route path="/experience" element={<ExperiencePage content={content} />} />
            </Routes>
          </main>
          <Footer />
        </div>
      </BrowserRouter>
    </ThemeProvider>
  )
}

export default App
