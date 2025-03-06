import { FC, useState } from 'react'
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
  const [expandedProjects, setExpandedProjects] = useState<Record<string, boolean>>({});
  const [expandedAchievements, setExpandedAchievements] = useState<Record<number, boolean>>({});

  const toggleProject = (itemIndex: number, projectIndex: number) => {
    const key = `${itemIndex}-${projectIndex}`;
    setExpandedProjects(prev => ({
      ...prev,
      [key]: !prev[key]
    }));
  };

  const toggleAchievements = (itemIndex: number) => {
    setExpandedAchievements(prev => ({
      ...prev,
      [itemIndex]: !prev[itemIndex]
    }));
  };

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

            {item.technologies && (
              <div className={styles.technologies}>
                <span>Technologies: </span>
                {item.technologies.map((tech, i) => (
                  <span key={i} className={styles.tech}>
                    {tech}
                  </span>
                ))}
              </div>
            )}

            {/* Render general achievements if no projects */}
            {item.achievements && !item.projects && (
              <div className={styles.projectsList}>
                <div className={styles.projectContainer}>
                  <button
                    className={styles.projectToggle}
                    onClick={() => toggleAchievements(index)}
                    aria-expanded={expandedAchievements[index]}
                    aria-controls={`achievements-${index}`}
                  >
                    <div className={styles.projectName}>
                      Key Achievements
                      <span className={styles.toggleIcon} aria-hidden="true">
                        {expandedAchievements[index] ? '−' : '+'}
                      </span>
                    </div>
                  </button>
                  
                  <div 
                    id={`achievements-${index}`}
                    className={`${styles.expandableContent} ${expandedAchievements[index] ? styles.expanded : ''}`}
                  >
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
                  </div>
                </div>
              </div>
            )}

            {/* Render projects if they exist */}
            {item.projects && (
              <div className={styles.projectsSection}>
                <h4 className={styles.projectsSectionTitle}>Key Projects & Initiatives</h4>
                <div className={styles.projectsList}>
                  {item.projects.map((project, pIndex) => (
                    <div key={pIndex} className={styles.projectContainer}>
                      <button
                        className={styles.projectToggle}
                        onClick={() => toggleProject(index, pIndex)}
                        aria-expanded={expandedProjects[`${index}-${pIndex}`]}
                        aria-controls={`project-${index}-${pIndex}`}
                      >
                        <div className={styles.projectName}>
                          {project.name}
                          <span className={styles.toggleIcon} aria-hidden="true">
                            {expandedProjects[`${index}-${pIndex}`] ? '−' : '+'}
                          </span>
                        </div>
                      </button>
                      
                      <div 
                        id={`project-${index}-${pIndex}`}
                        className={`${styles.expandableContent} ${expandedProjects[`${index}-${pIndex}`] ? styles.expanded : ''}`}
                      >
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
                    </div>
                  ))}
                </div>
              </div>
            )}
          </div>
        ))}
      </div>
    </section>
  )
}

export default Experience