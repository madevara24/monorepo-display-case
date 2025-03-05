import Profile from './components/Profile'
import Experience from './components/Experience'
import './App.css'

const workExperience = [
  {
    organization: "Company Name",
    description: "A leading tech company focused on...",
    location: "Jakarta, Indonesia",
    duration: {
      start: "March 2020",
      end: "October 2021"
    },
    role: "Senior Backend Engineer",
    achievements: [
      "Led development of microservices architecture serving 1M+ users",
      "Reduced API response time by 40% through caching implementation",
      "Mentored 3 junior developers"
    ],
    technologies: ["Go", "PostgreSQL", "Redis", "Docker", "Kubernetes"]
  },
  // ... more experiences
]

function App() {
  return (
    <div>
      <Profile />
      <Experience title="Work Experience" items={workExperience} />
      {/* Other components will go here */}
    </div>
  )
}

export default App
