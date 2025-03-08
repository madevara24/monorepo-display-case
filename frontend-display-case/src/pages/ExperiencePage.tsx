import { FC } from 'react'
import Experience from '../components/Experience'
import styles from './ExperiencePage.module.css'

interface ExperiencePageProps {
  content: any
}

const ExperiencePage: FC<ExperiencePageProps> = ({ content }) => {
  return (
    <div className={styles.experiencePage}>
      <header className={styles.pageHeader}>
        <h1 className={styles.pageTitle}>Professional Experience</h1>
        <p className={styles.pageSubtitle}>
          A comprehensive overview of my professional journey and key achievements
        </p>
      </header>
      
      <Experience title={content.experience.title} items={content.experience.items} />
    </div>
  )
}

export default ExperiencePage 