import { FC } from 'react'
import ReactMarkdown from 'react-markdown'
import styles from './Experience.module.css'

interface Project {
  name: string
  achievements: string[]
}

interface ExperienceItem {
  organization: string
  description: string
  location?: string
  duration: {
    start: string
    end: string
  }
  role: string
  projects?: Project[]
  achievements?: string[]
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
              <h3 className={styles.organization}>
                {item.organization} <span className={styles.separator}>|</span> <span className={styles.role}>{item.role}</span>
              </h3>
              {/* <span className={styles.location}>{item.location}</span> */}
            </div>
            
            <div className={styles.duration}>
              {item.duration.start} - {item.duration.end}
            </div>

            <p className={styles.description}>{item.description}</p>

            {/* Render general achievements if no projects */}
            {item.achievements && (
              <ul className={styles.achievements}>
                {item.achievements.map((achievement, i) => (
                  <li key={i}>
                    <ReactMarkdown
                      components={{
                        strong: ({node, ...props}) => <strong className={styles.markdown} {...props} />,
                        a: ({node, ...props}) => <a className={styles.markdown} {...props} />,
                        em: ({node, ...props}) => <em className={styles.markdown} {...props} />,
                        code: ({node, ...props}) => <code className={styles.markdown} {...props} />
                      }}
                    >
                      {achievement}
                    </ReactMarkdown>
                  </li>
                ))}
              </ul>
            )}

            {/* Render projects if they exist */}
            {item.projects && item.projects.map((project, pIndex) => (
              <div key={pIndex} className={styles.project}>
                <h4 className={styles.projectName}>{project.name}</h4>
                <ul className={styles.achievements}>
                  {project.achievements.map((achievement, i) => (
                    <li key={i}>
                      <ReactMarkdown
                        components={{
                          strong: ({node, ...props}) => <strong className={styles.markdown} {...props} />,
                          a: ({node, ...props}) => <a className={styles.markdown} {...props} />,
                          em: ({node, ...props}) => <em className={styles.markdown} {...props} />,
                          code: ({node, ...props}) => <code className={styles.markdown} {...props} />
                        }}
                      >
                        {achievement}
                      </ReactMarkdown>
                    </li>
                  ))}
                </ul>
              </div>
            ))}

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