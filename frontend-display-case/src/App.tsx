import Profile from './components/Profile'
import Experience from './components/Experience'
import content from './data/content.json'
import './App.css'

function App() {
  return (
    <div>
      <Profile data={content.profile} />
      <div className="section-separator" />
      <Experience title={content.experience.title} items={content.experience.items} />
      {/* Other components will go here */}
    </div>
  )
}

export default App
