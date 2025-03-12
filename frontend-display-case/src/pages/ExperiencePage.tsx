import { FC } from 'react'
import Experience from '../components/Experience'
import Education from '../components/Education'
import styles from './ExperiencePage.module.css'

interface ExperiencePageProps {
  content: any
}

// Define education item from content.json
interface EducationItem {
  name: string
  location: string
  duration: {
    start: string
    end: string
  }
  degree: string
  major: string
  gpa: string
  highlights?: string[]
  technologies?: string[]
}

const ExperiencePage: FC<ExperiencePageProps> = ({ content }) => {
  return (
    <div className={styles.experiencePage}>
      <header className={styles.pageHeader}>
        <h1 className={styles.pageTitle}>Professional & Academic Journey</h1>
        <p className={styles.pageSubtitle}>
          My professional career path and educational achievements
        </p>
      </header>
      
      <Experience title={content.experience.title} items={content.experience.items} />
      
      <div className={styles.educationSection}>
        <Education title={content.education.title} items={
          content.education.items.map((item: EducationItem) => ({
            institution: item.name,
            description: `${item.degree} in ${item.major}. GPA: ${item.gpa}`,
            location: item.location,
            duration: item.duration,
            degree: item.degree,
            highlights: item.highlights || [],
            technologies: item.technologies || []
          }))
        } />
      </div>
    </div>
  )
}

export default ExperiencePage 