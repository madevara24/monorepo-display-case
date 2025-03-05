import { FC } from 'react'
import styles from './Experience.module.css'

interface ExperienceItem {
  organization: string
  description: string
  location: string
  duration: {
    start: string
    end: string
  }
  role: string
  achievements: string[]
  technologies?: string[] // Optional for education entries
}

interface ExperienceProps {
  title: string // "Work Experience" or "Education"
  items: ExperienceItem[]
}

const Experience: FC<ExperienceProps> = ({ title, items }) => {
  return (
    <section className={styles.experience}>
      <h2 className={styles.sectionTitle}>{title}</h2>
      <div className={styles.timeline}>
        {items.map((item, index) => (
          <div key={index} className={styles.timelineItem}>
            <div className={styles.header}>
              <h3 className={styles.organization}>{item.organization}</h3>
              <span className={styles.location}>{item.location}</span>
            </div>
            
            <div className={styles.subHeader}>
              <span className={styles.role}>{item.role}</span>
              <span className={styles.duration}>
                {item.duration.start} - {item.duration.end}
              </span>
            </div>

            <p className={styles.description}>{item.description}</p>

            <ul className={styles.achievements}>
              {item.achievements.map((achievement, i) => (
                <li key={i}>{achievement}</li>
              ))}
            </ul>

            {item.technologies && (
              <div className={styles.technologies}>
                {item.technologies.map((tech, i) => (
                  <span key={i} className={styles.tech}>
                    {tech}
                  </span>
                ))}
              </div>
            )}
          </div>
        ))}
      </div>
    </section>
  )
}

export default Experience