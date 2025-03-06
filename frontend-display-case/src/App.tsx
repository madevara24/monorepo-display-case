import Profile from './components/Profile'
import Experience from './components/Experience'
import content from './data/content.json'
import './App.css'
import { ThemeProvider } from './context/ThemeContext'
import ThemeToggle from './components/ThemeToggle/ThemeToggle'

function App() {
  return (
    <ThemeProvider>
      <Profile data={content.profile} />
      <div className="section-separator" />
      <Experience title={content.experience.title} items={content.experience.items} />
      <ThemeToggle />
      {/* Other components will go here */}
    </ThemeProvider>
  )
}

export default App
