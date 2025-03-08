import { FC } from 'react'
import { Link } from 'react-router-dom'
import Profile from '../components/Profile'
import ExperienceCard from '../components/ExperienceCard'
import ChatBot from '../components/ChatBot'
import styles from './HomePage.module.css'

interface HomePageProps {
  content: any
}

const HomePage: FC<HomePageProps> = ({ content }) => {
  return (
    <div className={styles.homePage}>
      <Profile data={content.profile} />
      
      <div className={styles.sectionSeparator} />
      
      <section className={styles.experienceSummary}>
        <h2 className={styles.sectionTitle}>Work Experience</h2>
        <p className={styles.sectionSubtitle}>
          Companies I've worked with. <Link to="/experience" className={styles.viewAll}>View detailed experience →</Link>
        </p>
        
        <div className={styles.cardGrid}>
          {content.experience.items.map((item: any, index: number) => (
            <ExperienceCard 
              key={index}
              organization={item.organization}
              role={item.role}
              duration={item.duration}
              technologies={item.technologies || []}
              location={item.location}
              logo={item.logo}
            />
          ))}
        </div>
      </section>
      
      <div className={styles.sectionSeparator} />
      
      <ChatBot />
    </div>
  )
}

export default HomePage 